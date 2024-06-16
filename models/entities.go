package models

type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoSearchParams struct {
	Title       string
	Description string
	Completed   *bool
}