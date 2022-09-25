package models

type Question struct {
	Title    string `json:"title"`
	Response string `json:"response"`
}

type Quiz struct {
	Id        string     `json:"id"`
	Questions []Question `json:"questions"`
	CreatedAt string     `json:"created_at"`
	EditedAt  string     `json:"edited_at"`
}
