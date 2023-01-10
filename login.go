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
	fmt.Printf("loginEmail : %s \n", loginEmail)
	fmt.Printf("LoginPass : %s \n", loginPass)
	token, err := loginVerification(loginEmail, loginPass)
	if err != nil {
		log.Panicf("Incorrect User/Password %s", err.Error())
		c.Redirect(http.StatusMovedPermanently, "/error")
	} else {
		createCookie(token, c)
		c.Redirect(http.StatusMovedPermanently, "/userpage")
		fmt.Println("\n Login Succes coockie:")
		fmt.Println(c.Request.Cookie("AuthenticationCookie"))

		fmt.Println(c.Request.Cookie("AuthenticationCookie"))
	}
}

func loginVerification(loginEmail, loginPass string) (string, error) {
	var passFromDB string
	var idFromDB ksuid.KSUID
	var roleFromDB string
	var validToken string
	err := usersDb.QueryRow("SELECT Passw, URole, Secret_Key FROM `DropItUsersDB` WHERE (Email= ? );", loginEmail).Scan(&passFromDB, &roleFromDB, &idFromDB)
	fmt.Printf("idFromDB : %s \n", idFromDB)
	fmt.Printf("roleFromDB : %s \n", roleFromDB)
	fmt.Printf("passFromDB : %s \n", passFromDB)

	fmt.Println(err)
	if err != nil {
		log.Panicln("no such user")

	}
	if loginPass != passFromDB {
		log.Println("password is incorrect")
	} else {
		validToken, err = generateJWT(loginEmail, idFromDB, roleFromDB)
		fmt.Printf("validtoke= %s", validToken)
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

/*func isAuthorized(c *gin.Context) {
	cookie, err := c.Request.Cookie("AuthenticationCookie")
	if err != nil {
		if err == http.ErrNoCookie {
			c.Redirect(http.StatusUnauthorized, "/error")
		}
		c.Redirect(http.StatusBadRequest, "/error")
	}
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("something went wrong")
		}
		return secretkey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.Redirect(http.StatusUnauthorized, "/error")
		}
		c.Redirect(http.StatusBadRequest, "/error")
	}
	if token.Valid {
		c.Redirect(http.StatusOK, "/userpage")
	}
}*/
