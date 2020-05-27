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
	Title    string `json:"title"`
	Products models.Products
	Pages    models.Pages
	Blogs    models.Blogs
	Article  models.Article
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
	//url := "http://localhost:9001"

	title := r.FormValue("api-title")
	data := r.FormValue("api_key")

	values := map[string]string{"data-title": title, "data-value": data}

	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post("https://httpbin.org/post",
		"application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
}

func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}
