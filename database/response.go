package database

type APIResponse[T any] []struct {
	Time   string `json:"time"`
	Status string `json:"status"`
	Result []T    `json:"result"`
}

type Result[T any] struct {
	Time   string `json:"time"`
	Status string `json:"status"`
	Result []T    `json:"result"`
}
