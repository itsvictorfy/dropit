package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

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
		fmt.Println(err.Error())
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
