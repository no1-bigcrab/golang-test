package models

//Page abcs
type Page []struct {
	Links struct {
		About []struct {
			Href string `json:"href"`
		} `json:"about"`
		Author []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"author"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		Curies []struct {
			Href      string `json:"href"`
			Name      string `json:"name"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
		Replies []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"replies"`
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		VersionHistory []struct {
			Count int    `json:"count"`
			Href  string `json:"href"`
		} `json:"version-history"`
		WpAttachment []struct {
			Href string `json:"href"`
		} `json:"wp:attachment"`
	} `json:"_links"`
	Author        int    `json:"author"`
	CommentStatus string `json:"comment_status"`
	Content       struct {
		Protected bool   `json:"protected"`
		Rendered  string `json:"rendered"`
	} `json:"content"`
	Date    string `json:"date"`
	DateGmt string `json:"date_gmt"`
	Excerpt struct {
		Protected bool   `json:"protected"`
		Rendered  string `json:"rendered"`
	} `json:"excerpt"`
	FeaturedMedia int `json:"featured_media"`
	GUID          struct {
		Rendered string `json:"rendered"`
	} `json:"guid"`
	ID          int           `json:"id"`
	Link        string        `json:"link"`
	MenuOrder   int           `json:"menu_order"`
	Meta        []interface{} `json:"meta"`
	Modified    string        `json:"modified"`
	ModifiedGmt string        `json:"modified_gmt"`
	Parent      int           `json:"parent"`
	PingStatus  string        `json:"ping_status"`
	Slug        string        `json:"slug"`
	Status      string        `json:"status"`
	Template    string        `json:"template"`
	Title       struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Type string `json:"type"`
}
