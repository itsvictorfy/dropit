package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

//var SProducts []product

func searchNew(c *gin.Context) {
	var sProducts []product
	//getProducts(searchQuery(c))
	//if (i can get estimated sales)
	for i, prod := range sProducts {
		sProducts[i].MonthlySalesEstimate = getEstimatedsales(editQuery(prod))
		sProducts[i].MonthlyRevenueEstimate = sProducts[i].MonthlySalesEstimate * sProducts[i].Price
	}

}

// sales estimation per ASIN
func getEstimatedsales(url string) float64 {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("new request error %s", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("error getting response from %s: \n %v", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("request failed :\n %d", resp.StatusCode)
	}
	var salesEstimation EstimatedSales
	json.NewDecoder(resp.Body).Decode(&salesEstimation)
	return float64(salesEstimation.SalesEstimation.MonthlySalesEstimate)

}

// get Asins from Search
func getProducts(url string) product {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("new request error %s", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("error getting response from %s: \n %v", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("request failed :\n %d", resp.StatusCode)
	}
	var searchReq searchResult
	json.NewDecoder(resp.Body).Decode(&searchReq)
	fmt.Println(resp.Body)

	b, err := json.MarshalIndent(searchReq, "", " ")
	if err != nil {
		log.Printf("error w/ Marshelindent :\n %s", err)
	}
	log.Printf("JSON:\n %s", b)
	var searchProducts []product
	for i := 0; i < len(searchReq.SearchResults); i++ {
		var prod product
		prod.Asin = searchReq.SearchResults[i].Asin
		prod.ImageUrl = searchReq.SearchResults[i].Image
		prod.Link = searchReq.SearchResults[i].Link
		prod.Price = searchReq.SearchResults[i].Price.Value
		prod.Rating = searchReq.SearchResults[i].Rating
		prod.TotalRatings = searchReq.SearchResults[i].RatingsTotal
		prod.Title = searchReq.SearchResults[i].Title
		searchProducts = append(searchProducts, prod)
		cachedProd, _ := json.Marshal(prod)
		productCache.Set(prod.Asin, cachedProd, 5)
	}
	fmt.Println(searchProducts)
	return searchProducts[1]
}

// edit Search Query for Get Products
func searchQuery(c *gin.Context) string {
	url, err := url.Parse("https://api.rainforestapi.com/request?api_key=63FFD605C840421D9F5FC4433C106F90&type=search&amazon_domain=amazon.com&search_term=xbox1&output=json")
	if err != nil {
		log.Printf("error")
	}
	fmt.Printf("%s \n", url.String())

	values := url.Query()
	keyword := c.PostForm("keyword")
	values.Set("search_term", keyword)

	url.RawQuery = values.Encode()
	fmt.Printf("test:%s \n", url)
	return string(url.String())
}

// edit Api Query to get Estimated sales data per product
func editQuery(prod product) string {
	url, err := url.Parse("https://api.rainforestapi.com/request?api_key=63FFD605C840421D9F5FC4433C106F90&type=sales_estimation&amazon_domain=amazon.com&asin=B07GY4DS42&output=json")
	if err != nil {
		log.Printf("error")
	}
	fmt.Printf("%s \n", url.String())

	values := url.Query()
	asin := prod.Asin
	values.Set("asin", asin)

	url.RawQuery = values.Encode()
	fmt.Printf("test:%s \n", url)
	return string(url.String())
}
