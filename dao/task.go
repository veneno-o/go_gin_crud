package dao

type Task struct {
	ID      int     `json:"id"`
	Content *string `json:"content"`
	Status  int     `json:"status"`
}
