package models

//Article cd
type Article struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Tags     string `json:"tags"`
	BodyHTML string `json:"body_html"`
	Image    struct {
		Attachment string `json:"attachment"`
	} `json:"image"`
}
