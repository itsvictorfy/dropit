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

func login_auth(c *gin.Context) {
	loginEmail := c.PostForm("login_email")
	loginPass := c.PostForm("login_password")
	token, err := loginVerification(loginEmail, loginPass)
	if err != nil {
		log.Printf("Incorrect User/Password %s \n", err.Error())
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		log.Printf("Login success for %s \n", loginEmail)
		createCookie(token, c)
		usersDb.Query("update DropItUsersDB set LastLogin = ?  WHERE Email = ?;", time.Now(), loginEmail)
		log.Printf("DB Update last login for %s to %s", loginEmail, time.Now())
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
		log.Printf("No Such user : %s", err)
		return "", err
	}
	if loginPass != passFromDB {
		log.Printf("Password Doesnt Match")
		return "", err
	} else {
		validToken, err = generateJWT(loginEmail, idFromDB, roleFromDB)
		if err != nil {
			log.Printf("Error with creating Token")
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
		log.Printf("Error signing String : %s\n", err.Error())
		return "", err
	}
	log.Printf("Token string Created for %s : %s\n", email, tokenString)
	return tokenString, nil
}

func createCookie(token string, c *gin.Context) {
	c.SetCookie("AuthenticationCookie", token, 60*60*48, "", "", false, true)
	log.Printf("Cookie Created:\n")
	log.Println(c.Request.Cookie("AuthenticationCookie"))
}

func isAuthorized(c *gin.Context) bool {
	cookie, err := c.Request.Cookie("AuthenticationCookie")
	if err != nil {
		log.Printf("No cookie Found: %s", err)
		return false
	}
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
			return nil, fmt.Errorf("something went wrong")
		}
		return secretkey, nil
	})
	if err != nil {
		log.Printf("Cant parse Token: %s", err)
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Printf("Cant map claims")
		return false
	}
	idFromCookie, ok := claims["id"].(string)
	if !ok {
		log.Printf("Cant claim id from Cookie")
		return false
	}
	emailFromCookie, ok := claims["email"].(string)
	if !ok {
		log.Printf("Cant claim email from Cookie")
		return false
	}
	var idFromDb string
	errs := usersDb.QueryRow("SELECT Secret_Key FROM `users` WHERE (Email= ? );", emailFromCookie).Scan(&idFromDb)
	if errs != nil {
		log.Printf("error getting DB Data")
		return false
	}
	if idFromDb != idFromCookie {
		log.Printf("Cookie ID != DB ID")
		return false
	}
	log.Printf("is Authorized")
	return true
}
