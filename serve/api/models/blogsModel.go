package models

//Blogs json cd
type Blogs []struct {
	ID                 int64       `json:"id"`
	Handle             string      `json:"handle"`
	Title              string      `json:"title"`
	UpdatedAt          string      `json:"updated_at"`
	Commentable        string      `json:"commentable"`
	Feedburner         interface{} `json:"feedburner"`
	FeedburnerLocation interface{} `json:"feedburner_location"`
	CreatedAt          string      `json:"created_at"`
	TemplateSuffix     interface{} `json:"template_suffix"`
	Tags               []string    `json:"tags"`
	AdminGraphqlAPIID  string      `json:"admin_graphql_api_id"`
}
