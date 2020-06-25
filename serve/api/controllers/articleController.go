package controllers

import (
	"encoding/json"
	"fmt"
	"golang-test/serve/api/models"
	"net/http"
	"strconv"
	"strings"
)

//Article json cd
type Article struct {
	ID   int `json:"id"`
	GUID struct {
		Rendered string `json:"rendered"`
	} `json:"guid"`
	Modified    string `json:"modified"`
	ModifiedGmt string `json:"modified_gmt"`
	Slug        string `json:"slug"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Link        string `json:"link"`
	Title       struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Content struct {
		Rendered  string `json:"rendered"`
		Protected bool   `json:"protected"`
	} `json:"content"`
	Excerpt struct {
		Rendered  string `json:"rendered"`
		Protected bool   `json:"protected"`
	} `json:"excerpt"`
	Author                  int           `json:"author"`
	FeaturedMedia           int           `json:"featured_media"`
	CommentStatus           string        `json:"comment_status"`
	PingStatus              string        `json:"ping_status"`
	Sticky                  bool          `json:"sticky"`
	Template                string        `json:"template"`
	Format                  string        `json:"format"`
	Meta                    []interface{} `json:"meta"`
	Categories              []int         `json:"categories"`
	Tags                    []int         `json:"tags"`
	JetpackFeaturedMediaURL string        `json:"jetpack_featured_media_url"`
}

//ArticlePageGet it here
func (a *App) ArticlePageGet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postURL := r.FormValue("url")
	postDomain := r.FormValue("domain")
	postPassword := r.FormValue("password")
	postAPIkey := r.FormValue("apikey")
	apiTitle := r.FormValue("api-title")
	url := postURL + "posts"
	body := getValueFromWp(url)
	var dataArticle models.Article
	json.Unmarshal(body, &dataArticle)
	// jsonValue, _ := json.Marshal(dataArticle)
	// fmt.Println(bytes.NewBuffer(jsonValue))
	for i := 0; i < len(dataArticle); i++ {
		urlCategories := postURL + "categories" + "/" + strconv.Itoa(dataArticle[i].ID)
		fmt.Println(urlCategories)
		bodyCategories := getValueFromWp(urlCategories)
		var dataCategories models.Categories
		json.Unmarshal(bodyCategories, &dataCategories)
		// jsonValue, _ := json.Marshal(dataArticle)
		// fmt.Println(bytes.NewBuffer(jsonValue))
		dataHandle := strings.ReplaceAll(strings.ToLower(dataCategories.Name), " ", "-")
		fmt.Println(dataHandle, urlCategories)
		urlCategoriesStore := "https://" + postAPIkey + ":" + postPassword + "@" + postDomain + "/admin/api/2020-04/blogs.json?handle=" + dataHandle
		dataHandel, err := checkValueDataFromStore(urlCategoriesStore)
		if err != nil {
			fmt.Println(err)
		}
		if len(dataHandel) > 0 {
			var dataCategoriesStore interface{}
			json.Unmarshal(dataHandel, &dataCategoriesStore)
			//fmt.Println(dataCategoriesStore)

			urlArticle := "https://" + postAPIkey + ":" + postPassword + "@" + postDomain + "/admin/api/2020-04/blogs/" + strconv.Itoa(dataCategories.ID) + "/" + apiTitle + ".json"
		}

		//postValueToStore( ,urlArticle)
	}

}
