package model

type Book struct {
	ID     string `json:"id,omitempty"` // will be set after insert
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Year   int    `json:"year,omitempty"`
}
