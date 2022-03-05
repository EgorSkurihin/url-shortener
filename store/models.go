package store

type URLModel struct {
	Id       int    `json:"id"`
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}
