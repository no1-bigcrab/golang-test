package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-test/serve/api/models"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

//Asset in here
type Asset struct {
	Assets models.Assets `json:"asset"`
}

//PageGet it here
func (a *App) PageGet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postURL := r.FormValue("url")
	postDomain := r.FormValue("domain")
	postPassword := r.FormValue("password")
	postAPIkey := r.FormValue("apikey")
	apiTitle := r.FormValue("api-title")

	url := postURL + "pages"
	body := getValueFromWp(url)
	var dataPage models.Page
	json.Unmarshal(body, &dataPage)

	for i := 0; i < len(dataPage); i++ {

		var myString = dataPage[i].Content.Rendered
		var myRegex = regexp.MustCompile(`<img[^>]+\bsrcset=["']([^"']+)["']`)
		var myRegexSrcSet = regexp.MustCompile(`<img[^>]+\bsrcset=["']([^"']+)["']`)
		var imgTags = myRegex.FindAllStringSubmatch(myString, -1)
		var imgTagsSrcSet = myRegexSrcSet.FindAllStringSubmatch(myString, -1)
		if len(imgTags) > 0 {
			out := make([]string, len(imgTags))

			title := dataPage[i].Title.Rendered

			keyImage := strings.ReplaceAll(title, " ", "-")
			for i1 := range out {
				//fmt.Println(imgTagsSrcSet[i1][1])
				valuesPage := map[string]map[string]interface{}{
					"asset": {
						"key": "assets/" + keyImage + ".jpg",
						"src": imgTags[i1][1],
					},
				}

				urlAsset := "https://" + postAPIkey + ":" + postPassword + "@" + postDomain + "/admin/api/2020-04/themes/99734585509/assets.json"
				urlAssetValue := urlAsset + "?asset[key]=assets/" + keyImage + ".jpg"

				CheckAssets, err := checkValueDataFromStore(urlAssetValue)
				if err != nil {
					fmt.Println(err)
				}

				if len(CheckAssets) > 0 {

					var dataPutPage Asset
					json.Unmarshal(CheckAssets, &dataPutPage)
					myString = strings.ReplaceAll(myString, imgTags[i1][1], dataPutPage.Assets.PublicURL)
					myString = strings.ReplaceAll(myString, imgTagsSrcSet[i1][1], dataPutPage.Assets.PublicURL)

				} else {

					dataPutImg := putDataToStore(valuesPage, urlAsset)
					var dataPutPage Asset
					json.Unmarshal(dataPutImg, &dataPutPage)
					myString = strings.ReplaceAll(myString, imgTags[i1][1], dataPutPage.Assets.PublicURL)
					myString = strings.ReplaceAll(myString, imgTagsSrcSet[i1][1], dataPutPage.Assets.PublicURL)

				}
				//fmt.Println(dataPage[i].Title.Rendered, postURL, postDomain, postPassword, postAPIkey, apiTitle)
				checkAndPostValuePage(dataPage[i].Title.Rendered, myString, postURL, postDomain, postPassword, postAPIkey, apiTitle)

			}
		} else {
			checkAndPostValuePage(dataPage[i].Title.Rendered, myString, postURL, postDomain, postPassword, postAPIkey, apiTitle)
		}

	}
}

func putDataToStore(value map[string]map[string]interface{}, url string) []byte {
	jsonValue, _ := json.Marshal(value)
	//fmt.Println(bytes.NewBuffer(jsonValue))
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func checkAndPostValuePage(title string, datakeyImage string, postURL string, postDomain string, postPassword string, postAPIkey string, apiTitle string) {
	url := "https://" + postAPIkey + ":" + postPassword + "@" + postDomain + "/admin/api/2020-04/pages.json"
	titleCheck := strings.ReplaceAll(title, " ", "%20")

	urlCheck := url + "?title=" + titleCheck

	//fmt.Println(urlCheck)

	dataCheck, err := checkValueDataFromStore(urlCheck)
	if err != nil {
		fmt.Println(err)
	}

	var dataPage models.Page
	json.Unmarshal(dataCheck, &dataPage)

	//fmt.Println(len(dataCheck))

	if len(dataCheck) == 12 {
		dataCreatePage := map[string]map[string]interface{}{
			"page": {
				"title":     title,
				"body_html": datakeyImage,
			},
		}
		postValueToStore(dataCreatePage, url)
	}
}

func checkValueDataFromStore(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	var body []byte

	if err != nil {
		log.Print(err)
		body = []byte("Error")
	} else {
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		fmt.Println("response Status:", resp.Status)

		body, _ = ioutil.ReadAll(resp.Body)
	}

	return body, nil
}
