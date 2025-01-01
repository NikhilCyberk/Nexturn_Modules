package model

type Blog struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	TimeStamp string `json:"timestamp"`
}
