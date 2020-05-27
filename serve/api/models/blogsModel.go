package models

//Blogs json cd
type Blogs struct {
	Handle             string      `json:"handle"`
	Title              string      `json:"title"`
	Commentable        string      `json:"commentable"`
	Feedburner         interface{} `json:"feedburner"`
	FeedburnerLocation interface{} `json:"feedburner_location"`
	TemplateSuffix     interface{} `json:"template_suffix"`
	Tags               string      `json:"tags"`
	AdminGraphqlAPIID  string      `json:"admin_graphql_api_id"`
}
