package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/segmentio/ksuid"
)

/*
	type product struct {
		asin          string  `json:"asin"`
		productTitle  string  `json:"Product Title"`
		price         float32 `json:"price"`
		productRating float32 `json:"productRating"`
		sellerID      string  `json:"Seller ID"`
		imageUrl      string  `json:"Image Url"`
		//keywords      []string //depends on API Used Rainforest provides but Axesso doesnt
	}*/

type user struct {
	SecretKey    ksuid.KSUID
	Role         string    `json:"Role"`
	Email        string    `json:"Email"`
	Pass         string    `json:"Password"`
	FName        string    `json:"FirstName"`
	LName        string    `json:"Last Name"`
	LastSearch   string    `json:"Last Search"`
	LastLogin    string    `json:"Last Login"`
	Verified     bool      `json:"verified"`
	CreationDate time.Time `json:"Created "`
}

var secretkey = []byte("whatsecretkeywedonthaveasecretkey")
var usersDb *sql.DB

// handle root directory
func homePage(c *gin.Context) {
	if !isAuthorized(c) {
		c.HTML(http.StatusOK, "HomePage.html", nil)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/userpage")
	}
}
func userpage(c *gin.Context) {
	if !isAuthorized(c) {
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		c.HTML(http.StatusOK, "userpage.html", nil)
	}
}

func registerPage(c *gin.Context) {
	if !isAuthorized(c) {
		c.HTML(http.StatusOK, "register.html", nil)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/userpage")
	}
}
func logout(c *gin.Context) {
	c.SetCookie("AuthenticationCookie", "expired", -1, "", "", false, false) //cant delete the fucking cookie
	c.Redirect(http.StatusMovedPermanently, "/")
}

func getCookie(c *gin.Context) {
	fmt.Println(c.Request.Cookie("AuthenticationCookie"))
}

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "admin",
		Net:                  "tcp",
		Addr:                 "userDB",
		DBName:               "userdDB",
		AllowNativePasswords: true,
	}
	var err error
	usersDb, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Println(err)
	}

	pingErr := usersDb.Ping()
	if pingErr != nil {
		log.Println(pingErr)
	}

	var version string
	usersDb.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	router := gin.Default()
	router.LoadHTMLGlob("pages/*.html")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", homePage)             //home page (no cookie)
	router.GET("/logout", logout)         // handles logout (delete cookie)
	router.GET("/userpage", userpage)     // home page ( with cookie)
	router.POST("/regaval", register)     // handles registration
	router.POST("/loginver", login)       //handles login
	router.GET("/register", registerPage) //register page
	router.GET("/cookie", getCookie)      //check cookie

	apiReq := router.Group("/apireq")
	{
		apiReq.GET("/csv", reinforstCSV)     //using Rainforst API for spesific Amazon product data to CSV
		apiReq.GET("/json", reinforstJSON)   //using Rainforest API for spesific Amazon Product data to JSON
		apiReq.POST("/search", searchResult) //using Axesso API for ASIN numbers from search results by Keyboard to JSON
		apiReq.GET("/newapi", newapi)
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}

/*
IDEAS:
-monitoring all logs with Elsatic Search and represent number of issues / requests
-Side Quest : calculate Formula 1 pottential winning round for driver if possible
-side quest Email verefication API

CREATE TABLE IF NOT EXISTS dropitusersdb(
    Secret_Key VARCHAR(255) NOT NULL,
	URole VARCHAR(255) NOT NULL,
	Email VARCHAR (255) NOT NULL,
	Passw VARCHAR (255) NOT NULL,
	First_Name VARCHAR (255) NOT NULL,
	Last_Name VARCHAR (255) NOT NULL,
	LastSearch VARCHAR (255) NULL,
	LastLogin DATE NULL,
	Is_Verified BOOLEAN NOT NULL,
	Creation_Date DATE NOT NULL
) COMMENT '';
*/
