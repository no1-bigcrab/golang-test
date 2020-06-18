package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

//PageGet it here
func (a *App) PageGet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postURL := r.FormValue("url")
	// postDomain := r.FormValue("domain")
	// postPassword := r.FormValue("password")
	// postAPIkey := r.FormValue("apikey")
	for count := 1; count < 10; count++ {
		url := postURL + "pages/?per_pages=" + strconv.Itoa(count)
		body := getValueFromWp(url)
		var data interface{}
		json.Unmarshal(body, &data)
		jsonValue, _ := json.Marshal(data)
		fmt.Println(bytes.NewBuffer(jsonValue))
	}
}
