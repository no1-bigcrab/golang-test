package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-test/serve/api/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//Product is
type Product []struct {
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
	Name             string      `json:"name"`
	OnSale           bool        `json:"on_sale"`
	ParentID         int         `json:"parent_id"`
	Permalink        string      `json:"permalink"`
	Price            string      `json:"price"`
	PriceHTML        string      `json:"price_html"`
	Purchasable      bool        `json:"purchasable"`
	PurchaseNote     string      `json:"purchase_note"`
	RatingCount      int         `json:"rating_count"`
	RegularPrice     string      `json:"regular_price"`
	RelatedIds       []int       `json:"related_ids"`
	ReviewsAllowed   bool        `json:"reviews_allowed"`
	SalePrice        string      `json:"sale_price"`
	ShippingClass    string      `json:"shipping_class"`
	ShippingClassID  int         `json:"shipping_class_id"`
	ShippingRequired bool        `json:"shipping_required"`
	ShippingTaxable  bool        `json:"shipping_taxable"`
	ShortDescription string      `json:"short_description"`
	Sku              string      `json:"sku"`
	Slug             string      `json:"slug"`
	SoldIndividually bool        `json:"sold_individually"`
	Status           string      `json:"status"`
	StockQuantity    interface{} `json:"stock_quantity"`
	StockStatus      string      `json:"stock_status"`
	Tags             []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}
	TaxClass   string        `json:"tax_class"`
	TaxStatus  string        `json:"tax_status"`
	TotalSales int           `json:"total_sales"`
	Type       string        `json:"type"`
	UpsellIds  []interface{} `json:"upsell_ids"`
	Variations []int         `json:"variations"`
	Virtual    bool          `json:"virtual"`
	Weight     string        `json:"weight"`
}

//CustomCollections is
type CustomCollections []struct {
	ID                int64       `json:"id"`
	Handle            string      `json:"handle"`
	Title             string      `json:"title"`
	UpdatedAt         string      `json:"updated_at"`
	BodyHTML          interface{} `json:"body_html"`
	PublishedAt       string      `json:"published_at"`
	SortOrder         string      `json:"sort_order"`
	TemplateSuffix    interface{} `json:"template_suffix"`
	PublishedScope    string      `json:"published_scope"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
}

//ProductsPagePost in function
func (a *App) ProductsPagePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postURL := r.FormValue("url")
	postDomain := r.FormValue("domain")
	postPassword := r.FormValue("password")
	postAPIkey := r.FormValue("apikey")
	for count := 1; count < 100; count++ {
		url := postURL + "products/?page=" + strconv.Itoa(count)

		body := getValueFromWp(url)
		var data Product
		json.Unmarshal(body, &data)
		if len(data) > 0 {
			for i := 0; i < len(data); i++ {

				variantsData := []interface{}{}
				optionsData := []interface{}{}
				imgData := []interface{}{}
				tagsData := []string{}
				var dataPrice string
				var dataRegPrice string
				collectData := []interface{}{}
				//check data Variations nếu khác rỗng. thì duyệt
				if len(data[i].Variations) > 0 {
					for i1 := 0; i1 < len(data[i].Variations); i1++ {
						url1 := postURL + "products/" + strconv.Itoa(data[i].ID) + "/variations/" + strconv.Itoa(data[i].Variations[i1])

						body1 := getValueFromWp(url1)
						var data1 models.Products
						json.Unmarshal(body1, &data1)
						dataAttr := map[string]interface{}{
							"src": data1.Image.Src,
						}
						if data1.Price == data1.RegularPrice || len(data1.RegularPrice) == 0 {
							dataRegPrice = ""
						} else {
							dataRegPrice = data1.RegularPrice
						}
						dataPrice = data1.Price

						if len(data1.Attributes) > 0 {
							if len(data1.Attributes) == 2 {
								varDatas := map[string]interface{}{
									"option1":          data1.Attributes[0].Option,
									"option2":          data1.Attributes[1].Option,
									"option3":          "",
									"price":            dataPrice,
									"sku":              data1.Sku,
									"image":            dataAttr,
									"compare_at_price": dataRegPrice,
								}
								variantsData = append(variantsData, varDatas)

							} else if len(data1.Attributes) == 1 {
								varDatas := map[string]interface{}{
									"option1":          data1.Attributes[0].Option,
									"option2":          "All Size",
									"option3":          "",
									"price":            dataPrice,
									"sku":              data1.Sku,
									"image":            dataAttr,
									"compare_at_price": dataRegPrice,
								}
								variantsData = append(variantsData, varDatas)

							} else if len(data1.Attributes) == 3 {
								varDatas := map[string]interface{}{
									"option1":          data1.Attributes[0].Option,
									"option2":          data1.Attributes[1].Option,
									"option3":          data1.Attributes[2].Option,
									"price":            dataPrice,
									"sku":              data1.Sku,
									"image":            dataAttr,
									"compare_at_price": dataRegPrice,
								}
								variantsData = append(variantsData, varDatas)

							}
						}
					}
				} else {
					dataPrice = data[i].Price
					imgData = append(imgData, data[i].Images)
					if data[i].Price == data[i].RegularPrice || len(data[i].RegularPrice) == 0 {
						dataRegPrice = ""
					} else {
						dataRegPrice = data[i].RegularPrice
					}
					varDatas := map[string]interface{}{
						"price":            dataPrice,
						"sku":              data[i].Sku,
						"compare_at_price": dataRegPrice,
					}
					variantsData = append(variantsData, varDatas)
				}
				if len(data[i].Attributes) > 0 {
					for countAttr := 0; countAttr < len(data[i].Attributes); countAttr++ {
						nameData := data[i].Attributes[countAttr].Name
						valueAttr := data[i].Attributes[countAttr].Options
						optDatas := map[string]interface{}{
							"name":   nameData,
							"values": valueAttr,
						}
						optionsData = append(optionsData, optDatas)
					}
				}

				if data[i].Images != nil {
					for i2 := 0; i2 < len(data[i].Images); i2++ {
						dataImage := map[string]string{
							"alt": "",
							"src": data[i].Images[i2].Src,
						}
						imgData = append(imgData, dataImage)
					}
				}

				if len(data[i].Tags) > 0 {
					for i3 := 0; i3 < len(data[i].Tags); i3++ {
						tags := data[i].Tags[i3].Name
						tagsData = append(tagsData, tags)
					}
				}
				if len(data[i].Categories) > 0 {
					for i4 := 0; i4 < len(data[i].Categories); i4++ {
						dataCategory := data[i].Categories[i4]
						IDCollection := createCollection(dataCategory.Name, postURL, postDomain, postPassword, postAPIkey)
						//fmt.Println(IDCollection)
						collectData = append(collectData, strconv.FormatInt(IDCollection, 10))
					}
				}
				values := map[string]map[string]interface{}{
					"product": {
						"title":       data[i].Name,
						"body_html":   data[i].Description,
						"vendor":      "",
						"productType": "",
						"variants":    variantsData,
						"options":     optionsData,
						"images":      imgData,
						"image":       "",
						"tags":        tagsData,
					},
				}
				// jsonValue, _ := json.Marshal(values)
				// fmt.Println(bytes.NewBuffer(jsonValue))
				url2 := "https://" + postAPIkey + ":" + postPassword + "@" + postDomain + "/admin/api/2020-04/" + r.FormValue("api-title") + ".json"
				nameProduct := strings.ReplaceAll(data[i].Name, " ", "%20")
				urlProducts := url2 + "?title=" + nameProduct
				bodyProducts := getValueFromStore(urlProducts)

				var productCheck struct {
					Products Product `json:"products"`
				}

				json.Unmarshal(bodyProducts, &productCheck)

				if len(productCheck.Products) != 0 {
					for i5 := 0; i5 < len(productCheck.Products); i5++ {
						IDProduct := productCheck.Products[i5].ID
						for countCollectData := 0; countCollectData < len(collectData); countCollectData++ {
							valuesProduct := map[string]map[string]interface{}{
								"collect": {
									"product_id":    IDProduct,
									"collection_id": collectData[countCollectData],
								},
							}
							urlColection := "https://" + postAPIkey + ":" + postPassword + "@" + postDomain + "/admin/api/2020-04/collects.json"
							postValueToStore(valuesProduct, urlColection)
						}
					}
				} else {
					postValueToStore(values, url2)

					bodyProducts := getValueFromStore(urlProducts)
					json.Unmarshal(bodyProducts, &productCheck)

					if len(productCheck.Products) != 0 {
						for i5 := 0; i5 < len(productCheck.Products); i5++ {
							IDProduct := productCheck.Products[i5].ID
							for countCollectData := 0; countCollectData < len(collectData); countCollectData++ {
								valuesProduct := map[string]map[string]interface{}{
									"collect": {
										"product_id":    IDProduct,
										"collection_id": collectData[countCollectData],
									},
								}
								urlColection := "https://" + postAPIkey + ":" + postPassword + "@" + postDomain + "/admin/api/2020-04/collects.json"
								postValueToStore(valuesProduct, urlColection)
							}
						}
					}
				}
			}
		}
	}
}

func postValueToStore(values map[string]map[string]interface{}, url string) {

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

func getValueFromWp(url string) []byte {
	token1 := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9kZXYubG9jYWwiLCJpYXQiOjE1OTE5MzI0MDIsIm5iZiI6MTU5MTkzMjQwMiwiZXhwIjoxNTkyNTM3MjAyLCJkYXRhIjp7InVzZXIiOnsiaWQiOjEsInR5cGUiOiJ3cF91c2VyIiwidXNlcl9sb2dpbiI6ImFkbWluIiwidXNlcl9lbWFpbCI6ImRldi1lbWFpbEBmbHl3aGVlbC5sb2NhbCIsImFwaV9rZXkiOiIxQWZCZXlvU0U1a3Axa2lDMDNaYjJpSURZIn19fQ.5ls54WhX6GDPeMbPTOoVF_aqUqwg7OnkxjXn9qowNR8"

	res, err1 := http.NewRequest("GET", url, nil)
	res.Header.Set("Content-Type", "application/json")
	res.Header.Set("Authorization", "Bearer "+token1)
	client1 := &http.Client{}
	resp1, err1 := client1.Do(res)
	if err1 != nil {
		panic(err1)
	}
	defer resp1.Body.Close()
	body, _ := ioutil.ReadAll(resp1.Body)

	return body
}

func createCollection(Collection string, postURL string, postDomain string, postPassword string, postAPIkey string) int64 {
	url := "https://" + postAPIkey + ":" + postPassword + "@" + postDomain + "/admin/api/2020-04/custom_collections.json"
	nameCollection := strings.ReplaceAll(Collection, " ", "%20")
	urlCollection := url + "?title=" + nameCollection
	body := getValueFromStore(urlCollection)

	var data struct {
		CustomCollections CustomCollections `json:"custom_collections"`
	} // dataCollection

	json.Unmarshal(body, &data)

	var IDCollection int64
	if len(data.CustomCollections) == 0 {
		valCategoies := map[string]map[string]interface{}{
			"custom_collection": {
				"title": nameCollection,
			},
		}
		postValueToStore(valCategoies, url)
		urlCollectionCreateNew := url + "?title=" + nameCollection
		bodyCollection := getValueFromStore(urlCollectionCreateNew)

		var dataCollectionCreateNew struct {
			CustomCollections CustomCollections `json:"custom_collections"`
		}
		json.Unmarshal(bodyCollection, &dataCollectionCreateNew)

		IDCollection = dataCollectionCreateNew.CustomCollections[0].ID

	} else {
		IDCollection = data.CustomCollections[0].ID
	}

	return IDCollection

}

func getValueFromStore(url string) []byte {
	r, err1 := http.NewRequest("GET", url, nil)
	r.Header.Set("Content-Type", "application/json")
	client1 := &http.Client{}
	resp1, err1 := client1.Do(r)
	if err1 != nil {
		panic(err1)
	}
	defer resp1.Body.Close()
	body, _ := ioutil.ReadAll(resp1.Body)

	return body
}
