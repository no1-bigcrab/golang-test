package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-test/serve/api/urls"
	"html/template"
	"io/ioutil"
	"net/http"
)

//Products ahhhu
type Products []struct {
	Links struct {
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Attributes        []interface{} `json:"attributes"`
	AverageRating     string        `json:"average_rating"`
	Backordered       bool          `json:"backordered"`
	Backorders        string        `json:"backorders"`
	BackordersAllowed bool          `json:"backorders_allowed"`
	ButtonText        string        `json:"button_text"`
	CatalogVisibility string        `json:"catalog_visibility"`
	Categories        []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"categories"`
	CrossSellIds      []interface{} `json:"cross_sell_ids"`
	DateOnSaleFrom    interface{}   `json:"date_on_sale_from"`
	DateOnSaleFromGmt interface{}   `json:"date_on_sale_from_gmt"`
	DateOnSaleTo      interface{}   `json:"date_on_sale_to"`
	DateOnSaleToGmt   interface{}   `json:"date_on_sale_to_gmt"`
	DefaultAttributes []interface{} `json:"default_attributes"`
	Description       string        `json:"description"`
	Dimensions        struct {
		Height string `json:"height"`
		Length string `json:"length"`
		Width  string `json:"width"`
	} `json:"dimensions"`
	DownloadExpiry  int           `json:"download_expiry"`
	DownloadLimit   int           `json:"download_limit"`
	Downloadable    bool          `json:"downloadable"`
	Downloads       []interface{} `json:"downloads"`
	ExternalURL     string        `json:"external_url"`
	Featured        bool          `json:"featured"`
	GroupedProducts []interface{} `json:"grouped_products"`
	ID              int           `json:"id"`
	Images          []struct {
		Alt  string `json:"alt"`
		ID   int    `json:"id"`
		Name string `json:"name"`
		Src  string `json:"src"`
	} `json:"images"`
	ManageStock bool `json:"manage_stock"`
	MenuOrder   int  `json:"menu_order"`
	MetaData    []struct {
		ID    int    `json:"id"`
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"meta_data"`
	Name             string        `json:"name"`
	OnSale           bool          `json:"on_sale"`
	ParentID         int           `json:"parent_id"`
	Permalink        string        `json:"permalink"`
	Price            string        `json:"price"`
	PriceHTML        string        `json:"price_html"`
	Purchasable      bool          `json:"purchasable"`
	PurchaseNote     string        `json:"purchase_note"`
	RatingCount      int           `json:"rating_count"`
	RegularPrice     string        `json:"regular_price"`
	RelatedIds       []interface{} `json:"related_ids"`
	ReviewsAllowed   bool          `json:"reviews_allowed"`
	SalePrice        string        `json:"sale_price"`
	ShippingClass    string        `json:"shipping_class"`
	ShippingClassID  int           `json:"shipping_class_id"`
	ShippingRequired bool          `json:"shipping_required"`
	ShippingTaxable  bool          `json:"shipping_taxable"`
	ShortDescription string        `json:"short_description"`
	Sku              string        `json:"sku"`
	Slug             string        `json:"slug"`
	SoldIndividually bool          `json:"sold_individually"`
	Status           string        `json:"status"`
	StockQuantity    interface{}   `json:"stock_quantity"`
	StockStatus      string        `json:"stock_status"`
	Tags             []interface{} `json:"tags"`
	TaxClass         string        `json:"tax_class"`
	TaxStatus        string        `json:"tax_status"`
	TotalSales       int           `json:"total_sales"`
	Type             string        `json:"type"`
	UpsellIds        []interface{} `json:"upsell_ids"`
	Variations       []interface{} `json:"variations"`
	Virtual          bool          `json:"virtual"`
	Weight           string        `json:"weight"`
	DateCreated      string        `json:"date_created,omitempty"`
	DateCreatedGmt   string        `json:"date_created_gmt,omitempty"`
	DateModified     string        `json:"date_modified,omitempty"`
	DateModifiedGmt  string        `json:"date_modified_gmt,omitempty"`
}

//HomePageGet in function
func (a *App) HomePageGet(w http.ResponseWriter, r *http.Request) {

	path := urls.PathUrl()

	url := "http://dev.local/wp-json/wc/v3/products"

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9kZXYubG9jYWwiLCJpYXQiOjE1OTEyNTYxMTksIm5iZiI6MTU5MTI1NjExOSwiZXhwIjoxNTkxODYwOTE5LCJkYXRhIjp7InVzZXIiOnsiaWQiOjEsInR5cGUiOiJ3cF91c2VyIiwidXNlcl9sb2dpbiI6ImFkbWluIiwidXNlcl9lbWFpbCI6ImRldi1lbWFpbEBmbHl3aGVlbC5sb2NhbCIsImFwaV9rZXkiOiIxaUoxNHJydmFjbEZCTjVBSmpiMUpDa2ZFIn19fQ.2Pmo5DefSc69txjGtAJX8zwU7Oxw9rS_V7_wm5ARiYg"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var data Products
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := template.Must(template.ParseFiles(path.TEMPLATE_PATH + "/form.html"))
	html.Execute(w, data)
}

//HomePagePost is func
func (a *App) HomePagePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := "https://" + r.FormValue("apikey") + ":" + r.FormValue("password") + "@" + r.FormValue("hostname") + "/admin/api/2020-04/" + r.FormValue("api-title") + ".json"
	title := r.FormValue("api-title")

	//fmt.Printf(r.ParseForm().Error())
	switch title {
	case "products":
		varData := map[string]interface{}{
			"option1": r.FormValue("variants-color"),
			"option2": r.FormValue("variants-size"),
		}

		opData := map[string]interface{}{
			"name":   r.FormValue("options-name"),
			"values": r.Form["select2_multi"],
		}

		variantsData := []interface{}{
			varData,
		}

		optionsData := []interface{}{
			opData,
		}

		values := map[string]map[string]interface{}{
			"product": {
				"title":       r.FormValue("title"),
				"body_html":   r.FormValue("html-body"),
				"vendor":      r.FormValue("vendor"),
				"productType": r.FormValue("product-type"),
				"variants":    variantsData,
				"options":     optionsData,
				"images":      "",
				"image":       "",
			},
		}

		//fmt.Println(values)
		checkValueRequest(values, url)

	case "pages":
		values := map[string]map[string]interface{}{
			"page": {
				"title":     r.FormValue("title"),
				"body_html": r.FormValue("html-body"),
				"author":    r.FormValue("author"),
			},
		}
		checkValueRequest(values, url)

	case "blogs":
		values := map[string]map[string]interface{}{
			"blog": {
				"title": r.FormValue("title"),
			},
		}
		checkValueRequest(values, url)

	case "article":
		values := map[string]map[string]interface{}{
			"article": {
				"title":     r.FormValue("title"),
				"body_html": r.FormValue("html-body"),
				"author":    r.FormValue("author"),
				"tags":      r.FormValue("tags"),
				"image":     r.FormValue("image"),
			},
		}
		checkValueRequest(values, url)

	default:
		json.NewEncoder(w).Encode("No created !!!")
	}
}

func checkValueRequest(values map[string]map[string]interface{}, url string) {

	jsonValue, _ := json.Marshal(values)
	//fmt.Println(bytes.NewBuffer(jsonValue))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
}
