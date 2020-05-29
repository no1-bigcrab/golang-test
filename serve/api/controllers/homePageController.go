package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-test/serve/api/models"
	"golang-test/serve/api/urls"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

//Context is

//superData
type superData struct {
	Title   string `json:"title"`
	Product models.Product
	Pages   models.Pages
	Blogs   models.Blogs
	Article models.Article
}

//HomePageGet in function
func (a *App) HomePageGet(w http.ResponseWriter, r *http.Request) {

	path := urls.PathUrl()

	jsonFile, err := os.Open(path.CONFIG_PATH + "dataConfig.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	body, _ := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err.Error())
	}

	var data superData
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		panic(err)
	}
	//fmt.Println(data.Products.Images[0].Src)

	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := template.Must(template.ParseFiles(path.TEMPLATE_PATH + "/form.html"))
	html.Execute(w, data)
}

//HomePagePost is func
func (a *App) HomePagePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := "https://" + r.FormValue("apikey") + ":" + r.FormValue("password") + "@" + r.FormValue("hostname") + "/admin/api/2020-04/" + r.FormValue("api-title") + ".json"

	if r.FormValue("api-title") == "products" {

		values := map[string]map[string]string{
			"product": {
				"title":       r.FormValue("title"),
				"body_html":   r.FormValue("html-body"),
				"vendor":      r.FormValue("vendor"),
				"productType": r.FormValue("product-type"),
			},
		}
		jsonValue, _ := json.Marshal(values)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		req.Header.Set("X-Shopify-Access-Token", r.FormValue("password"))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		json.NewEncoder(w).Encode(map[string]string{"code": resp.Status})

	} else if r.FormValue("api-title") == "blogs" {

		values := map[string]map[string]string{
			"blog": {
				"title":      r.FormValue("title"),
				"metafields": "global",
			},
		}
		jsonValue, _ := json.Marshal(values)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		req.Header.Set("X-Shopify-Access-Token", r.FormValue("password"))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		json.NewEncoder(w).Encode(map[string]string{"code": resp.Status})

	} else if r.FormValue("api-title") == "pages" {

		values := map[string]map[string]string{
			"page": {
				"title":     r.FormValue("title"),
				"body_html": r.FormValue("html-body"),
				"author":    r.FormValue("author"),
			},
		}
		jsonValue, _ := json.Marshal(values)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		req.Header.Set("X-Shopify-Access-Token", r.FormValue("password"))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		json.NewEncoder(w).Encode(map[string]string{"code": resp.Status})

	} else {
		values := map[string]map[string]string{
			"article": {
				"title":     r.FormValue("title"),
				"body_html": r.FormValue("html-body"),
				"author":    r.FormValue("author"),
				"tags":      r.FormValue("tags"),
				"image":     r.FormValue("image"),
			},
		}
		jsonValue, _ := json.Marshal(values)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		req.Header.Set("X-Shopify-Access-Token", r.FormValue("password"))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		json.NewEncoder(w).Encode(map[string]string{"code": resp.Status})

	}
}
