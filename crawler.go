package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"

	"github.com/gocolly/colly/v2"
)

func getBsr(url string) []bsr {
	var bsr1 []bsr
	dataColector := colly.NewCollector(
		colly.AllowedDomains("amazon.com", "www.amazon.com"),
	)
	dataColector.OnRequest(func(r *colly.Request) {
		log.Printf("Colecting Data From : %s\n", r.URL)
	})
	dataColector.OnError(func(r *colly.Response, e error) {
		log.Printf("Error While Scrapping: %s\n", e.Error())
	})
	dataColector.OnHTML("ul.detail-bullet-list li", func(r *colly.HTMLElement) {
		selection := r.DOM
		chnode := selection.Children().Nodes
		for i := 0; i < len(chnode); i++ {
			if strings.Contains(selection.AddNodes(chnode[i]).Text(), "Best Sellers Rank") {
				str := strings.ReplaceAll(selection.AddNodes(chnode[i]).Text(), "  ", "")
				re := regexp.MustCompile(`\([^)]*\)`)
				str1 := re.ReplaceAllString(str, "")
				strf := strings.Split(str1, "#")
				// strf = removeDuplicateValues(removeDuplicateValues(strf))
				for i := 0; i < len(strf); i++ {
					strf[i] = strings.ReplaceAll(strf[i], "Best Sellers Rank:", "")
					var tempbsr bsr
					words := strings.Split(strf[i], " ")
					for _, word := range words {
						if len(word) > 0 && unicode.IsDigit(rune(word[0])) {
							tempbsr.Rank = strings.ReplaceAll(string(word), ",", "")
							tempbsr.Category = strf[i]
							bsr1 = append(bsr1, tempbsr)
						}
					}
				}
			}
		}
	})
	dataColector.OnHTML("tbody tr td span span", func(r *colly.HTMLElement) {
		if strings.HasPrefix(r.Text, "#") {
			re := regexp.MustCompile(`\([^)]*\)`)
			str := re.ReplaceAllString(r.Text, "")
			str = strings.ReplaceAll(str, "#", "")
			fmt.Printf("prod bsr: %s\n", str)
			str1 := strings.Split(str, " ")
			for i := 0; i < len(str1); i++ {
				var tempbsr bsr
				if len(str1[i]) > 0 && unicode.IsDigit(rune(str1[i][0])) {
					tempbsr.Rank = strings.ReplaceAll(str1[i], ",", "")
					tempbsr.Category = str
					tempbsr.Link = ""
					bsr1 = append(bsr1, tempbsr)
				}
			}

		}
	})
	dataColector.Visit(url)
	return removeDuplicateValues(bsr1)
}

func removeDuplicateValues(bsr1 []bsr) []bsr {
	keys := make(map[string]bool)
	list := []bsr{}
	for _, entry := range bsr1 {
		entry.Category = strings.TrimSpace(entry.Category)
		if _, value := keys[entry.Category]; !value {
			keys[entry.Category] = true
			list = append(list, entry)
		}
	}
	return list
}

// func updateBSR() {
// 	keys, err := keywordCache.Keys("*").Result()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	for _, key := range keys {
// 		url, err := url.Parse("https://api.rainforestapi.com/request?api_key=43BBBEBB52D3470598EFDC8F16AA6B60&type=search&amazon_domain=amazon.com&search_term=memory+cards&output=json")
// 		if err != nil {
// 			log.Printf("error")
// 		}
// 		values := url.Query()
// 		values.Set("search_term", key)

// 		url.RawQuery = values.Encode()
// 		prod := getProdListFromApi(url.String(), key)
// 		fmt.Println(prod)
// 	}
// }
