package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-test/serve/api/urls"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

//DataURL in herhe
type DataURL []struct {
	APIKey   string `json:"apiKey"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
	URLGet   string `json:"urlGet"`
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

	var data DataURL
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		panic(err)
	}

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
