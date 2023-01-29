package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/segmentio/ksuid"
)

func login(c *gin.Context) {
	loginEmail := c.PostForm("loginEmail")
	loginPass := c.PostForm("loginPass")
	token, err := loginVerification(loginEmail, loginPass)
	if err != nil {
		log.Panicf("Incorrect User/Password %s", err.Error())
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		log.Printf("Login success ")
		createCookie(token, c)
		usersDb.Query("update DropItUsersDB set LastLogin = ?  WHERE Email = ?;", time.Now(), loginEmail)
		fmt.Println("\n Login Succes coockie:")
		fmt.Println(c.Request.Cookie("AuthenticationCookie"))
		c.Redirect(http.StatusMovedPermanently, "/userpage")

	}
}

func loginVerification(loginEmail, loginPass string) (string, error) {
	var passFromDB string
	var idFromDB ksuid.KSUID
	var roleFromDB string
	var validToken string
	err := usersDb.QueryRow("SELECT Passw, URole, Secret_Key FROM `DropItUsersDB` WHERE (Email= ? );", loginEmail).Scan(&passFromDB, &roleFromDB, &idFromDB)
	if err != nil {
		log.Println("error getting DB Data")
	}
	if err != nil {
		log.Panicln("no such user")
	}
	if loginPass != passFromDB {
		log.Println("password is incorrect")
		return "", err
	} else {
		validToken, err = generateJWT(loginEmail, idFromDB, roleFromDB)
		if err != nil {
			log.Panicln("Error with creating Token")
			return "", err
		}
	}
	return validToken, nil
}
func generateJWT(email string, id ksuid.KSUID, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["id"] = id
	claims["role"] = role
	claims["experation"] = time.Now().Add(time.Hour * 48).Unix()

	tokenString, err := token.SignedString(secretkey)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	fmt.Printf("tokenstring= %s", tokenString)
	return tokenString, nil
}

func createCookie(token string, c *gin.Context) {
	c.SetCookie("AuthenticationCookie", token, 60*60*48, "", "", false, true)
}

// check if coocie legit
func isAuthorized(c *gin.Context) bool {
	cookie, err := c.Request.Cookie("AuthenticationCookie")
	if err != nil {
		log.Printf("No cookie Found")
		return false
	}
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
			return nil, fmt.Errorf("something went wrong")
		}
		return secretkey, nil
	})
	if err != nil {
		log.Printf("cant parse Token")
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Printf("cant map claims ")
		return false
	}
	idFromCookie, ok := claims["id"].(string)
	if !ok {
		log.Printf("cant claim id from Cookie")
		return false
	}
	emailFromCookie, ok := claims["email"].(string)
	if !ok {
		log.Printf("cant claim email from Cookie")
		return false
	}
	var idFromDb string
	errs := usersDb.QueryRow("SELECT Secret_Key FROM `DropItUsersDB` WHERE (Email= ? );", emailFromCookie).Scan(&idFromDb)
	if errs != nil {
		log.Println("error getting DB Data")
		return false
	}
	if idFromDb != idFromCookie {
		log.Printf("Cookie ID != DB ID")
		return false
	}
	log.Printf("is Authorized")
	return true
}
