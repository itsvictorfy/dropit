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
func getProducts(url string, keyword string) []product { //fix : if BSR in DB no need to getBSR
	kwexistsinDB, _ := keywordCache.Exists(keyword).Result()
	if kwexistsinDB == 1 {
		log.Printf("Found Privious search in Database")
		return getProdListFromDB(keyword)
	} else {
		return getProdListFromApi(url, keyword)
	}
}

// edit Search Query for Get Products
func searchQuery(c *gin.Context) (string, string) {
	url, err := url.Parse("https://api.rainforestapi.com/request?api_key=43BBBEBB52D3470598EFDC8F16AA6B60&type=search&amazon_domain=amazon.com&search_term=memory+cards&output=json")
	if err != nil {
		log.Printf("error")
	}
	values := url.Query()
	keyword := c.PostForm("keyword")
	values.Set("search_term", keyword)

	url.RawQuery = values.Encode()
	log.Printf("Getting Data from :%s \n", url)
	return string(url.String()), keyword
}

func getProdListFromApi(url, keyword string) []product {
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

	var searchProducts []product
	var asinlist keywordAsins
	asinlist.Keyword = keyword
	for i := 0; i < len(searchReq.SearchResults); i++ {
		var prod product
		url1 := "https://www.amazon.com/dp/" + searchReq.SearchResults[i].Asin
		prod.Asin = searchReq.SearchResults[i].Asin
		prod.ImageUrl = searchReq.SearchResults[i].Image
		prod.Link = searchReq.SearchResults[i].Link
		prod.Price = searchReq.SearchResults[i].Price.Value
		prod.Rating = searchReq.SearchResults[i].Rating
		prod.TotalRatings = searchReq.SearchResults[i].RatingsTotal
		prod.Title = searchReq.SearchResults[i].Title
		prod.BestsellersRank = getBsr(url1)
		searchProducts = append(searchProducts, prod)
		asinlist.Asins = append(asinlist.Asins, prod.Asin)
		cachedProd, _ := json.Marshal(prod)
		productCache.Set(prod.Asin, cachedProd, 0)
	}
	cachedkeyword, _ := json.Marshal(asinlist)
	keywordCache.Set(asinlist.Keyword, cachedkeyword, 0)
	log.Printf("Get Products Completed")
	return searchProducts
}

func getProdListFromDB(keyword string) []product {
	var prodlist []product
	kwdbvalue, _ := keywordCache.Get(keyword).Result()
	var asinlist keywordAsins
	err := json.Unmarshal([]byte(kwdbvalue), &asinlist)
	if err != nil {
		log.Printf("error unmarshalling : %s", err)
	} else {
		for i := 0; i < len(asinlist.Asins); i++ {
			var prod product
			proddbvalue, _ := productCache.Get(asinlist.Asins[i]).Result()
			err := json.Unmarshal([]byte(proddbvalue), &prod)
			if err != nil {
				log.Printf("error unmarshalling : %s", err)
			}
			if prod.BestsellersRank == nil {
				fmt.Printf("BSR from for: %s ", prod.BestsellersRank)
				url := "https://www.amazon.com/dp/" + prod.Asin
				bsr := getBsr(url)
				prod.BestsellersRank = append(prod.BestsellersRank, bsr...)
				cachedProd, _ := json.Marshal(prod)
				productCache.Set(prod.Asin, cachedProd, 0)
			}
			prodlist = append(prodlist, prod)
		}
	}
	return prodlist
}

// edit Api Query to get Estimated sales data per product
func editQuery(prod product) string {
	url, err := url.Parse("https://api.rainforestapi.com/request?api_key=63FFD605C840421D9F5FC4433C106F90&type=sales_estimation&amazon_domain=amazon.com&asin=B07GY4DS42&output=json")
	if err != nil {
		log.Printf("error")
	}
	log.Printf("%s \n", url.String())

	values := url.Query()
	asin := prod.Asin
	values.Set("asin", asin)

	url.RawQuery = values.Encode()
	log.Printf("test:%s \n", url)
	return string(url.String())
}
