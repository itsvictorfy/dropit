package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/segmentio/ksuid"
)

type product struct { //from getProducts 24 items
	Asin                   string
	Title                  string
	Link                   string
	Price                  float64
	Rating                 float64
	TotalRatings           int
	SellerID               string
	ImageUrl               string
	MonthlySalesEstimate   float64
	MonthlyRevenueEstimate float64
	TotalProducts          int
	BestsellersRank        []struct {
		Category string
		Rank     int
		Link     string
	}
}

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
		c.HTML(http.StatusOK, "home_page.html", nil)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/auth/userpage")
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
		c.HTML(http.StatusOK, "register_page.html", nil)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/auth/userpage")
	}
}
func logout(c *gin.Context) {
	c.SetCookie("AuthenticationCookie", "expired", -1, "", "", false, false)
	c.Redirect(http.StatusMovedPermanently, "/")
}
func testpage(c *gin.Context) {
	c.HTML(http.StatusOK, "HomePage.html", nil)
	c.SetCookie("AuthenticationCookie", "expired", -1, "", "", false, false)
}

func loginpage(c *gin.Context) {
	c.HTML(http.StatusOK, "login_page.html", nil)
}

func basicSearch(c *gin.Context) {
	sProducts := getProducts(searchQuery(c))
	tmpl := template.Must(template.ParseFiles("test.html"))
	tmpl.Execute(c.Writer, sProducts)
	//tmpl.Execute(c.Writer, product2)

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
	router.LoadHTMLGlob("templates/*.*")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", homePage)                //home page (no cookie)
	router.GET("/login", loginpage)          //Login Page
	router.POST("/loginauth", login_auth)    //login authentication
	router.POST("/basicsearch", basicSearch) //basic search for any new user
	router.GET("/test", testpage)            //testing - atm basic search

	router.Static("/css/", "./templates/css")

	reg := router.Group("/register")
	{
		reg.POST("/newRequest", register) //handle registration in register.go
		reg.GET("", registerPage)         //register page
	}
	auth := router.Group("/auth")
	{
		auth.GET("/userpage", userpage) // home page
		auth.GET("/logout", logout)     // handles logout (delete cookie)

		apiReq := router.Group("/apireq")
		{
			apiReq.POST("/search", searchNew) //unfinished
		}
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}

/*
IDEAS:
-monitoring all logs with Elsatic Search and represent number of issues / requests
-adding google trends per product / Group



-Side Quest : calculate Formula 1 pottential winning round for driver if possible
-side quest Email verefication API
-export all datas to env vars

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
