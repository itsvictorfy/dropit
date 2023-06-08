package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
	"github.com/segmentio/ksuid"
)

type bsr struct {
	Category string `json:"bestsellers_Category"`
	Rank     string `json:"bestsellers_rank"`
	Link     string `json:"bestsellers_Link"`
}

type product struct { //from getProducts 24 items
	Asin                   string  `json:"Asin"`
	Title                  string  `json:"Title"`
	Link                   string  `json:"Link"`
	Price                  float64 `json:"Price"`
	Rating                 float64 `json:"rating"`
	TotalRatings           int     `json:"TotalRating"`
	SellerID               string  `json:"SellerID"`
	ImageUrl               string  `json:"ImgaeURL"`
	MonthlySalesEstimate   float64 `json:"MonthlySalesEstimate"`
	MonthlyRevenueEstimate float64 `json:"MonthlyRevenueEstimate"`
	BestsellersRank        []bsr   `json:"bestsellers_rank"`
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

type keywordAsins struct {
	Keyword string   `json:"keyword"`
	Asins   []string `json:"asins"`
}

var secretkey = []byte("whatsecretkeywedonthaveasecretkey")

// var googlekey := "AIzaSyAo9yLA6tPkzuvFdBHu-ySnccv-y5IVIc0"
var usersDb *sql.DB
var productCache *redis.Client
var keywordCache *redis.Client

// handle root directory
func homePage(c *gin.Context) {
	if !isAuthorized(c) {
		c.HTML(http.StatusOK, "home_page.html", nil)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/userpage")
	}
}
func userpage(c *gin.Context) {
	if !isAuthorized(c) {
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		c.HTML(http.StatusOK, "user_page.html", nil)
	}
}

func registerPage(c *gin.Context) {
	if !isAuthorized(c) {
		c.HTML(http.StatusOK, "register_page.html", nil)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/userpage")
	}
}
func logout(c *gin.Context) {
	c.SetCookie("AuthenticationCookie", "expired", -1, "", "", false, true)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func loginpage(c *gin.Context) {
	c.HTML(http.StatusOK, "login_page.html", nil)
}

func basicSearch(c *gin.Context) {
	sProducts := getProdListFromApi(searchQuery(c))
	tmpl := template.Must(template.ParseFiles("./templates/search_result.html"))
	err := tmpl.Execute(c.Writer, sProducts)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
	// //c.JSON(http.StatusOK, sProducts)
}

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "admin",
		Net:    "tcp",
		Addr:   "dropit_users_DB",
		//Addr:                 "127.0.0.1:3306",
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

	productCache = redis.NewClient(&redis.Options{
		Addr: "dropit_prod_cache:6379",
		// Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	pong, err := productCache.Ping().Result()
	if err != nil {
		fmt.Printf("Pong: %s, Err: %v", pong, err)
	}
	log.Printf("Connected to DB1): %s", pong)

	keywordCache = redis.NewClient(&redis.Options{
		Addr: "dropit_prod_cache:6379",
		//Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	})
	pong1, err := keywordCache.Ping().Result()
	if err != nil {
		fmt.Printf("Pong: %s, Err: %v", pong1, err)
	}
	log.Printf("Connected to DB1): %s", pong1)
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.*")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", homePage)                 //Home page (no cookie)
	router.GET("/login", loginpage)           //Login Page
	router.GET("/register", registerPage)     //register page
	router.GET("/userpage", userpage)         // User Home Page
	router.POST("/fsearch", basicSearch)      //basic search for any user
	router.Static("/css/", "./templates/css") //get css

	auth := router.Group("/auth") //all authentication and registretion
	{
		auth.POST("/login", login_auth) //handle Login
		auth.GET("/logout", logout)     // handles logout (delete cookie)
		reg := router.Group("/reg")
		{
			reg.POST("/newrequest", register) //handle registration in register.go
		}
	}
	apiReq := router.Group("/apireq")
	{
		apiReq.POST("/search", homePage) //unfinished
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}

/*
IDEAS:
-monitoring all logs with Elsatic Search and represent number of issues / requests Prometheus
-adding google trends per product / Group



//text to speach plugin for chrome
-Side Quest : calculate Formula 1 pottential winning round for driver if possible
-side quest Email verefication API
-export all datas to env vars
*/
