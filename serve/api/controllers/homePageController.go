package controllers

import (
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
	fmt.Println(url)
	for key, values := range r.PostForm {
		fmt.Println(key, values)
	}

	//fmt.Println(r.Form)
	// req, err := http.NewRequest("POST", url, nil)
	// req.Header.Set("X-Shopify-Access-Token", r.FormValue("password"))
	// req.Header.Set("Content-Type", "application/json")
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// fmt.Println("response Status:", resp.Status)
	// body, _ := ioutil.ReadAll(resp.Body)
	// var m interface{}
	// json.Unmarshal(body, &m)
	// fmt.Println(m)
	// switch v := m.(type) {
	// case map[string]interface{}:
	// 	{
	// 		switch h := v["products"].(type) {
	// 		case []interface{}:
	// 			{
	// 				fmt.Println("ok")
	// 				for _, x := range h {
	// 					if x.(map[string]interface{})["role"] == "main" {
	// 						json.NewEncoder(w).Encode(map[string]interface{}{"data": x.(map[string]interface{})["id"]})
	// 					}
	// 				}
	// 			}
	// 		default:
	// 			fmt.Println("No type found")
	// 		}
	// 	}
	// default:
	// 	fmt.Println("No type found")
	// }
}

func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}
