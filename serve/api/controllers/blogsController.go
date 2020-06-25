package controllers

import (
	"encoding/json"
	"golang-test/serve/api/models"
	"net/http"
)

//Blogs json cd
type Blogs []struct {
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

//BlogPageGet it here
func (a *App) BlogPageGet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postURL := r.FormValue("url")
	// postDomain := r.FormValue("domain")
	// postPassword := r.FormValue("password")
	// postAPIkey := r.FormValue("apikey")
	// apiTitle := r.FormValue("api-title")
	url := postURL + "posts"
	body := getValueFromWp(url)
	var dataPage models.Article
	json.Unmarshal(body, &dataPage)

}
