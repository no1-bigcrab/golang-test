package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Context is

//ProductsPageGet in function
func (a *App) ProductsPageGet(w http.ResponseWriter, r *http.Request) {
	url := "http://dev.local/wp-json/wc/v3/products"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9kZXYubG9jYWwiLCJpYXQiOjE1OTEyNTYxMTksIm5iZiI6MTU5MTI1NjExOSwiZXhwIjoxNTkxODYwOTE5LCJkYXRhIjp7InVzZXIiOnsiaWQiOjEsInR5cGUiOiJ3cF91c2VyIiwidXNlcl9sb2dpbiI6ImFkbWluIiwidXNlcl9lbWFpbCI6ImRldi1lbWFpbEBmbHl3aGVlbC5sb2NhbCIsImFwaV9rZXkiOiIxaUoxNHJydmFjbEZCTjVBSmpiMUpDa2ZFIn19fQ.2Pmo5DefSc69txjGtAJX8zwU7Oxw9rS_V7_wm5ARiYg")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	jsonValue, _ := json.Marshal(data)
	fmt.Println(bytes.NewBuffer(jsonValue))
	//fmt.Println(data)

}
