package main

import "time"

type Product struct {
	RequestInfo struct {
		Success                bool `json:"success"`
		CreditsUsed            int  `json:"credits_used"`
		CreditsRemaining       int  `json:"credits_remaining"`
		CreditsUsedThisRequest int  `json:"credits_used_this_request"`
	} `json:"request_info"`
	RequestParameters struct {
		AmazonDomain string `json:"amazon_domain"`
		Asin         string `json:"asin"`
		Type         string `json:"type"`
		Output       string `json:"output"`
	} `json:"request_parameters"`
	RequestMetadata struct {
		CreatedAt      time.Time `json:"created_at"`
		ProcessedAt    time.Time `json:"processed_at"`
		TotalTimeTaken float64   `json:"total_time_taken"`
		AmazonURL      string    `json:"amazon_url"`
	} `json:"request_metadata"`
	Product struct {
		Title       string `json:"title"`
		SearchAlias struct {
			Title string `json:"title"`
			Value string `json:"value"`
		} `json:"search_alias"`
		Keywords        string   `json:"keywords"`
		KeywordsList    []string `json:"keywords_list"`
		Asin            string   `json:"asin"`
		ParentAsin      string   `json:"parent_asin"`
		Link            string   `json:"link"`
		Brand           string   `json:"brand"`
		ProtectionPlans []struct {
			Asin  string `json:"asin"`
			Title string `json:"title"`
			Price struct {
				Symbol   string  `json:"symbol"`
				Value    float64 `json:"value"`
				Currency string  `json:"currency"`
				Raw      string  `json:"raw"`
			} `json:"price"`
		} `json:"protection_plans"`
		AddAnAccessory []struct {
			Asin  string `json:"asin"`
			Title string `json:"title"`
			Price struct {
				Symbol   string  `json:"symbol"`
				Value    float64 `json:"value"`
				Currency string  `json:"currency"`
				Raw      string  `json:"raw"`
			} `json:"price"`
		} `json:"add_an_accessory"`
		SellOnAmazon bool `json:"sell_on_amazon"`
		Variants     []struct {
			Asin             string `json:"asin"`
			Title            string `json:"title"`
			IsCurrentProduct bool   `json:"is_current_product"`
			Link             string `json:"link"`
			Dimensions       []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"dimensions"`
			MainImage string `json:"main_image"`
			Images    []struct {
				Variant string `json:"variant"`
				Link    string `json:"link"`
			} `json:"images"`
		} `json:"variants"`
		VariantAsinsFlat string `json:"variant_asins_flat"`
		Documents        []struct {
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"documents"`
		Categories []struct {
			Name       string `json:"name"`
			Link       string `json:"link"`
			CategoryID string `json:"category_id"`
		} `json:"categories"`
		CategoriesFlat string `json:"categories_flat"`
		Description    string `json:"description"`
		APlusContent   struct {
			HasAPlusContent bool `json:"has_a_plus_content"`
			HasBrandStory   bool `json:"has_brand_story"`
			ThirdParty      bool `json:"third_party"`
		} `json:"a_plus_content"`
		SubTitle struct {
			Text string `json:"text"`
			Link string `json:"link"`
		} `json:"sub_title"`
		Rating          float64 `json:"rating"`
		RatingBreakdown struct {
			FiveStar struct {
				Percentage int `json:"percentage"`
				Count      int `json:"count"`
			} `json:"five_star"`
			FourStar struct {
				Percentage int `json:"percentage"`
				Count      int `json:"count"`
			} `json:"four_star"`
			ThreeStar struct {
				Percentage int `json:"percentage"`
				Count      int `json:"count"`
			} `json:"three_star"`
			TwoStar struct {
				Percentage int `json:"percentage"`
				Count      int `json:"count"`
			} `json:"two_star"`
			OneStar struct {
				Percentage int `json:"percentage"`
				Count      int `json:"count"`
			} `json:"one_star"`
		} `json:"rating_breakdown"`
		RatingsTotal int `json:"ratings_total"`
		MainImage    struct {
			Link string `json:"link"`
		} `json:"main_image"`
		Images []struct {
			Link    string `json:"link"`
			Variant string `json:"variant"`
		} `json:"images"`
		ImagesCount int    `json:"images_count"`
		ImagesFlat  string `json:"images_flat"`
		Videos      []struct {
			DurationSeconds int    `json:"duration_seconds"`
			Width           int    `json:"width"`
			Height          int    `json:"height"`
			Link            string `json:"link"`
			Thumbnail       string `json:"thumbnail"`
			IsHeroVideo     bool   `json:"is_hero_video"`
			Variant         string `json:"variant"`
			GroupID         string `json:"group_id"`
			GroupType       string `json:"group_type"`
			Title           string `json:"title"`
		} `json:"videos"`
		VideosCount      int    `json:"videos_count"`
		VideosFlat       string `json:"videos_flat"`
		VideosAdditional []struct {
			ID                     string `json:"id"`
			ProductAsin            string `json:"product_asin"`
			ParentAsin             string `json:"parent_asin"`
			RelatedProducts        string `json:"related_products,omitempty"`
			SponsorProducts        string `json:"sponsor_products"`
			Title                  string `json:"title"`
			ProfileImageURL        string `json:"profile_image_url,omitempty"`
			ProfileLink            string `json:"profile_link,omitempty"`
			PublicName             string `json:"public_name"`
			CreatorType            string `json:"creator_type,omitempty"`
			VendorCode             string `json:"vendor_code"`
			VendorName             string `json:"vendor_name"`
			VideoImageID           string `json:"video_image_id"`
			VideoImageURL          string `json:"video_image_url"`
			VideoImageURLUnchanged string `json:"video_image_url_unchanged"`
			VideoImageWidth        string `json:"video_image_width"`
			VideoImageHeight       string `json:"video_image_height"`
			VideoImageExtension    string `json:"video_image_extension"`
			VideoURL               string `json:"video_url"`
			VideoPreviews          string `json:"video_previews"`
			VideoMimeType          string `json:"video_mime_type"`
			Duration               string `json:"duration"`
			Type                   string `json:"type"`
			ClosedCaptions         string `json:"closed_captions,omitempty"`
			VendorTrackingID       string `json:"vendor_tracking_id,omitempty"`
		} `json:"videos_additional"`
		IsBundle            bool     `json:"is_bundle"`
		FeatureBullets      []string `json:"feature_bullets"`
		FeatureBulletsCount int      `json:"feature_bullets_count"`
		FeatureBulletsFlat  string   `json:"feature_bullets_flat"`
		Attributes          []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"attributes"`
		TopReviews []struct {
			ID       string `json:"id"`
			Title    string `json:"title"`
			Body     string `json:"body"`
			BodyHTML string `json:"body_html"`
			Link     string `json:"link,omitempty"`
			Rating   int    `json:"rating"`
			Date     struct {
				Raw string    `json:"raw"`
				Utc time.Time `json:"utc"`
			} `json:"date"`
			Profile struct {
				Name string `json:"name"`
				Link string `json:"link"`
				ID   string `json:"id"`
			} `json:"profile,omitempty"`
			VineProgram      bool   `json:"vine_program"`
			VerifiedPurchase bool   `json:"verified_purchase"`
			ReviewCountry    string `json:"review_country"`
			IsGlobalReview   bool   `json:"is_global_review"`
			HelpfulVotes     int    `json:"helpful_votes,omitempty"`
			Profile0         struct {
				Name string `json:"name"`
			} `json:"profile1,omitempty"`
			Profile1 struct {
				Name string `json:"name"`
			} `json:"profile2,omitempty"`
			Profile2 struct {
				Name string `json:"name"`
			} `json:"profile3,omitempty"`
			Profile3 struct {
				Name  string `json:"name"`
				Image string `json:"image"`
			} `json:"profile4,omitempty"`
			Profile4 struct {
				Name string `json:"name"`
			} `json:"profile5,omitempty"`
		} `json:"top_reviews"`
		BuyboxWinner struct {
			MaximumOrderQuantity struct {
				Value       int  `json:"value"`
				HardMaximum bool `json:"hard_maximum"`
			} `json:"maximum_order_quantity"`
			OfferID          string `json:"offer_id"`
			MixedOffersCount int    `json:"mixed_offers_count"`
			MixedOffersFrom  struct {
				Symbol   string  `json:"symbol"`
				Value    float64 `json:"value"`
				Currency string  `json:"currency"`
				Raw      string  `json:"raw"`
			} `json:"mixed_offers_from"`
			IsPrime       bool `json:"is_prime"`
			IsAmazonFresh bool `json:"is_amazon_fresh"`
			Condition     struct {
				IsNew bool `json:"is_new"`
			} `json:"condition"`
			Availability struct {
				Raw          string `json:"raw"`
				DispatchDays int    `json:"dispatch_days"`
			} `json:"availability"`
			Fulfillment struct {
				Type             string `json:"type"`
				StandardDelivery struct {
					Date string `json:"date"`
					Name string `json:"name"`
				} `json:"standard_delivery"`
				FastestDelivery struct {
					Date string `json:"date"`
					Name string `json:"name"`
				} `json:"fastest_delivery"`
				IsSoldByAmazon          bool `json:"is_sold_by_amazon"`
				IsFulfilledByAmazon     bool `json:"is_fulfilled_by_amazon"`
				IsFulfilledByThirdParty bool `json:"is_fulfilled_by_third_party"`
				IsSoldByThirdParty      bool `json:"is_sold_by_third_party"`
				ThirdPartySeller        struct {
					Name string `json:"name"`
					Link string `json:"link"`
					ID   string `json:"id"`
				} `json:"third_party_seller"`
			} `json:"fulfillment"`
			Price struct {
				Symbol   string  `json:"symbol"`
				Value    float64 `json:"value"`
				Currency string  `json:"currency"`
				Raw      string  `json:"raw"`
			} `json:"price"`
			Rrp struct {
				Symbol   string  `json:"symbol"`
				Value    float64 `json:"value"`
				Currency string  `json:"currency"`
				Raw      string  `json:"raw"`
			} `json:"rrp"`
			Shipping struct {
				Raw string `json:"raw"`
			} `json:"shipping"`
		} `json:"buybox_winner"`
		MoreBuyingChoices []struct {
			Price struct {
				Symbol   string  `json:"symbol"`
				Value    float64 `json:"value"`
				Currency string  `json:"currency"`
				Raw      string  `json:"raw"`
			} `json:"price"`
			SellerName   string `json:"seller_name"`
			SellerLink   string `json:"seller_link"`
			FreeShipping bool   `json:"free_shipping"`
			Position     int    `json:"position"`
		} `json:"more_buying_choices"`
		Specifications []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"specifications"`
		SpecificationsFlat string `json:"specifications_flat"`
		BestsellersRank    []struct {
			Category string `json:"category"`
			Rank     int    `json:"rank"`
			Link     string `json:"link"`
		} `json:"bestsellers_rank"`
		Manufacturer        string   `json:"manufacturer"`
		Weight              string   `json:"weight"`
		FirstAvailable      string   `json:"first_available"`
		Dimensions          string   `json:"dimensions"`
		ModelNumber         string   `json:"model_number"`
		BestsellersRankFlat string   `json:"bestsellers_rank_flat"`
		WhatsInTheBox       []string `json:"whats_in_the_box"`
	} `json:"product"`
	BrandStore struct {
		ID   string `json:"id"`
		Link string `json:"link"`
	} `json:"brand_store"`
	UserGuide  string `json:"user_guide"`
	NewerModel struct {
		Title        string  `json:"title"`
		Asin         string  `json:"asin"`
		Link         string  `json:"link"`
		Image        string  `json:"image"`
		Rating       float64 `json:"rating"`
		RatingsTotal int     `json:"ratings_total"`
		Price        struct {
			Symbol   string  `json:"symbol"`
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
			Raw      string  `json:"raw"`
		} `json:"price"`
	} `json:"newer_model"`
	FrequentlyBoughtTogether struct {
		TotalPrice struct {
			Symbol   string  `json:"symbol"`
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
			Raw      string  `json:"raw"`
		} `json:"total_price"`
		Products []struct {
			Asin  string `json:"asin"`
			Title string `json:"title"`
			Link  string `json:"link"`
			Image string `json:"image"`
			Price struct {
				Symbol   string  `json:"symbol"`
				Value    float64 `json:"value"`
				Currency string  `json:"currency"`
				Raw      string  `json:"raw"`
			} `json:"price"`
		} `json:"products"`
	} `json:"frequently_bought_together"`
	CompareWithSimilar []struct {
		Asin         string  `json:"asin"`
		Image        string  `json:"image"`
		Title        string  `json:"title"`
		Rating       float64 `json:"rating"`
		RatingsTotal int     `json:"ratings_total"`
		Price        struct {
			Symbol   string  `json:"symbol"`
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
			Raw      string  `json:"raw"`
		} `json:"price"`
		Link string `json:"link"`
	} `json:"compare_with_similar"`
}

type EstimatedSales struct {
	RequestInfo struct {
		Success          bool `json:"success"`
		CreditsUsed      int  `json:"credits_used"`
		CreditsRemaining int  `json:"credits_remaining"`
	} `json:"request_info"`
	RequestMetadata struct {
		ID             string    `json:"id"`
		CreatedAt      time.Time `json:"created_at"`
		ProcessedAt    time.Time `json:"processed_at"`
		TotalTimeTaken float64   `json:"total_time_taken"`
	} `json:"request_metadata"`
	RequestParameters struct {
		Type         string `json:"type"`
		Asin         string `json:"asin"`
		AmazonDomain string `json:"amazon_domain"`
	} `json:"request_parameters"`
	SalesEstimation struct {
		HasSalesEstimation      bool   `json:"has_sales_estimation"`
		MonthlySalesEstimate    int    `json:"monthly_sales_estimate"`
		WeeklySalesEstimate     int    `json:"weekly_sales_estimate"`
		BestsellerRank          int    `json:"bestseller_rank"`
		SalesEstimationCategory string `json:"sales_estimation_category"`
	} `json:"sales_estimation"`
}

type searchResult struct {
	RequestInfo struct {
		Success                bool `json:"success"`
		CreditsUsed            int  `json:"credits_used"`
		CreditsRemaining       int  `json:"credits_remaining"`
		CreditsUsedThisRequest int  `json:"credits_used_this_request"`
	} `json:"request_info"`
	RequestParameters struct {
		Type         string `json:"type"`
		AmazonDomain string `json:"amazon_domain"`
		SearchTerm   string `json:"search_term"`
	} `json:"request_parameters"`
	RequestMetadata struct {
		CreatedAt      time.Time `json:"created_at"`
		ProcessedAt    time.Time `json:"processed_at"`
		TotalTimeTaken float64   `json:"total_time_taken"`
		AmazonURL      string    `json:"amazon_url"`
	} `json:"request_metadata"`
	SearchResults []struct {
		Position     int    `json:"position"`
		Title        string `json:"title"`
		Asin         string `json:"asin"`
		Link         string `json:"link"`
		Availability struct {
			Raw string `json:"raw"`
		} `json:"availability,omitempty"`
		Categories []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"categories"`
		Image        string  `json:"image"`
		Rating       float64 `json:"rating"`
		RatingsTotal int     `json:"ratings_total"`
		Prices       []struct {
			Symbol    string  `json:"symbol"`
			Value     float64 `json:"value"`
			Currency  string  `json:"currency"`
			Raw       string  `json:"raw"`
			Name      string  `json:"name"`
			IsPrimary bool    `json:"is_primary,omitempty"`
			Asin      string  `json:"asin,omitempty"`
			Link      string  `json:"link,omitempty"`
			IsRrp     bool    `json:"is_rrp,omitempty"`
		} `json:"prices,omitempty"`
		Price struct {
			Symbol    string  `json:"symbol"`
			Value     float64 `json:"value"`
			Currency  string  `json:"currency"`
			Raw       string  `json:"raw"`
			Name      string  `json:"name"`
			IsPrimary bool    `json:"is_primary"`
			Asin      string  `json:"asin"`
			Link      string  `json:"link"`
		} `json:"price,omitempty"`
	} `json:"search_results"`
	CategoryInformation struct {
		IsLandingPage bool `json:"is_landing_page"`
	} `json:"category_information"`
	RelatedSearches []struct {
		Query string `json:"query"`
		Link  string `json:"link"`
	} `json:"related_searches"`
	RelatedBrands []struct {
		Logo      string `json:"logo"`
		Image     string `json:"image"`
		Title     string `json:"title"`
		Link      string `json:"link"`
		StoreLink string `json:"store_link"`
		StoreID   string `json:"store_id"`
		StoreName string `json:"store_name"`
	} `json:"related_brands"`
	Pagination struct {
		TotalResults int    `json:"total_results"`
		CurrentPage  int    `json:"current_page"`
		NextPageLink string `json:"next_page_link"`
		TotalPages   int    `json:"total_pages"`
	} `json:"pagination"`
	Refinements struct {
		Prime []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"prime"`
		Delivery []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"delivery"`
		Departments []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"departments"`
		Reviews []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"reviews"`
		Price []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link,omitempty"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"price"`
		Brand []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"brand"`
		NewReleases []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"new_releases"`
		VideoGameConsoleWirelessCommunicationTechnology []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"video_game_console_wireless_communication_technology"`
		PackagingOption []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"packaging_option"`
		AmazonGlobalStore []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"amazon_global_store"`
		InternationalShipping []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"international_shipping"`
		Condition []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"condition"`
		Availability []struct {
			Name                  string `json:"name"`
			Value                 string `json:"value"`
			Link                  string `json:"link"`
			RefinementDisplayName string `json:"refinement_display_name"`
		} `json:"availability"`
	} `json:"refinements"`
	ShoppingAdvisors []struct {
		Position        int    `json:"position"`
		Title           string `json:"title"`
		Recommendations []struct {
			Position int `json:"position"`
			Product  struct {
				Title        string  `json:"title"`
				Asin         string  `json:"asin"`
				Link         string  `json:"link"`
				Image        string  `json:"image"`
				Rating       float64 `json:"rating"`
				RatingsTotal int     `json:"ratings_total"`
				Price        struct {
					Value    float64 `json:"value"`
					Currency string  `json:"currency"`
					Symbol   string  `json:"symbol"`
					Raw      string  `json:"raw"`
				} `json:"price"`
			} `json:"product"`
		} `json:"recommendations"`
	} `json:"shopping_advisors"`
	AdBlocks []struct {
		CampaignID      string `json:"campaign_id"`
		BrandLogo       string `json:"brand_logo"`
		BackgroundImage string `json:"background_image"`
		AdvertiserID    string `json:"advertiser_id"`
		AdID            string `json:"ad_id"`
		Link            string `json:"link"`
		Title           string `json:"title"`
		StoreLink       string `json:"store_link"`
		StoreID         string `json:"store_id"`
		StoreName       string `json:"store_name"`
		Products        []struct {
			Asin         string `json:"asin"`
			Link         string `json:"link"`
			Image        string `json:"image"`
			Title        string `json:"title"`
			IsPrime      bool   `json:"is_prime,omitempty"`
			Rating       int    `json:"rating"`
			RatingsTotal int    `json:"ratings_total"`
		} `json:"products"`
	} `json:"ad_blocks"`
	VideoBlocks []struct {
		VideoLink     string `json:"video_link"`
		ThumbnailLink string `json:"thumbnail_link"`
		CampaignID    string `json:"campaign_id"`
		AdvertiserID  string `json:"advertiser_id"`
		HasAudio      bool   `json:"has_audio"`
		Products      []struct {
			Asin         string  `json:"asin"`
			Link         string  `json:"link"`
			Image        string  `json:"image"`
			Title        string  `json:"title"`
			IsPrime      bool    `json:"is_prime"`
			Rating       float64 `json:"rating"`
			RatingsTotal int     `json:"ratings_total"`
			Price        struct {
				Value    float64 `json:"value"`
				Currency string  `json:"currency"`
				Symbol   string  `json:"symbol"`
				Raw      string  `json:"raw"`
			} `json:"price"`
		} `json:"products"`
	} `json:"video_blocks"`
}

var NewProduct Product
