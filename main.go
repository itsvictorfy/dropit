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
	c.HTML(http.StatusOK, "HomePage.html", nil)
	c.SetCookie("lastTimeentered", time.Now().String(), 60*60*48, "", "", false, true)
	fmt.Println(c.Request.Cookie("AuthenticationCookie"))
	//isAuthorized(c)
}
func userpage(c *gin.Context) {
	c.HTML(http.StatusOK, "userpage.html", nil)
}

func registerPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}
func logout(c *gin.Context) {
	c.SetCookie("AuthenticationCookie", "expired", -1, "", "", false, false)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "admin",
		Net:                  "tcp",
		Addr:                 "0.0.0.0:3306",
		DBName:               "userdDB",
		AllowNativePasswords: true,
	}
	var err error
	usersDb, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := usersDb.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
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

	router.GET("/", homePage)
	router.GET("/login", login)
	router.GET("/logout", logout)
	router.GET("/userpage", userpage)
	router.POST("/regaval", register)
	router.POST("/loginevr", login)
	router.GET("/register", registerPage)

	apiReq := router.Group("/apireq")
	{
		apiReq.GET("/csv", reinforstCSV)     //using Rainforst API for spesific Amazon product data to CSV
		apiReq.GET("/json", reinforstJSON)   //using Rainforest API for spesific Amazon Product data to JSON
		apiReq.POST("/search", searchResult) //using Axesso API for ASIN numbers from search results by Keyboard to JSON
	}
	router.Run() // listen and serve on 0.0.0.0:8080
	fmt.Println(ksuid.New())
}

/*
IDEAS:
-monitoring all logs with Elsatic Search and represent number of issues / requests
-Side Quest : calculate Formula 1 pottential winning round for driver if possible
-side quest Email verefication API
*/
