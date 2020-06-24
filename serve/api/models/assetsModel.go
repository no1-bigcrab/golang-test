package models

// Assets abc
type Assets struct {
	Attachment  string        `json:"attachment"`
	ContentType string        `json:"content_type"`
	CreatedAt   string        `json:"created_at"`
	Key         string        `json:"key"`
	PublicURL   string        `json:"public_url"`
	Size        int           `json:"size"`
	ThemeID     int64         `json:"theme_id"`
	UpdatedAt   string        `json:"updated_at"`
	Warnings    []interface{} `json:"warnings"`
}
