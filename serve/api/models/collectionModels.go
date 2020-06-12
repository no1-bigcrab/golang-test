package models

// Collection abc
type Collection []struct {
	ID                int         `json:"id"`
	Handle            string      `json:"handle"`
	Title             string      `json:"title"`
	BodyHTML          string      `json:"body_html"`
	SortOrder         string      `json:"sort_order"`
	TemplateSuffix    interface{} `json:"template_suffix"`
	PublishedScope    string      `json:"published_scope"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
	Image             struct {
		Alt    string `json:"alt"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Src    string `json:"src"`
	} `json:"image"`
}
