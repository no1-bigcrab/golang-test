package models

// CustomCollections abc
type CustomCollections struct {
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
	BodyHTML          interface{} `json:"body_html"`
	Handle            string      `json:"handle"`
	ID                int64       `json:"id"`
	PublishedAt       string      `json:"published_at"`
	PublishedScope    string      `json:"published_scope"`
	SortOrder         string      `json:"sort_order"`
	TemplateSuffix    interface{} `json:"template_suffix"`
	Title             string      `json:"title"`
	UpdatedAt         string      `json:"updated_at"`
}
