package API_Service

import "errors"

type Article struct {
	Id      int    `json:"id" db:"id"`
	UserId  string `json:"userId" db:"user_id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}

type UpdateArticle struct {
	Title   *string `json:"title" db:"title"`
	Content *string `json:"content" db:"content"`
}

func (i UpdateArticle) Validate() error {
	if i.Title == nil && i.Content == nil {
		return errors.New("invalid UpdateArticle")
	}

	return nil
}
