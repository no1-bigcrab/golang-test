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

//Users cs cs
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

//Context is struct
type Context struct {
	Title   string
	Results []User
}

//HomePageGet in function
func (a *App) HomePageGet(w http.ResponseWriter, r *http.Request) {

	path := urls.PathUrl()

	jsonFile, err := os.Open(path.CONFIG_PATH + "dataConfig.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	//push data
	push(w, "static/css/bootstrap.min.css")
	context := Context{
		Title:   "My Fruits",
		Results: users.Users,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := template.Must(template.ParseFiles(path.TEMPLATE_PATH + "/form.html"))
	html.Execute(w, context)

}

//HomePagePost is func
func (a *App) HomePagePost(w http.ResponseWriter, r *http.Request) {
	//url := "http://localhost:9001"

	//fmt.Println("URL:>", url)
	apiKey := r.FormValue("api_key")
	password := r.FormValue("password")

	values := map[string]string{"apiKey": apiKey, "password": password}

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
