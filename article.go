package API_Service

type Article struct {
	Id      int    `json:"id"`
	UserId  string `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
