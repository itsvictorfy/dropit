package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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
	}

	type user struct {
		id           uuid.UUID
		role         string    `json:"Role"`
		email        string    `json:"Email"`
		pass         string    `json:"Password"`
		fName        string    `json:"FirstName"`
		lName        string    `json:"Last Name"`
		lastSearch   string    `json:"Last Search"`
		lastLogin    string    `json:"Last Login"`
		verified     bool      `json:"verified"`
		creationDate time.Time `json:"Created "`
	}
*/

var secretkey = []byte("whatsecretkeywedonthaveasecretkey")

// handle root directory
func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "HomePage.html", nil)
	c.SetCookie("lastTimeentered", time.Now().String(), 60*60*48, "", "", false, true)
	fmt.Println(c.Request.Cookie("AuthenticationCookie"))
	//isAuthorized(c)
}

// use Rainforest API to get Product Data to Json
func reinforstJSON(c *gin.Context) {
	response, err := http.Get("https://api.rainforestapi.com/request?api_key=63FFD605C840421D9F5FC4433C106F90&amazon_domain=amazon.com&asin=B08P4WR6XB&type=product&output=json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	fmt.Println(response)

	file, _ := json.MarshalIndent(string(responseData), "", " ")
	_ = os.WriteFile("data.json", file, 0644)

}

// use Rainforest API to get Product Data to CSV
func reinforstCSV(c *gin.Context) {
	response, err := http.Get("https://api.rainforestapi.com/request?api_key=63FFD605C840421D9F5FC4433C106F90&amazon_domain=amazon.com&asin=B073JYC4XM&type=product&output=csv")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(responseData))

}

// uses Axesso API to get ASINs from keyword search
func searchResult(c *gin.Context) {
	url := editQueryParm(c)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "fa931de48dmshdfc1a4a261c7e1cp129019jsn7f650c68afb5")
	req.Header.Add("X-RapidAPI-Host", "axesso-axesso-amazon-data-service-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	//unable to write formated data to file - Please help See data,data1.json files
	file, _ := json.Marshal(body)
	_ = os.WriteFile("data.json", file, 0644)

	file1, _ := json.Marshal(string(body))
	_ = os.WriteFile("data1.json", file1, 0644)

}

// edits Axesso API query Parametrs by keyworkd searched
func editQueryParm(c *gin.Context) string {
	url, err := url.Parse("https://axesso-axesso-amazon-data-service-v1.p.rapidapi.com/amz/amazon-search-by-keyword-asin?domainCode=com&keyword=Jeans&page=1&excludeSponsored=false&sortBy=relevanceblender&withCache=true")
	if err != nil {
		log.Println(err)
	}

	values := url.Query()
	keyword := c.PostForm("keyword")
	values.Set("keyword", keyword)

	url.RawQuery = values.Encode()
	fmt.Print(url)
	return string(url.String())
}

func userpage(c *gin.Context) {
	c.HTML(http.StatusOK, "userpage.html", nil)
	isAuthorized(c)
}

// handles /login
func login(c *gin.Context) {
	loginEmail := c.PostForm("loginEmail")
	loginPass := c.PostForm("LoginPass")
	token, err := loginVerification(loginEmail, loginPass)
	if err != nil {
		createCookie(token, c)
		//move to user page
	} else {
		log.Panicf("Incorrect User/Password %s", err.Error())
		//show error on the login popup
	}
}

func loginVerification(loginEmail, loginPass string) (string, error) {
	//get DB data based on email
	//if err - Login failed
	//compare db Passowrd to login Passowrd
	//if err - Login failed
	//get userID
	var uId uuid.UUID //get user ID from DB
	var role string   // get user role from DB
	validToken, err := generateJWT(loginEmail, uId, role)
	if err != nil {
		log.Panicln("Error with creating Token")
		return "", err
	}
	return validToken, nil

}
func generateJWT(email string, id uuid.UUID, role string) (string, error) {
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

	return tokenString, nil
}

func createCookie(token string, c *gin.Context) {
	c.SetCookie("AuthenticationCookie", token, 60*60*48, "", "", false, true)
	fmt.Println(c.Request.Cookie("AuthenticationCookie"))
}

func isAuthorized(c *gin.Context) {
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
}

func logout(c *gin.Context) {
	c.SetCookie("AuthenticationCookie", "expired", -1, "", "", false, false)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func main() {
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

	apiReq := router.Group("/apireq")
	{
		apiReq.GET("/csv", reinforstCSV)     //using Rainforst API for spesific Amazon product data to CSV
		apiReq.GET("/json", reinforstJSON)   //using Rainforest API for spesific Amazon Product data to JSON
		apiReq.POST("/search", searchResult) //using Axesso API for ASIN numbers from search results by Keyboard to JSON
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}

/*
IDEAS:
-monitoring all logs with Elsatic Search and represent number of issues / requests
-Side Quest : calculate Formula 1 pottential winning round for driver if possible
-side quest Email verefication API
*/
