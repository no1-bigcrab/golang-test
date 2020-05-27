package models

//Pages abcs
type Pages struct {
	Title             string      `json:"title"`
	ShopID            int         `json:"shop_id"`
	Handle            string      `json:"handle"`
	BodyHTML          string      `json:"body_html"`
	Author            string      `json:"author"`
	TemplateSuffix    interface{} `json:"template_suffix"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
}
