package models

//Article cd
type Article []struct {
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
		PredecessorVersion []struct {
			Href string `json:"href"`
			ID   int    `json:"id"`
		} `json:"predecessor-version"`
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
		WpFeaturedmedia []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"wp:featuredmedia"`
		WpTerm []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
			Taxonomy   string `json:"taxonomy"`
		} `json:"wp:term"`
	} `json:"_links"`
	Author        int    `json:"author"`
	Categories    []int  `json:"categories"`
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
	FeaturedMedia int    `json:"featured_media"`
	Format        string `json:"format"`
	GUID          struct {
		Rendered string `json:"rendered"`
	} `json:"guid"`
	ID                      int           `json:"id"`
	JetpackFeaturedMediaURL string        `json:"jetpack_featured_media_url"`
	Link                    string        `json:"link"`
	Meta                    []interface{} `json:"meta"`
	Modified                string        `json:"modified"`
	ModifiedGmt             string        `json:"modified_gmt"`
	PingStatus              string        `json:"ping_status"`
	Slug                    string        `json:"slug"`
	Status                  string        `json:"status"`
	Sticky                  bool          `json:"sticky"`
	Tags                    []int         `json:"tags"`
	Template                string        `json:"template"`
	Title                   struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Type string `json:"type"`
}
