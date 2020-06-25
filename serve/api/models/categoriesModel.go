package models

//Categories json cd
type Categories struct {
	ID          int           `json:"id"`
	Count       int           `json:"count"`
	Description string        `json:"description"`
	Link        string        `json:"link"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Taxonomy    string        `json:"taxonomy"`
	Parent      int           `json:"parent"`
	Meta        []interface{} `json:"meta"`
	Links       struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		About []struct {
			Href string `json:"href"`
		} `json:"about"`
		WpPostType []struct {
			Href string `json:"href"`
		} `json:"wp:post_type"`
		Curies []struct {
			Name      string `json:"name"`
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
	} `json:"_links"`
}
