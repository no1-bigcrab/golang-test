package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-test/serve/api/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Product is
type Product struct {
	Links struct {
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Attributes []struct {
		ID        int      `json:"id"`
		Name      string   `json:"name"`
		Options   []string `json:"options"`
		Position  int      `json:"position"`
		Variation bool     `json:"variation"`
		Visible   bool     `json:"visible"`
	} `json:"attributes"`
	AverageRating     string `json:"average_rating"`
	Backordered       bool   `json:"backordered"`
	Backorders        string `json:"backorders"`
	BackordersAllowed bool   `json:"backorders_allowed"`
	ButtonText        string `json:"button_text"`
	CatalogVisibility string `json:"catalog_visibility"`
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
	RelatedIds       []int         `json:"related_ids"`
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
	Variations       []int         `json:"variations"`
	Virtual          bool          `json:"virtual"`
	Weight           string        `json:"weight"`
}

//ProductsPageGet in function
func (a *App) ProductsPageGet(w http.ResponseWriter, r *http.Request) {
	url := "http://dev.local/wp-json/wc/v3/products/12"

	body := sendRequest(url)
	var data Product
	json.Unmarshal(body, &data)

	//for i := 0; i < len(data); i++ {

	//check data Atribute nếu khác rỗng. thì duyệt
	variantsData := []interface{}{}

	for i := 0; i < len(data.Variations); i++ {
		//nên tạo 1 function call api.có các trường có thể nhập vàofor k,v := range {
		url1 := "http://dev.local/wp-json/wc/v3/products/12" + "/variations/" + strconv.Itoa(data.Variations[i])

		body1 := sendRequest(url1)
		var data1 models.Products
		json.Unmarshal(body1, &data1)
		// jsonValue, _ := json.Marshal(data1)
		// fmt.Println(bytes.NewBuffer(jsonValue))
		dataAttr := map[string]interface{}{
			"id":  data.Variations[i],
			"src": data1.Image.Src,
		}
		varData := map[string]interface{}{
			"option1": data1.Attributes[0].Option,
			"price":   data1.Price,
			"sku":     data1.Sku,
			"image":   dataAttr,
		}

		variantsData = append(variantsData, varData)
	}

	optionsData := []interface{}{}
	imgData := []interface{}{}

	for i := 0; i < len(data.Images); i++ {
		dataImage := map[string]string{
			"alt": "",
			"src": data.Images[i].Src,
		}
		imgData = append(imgData, dataImage)
	}

	values := map[string]map[string]interface{}{
		"product": {
			"title":       data.Name,
			"body_html":   data.Description,
			"vendor":      "",
			"productType": "",
			"variants":    variantsData,
			"options":     optionsData,
			"images":      imgData,
			"image":       "",
		},
	}
	jsonValue, _ := json.Marshal(values)
	fmt.Println(bytes.NewBuffer(jsonValue))
	url2 := "https://c8f4666a96a5f2dce771c1c04a427308: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@bigcrab-1.myshopify.com/admin/api/2020-04/products.json"

	checkValueRequest(values, url2)
}
func checkValue(values map[string]map[string]interface{}, url string) {

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

func sendRequest(url string) []byte {
	token1 := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9kZXYubG9jYWwiLCJpYXQiOjE1OTEyNTYxMTksIm5iZiI6MTU5MTI1NjExOSwiZXhwIjoxNTkxODYwOTE5LCJkYXRhIjp7InVzZXIiOnsiaWQiOjEsInR5cGUiOiJ3cF91c2VyIiwidXNlcl9sb2dpbiI6ImFkbWluIiwidXNlcl9lbWFpbCI6ImRldi1lbWFpbEBmbHl3aGVlbC5sb2NhbCIsImFwaV9rZXkiOiIxaUoxNHJydmFjbEZCTjVBSmpiMUpDa2ZFIn19fQ.2Pmo5DefSc69txjGtAJX8zwU7Oxw9rS_V7_wm5ARiYg"

	r, err1 := http.NewRequest("GET", url, nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+token1)
	client1 := &http.Client{}
	resp1, err1 := client1.Do(r)
	if err1 != nil {
		panic(err1)
	}
	defer resp1.Body.Close()
	body1, _ := ioutil.ReadAll(resp1.Body)

	return body1
}
