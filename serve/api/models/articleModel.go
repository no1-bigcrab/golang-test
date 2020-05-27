package models

//Article cd
type Article struct {
	Title             string      `json:"title"`
	BodyHTML          string      `json:"body_html"`
	BlogID            int         `json:"blog_id"`
	Author            string      `json:"author"`
	UserID            interface{} `json:"user_id"`
	SummaryHTML       interface{} `json:"summary_html"`
	TemplateSuffix    interface{} `json:"template_suffix"`
	Handle            string      `json:"handle"`
	Tags              string      `json:"tags"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
	Image             struct {
		Alt    string `json:"alt"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Src    string `json:"src"`
	} `json:"image"`
}
