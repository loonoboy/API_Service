package postgres

import (
	"API_Service/internal/dto"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

type ArticlePostgres struct {
	db *sqlx.DB
}

func NewArticlePostgres(db *sqlx.DB) *ArticlePostgres {
	return &ArticlePostgres{db: db}
}

func (r *ArticlePostgres) CreateArticle(userId int, article dto.Article) (int, error) {
	const op = "repository.article_postgres.CreateArticle"
	var id int
	CreateArticleQuery := fmt.Sprintf("INSERT INTO %s (user_id, title, content) VALUES ($1, $2, $3) RETURNING id", articlesTable)
	row := r.db.QueryRow(CreateArticleQuery, userId, article.Title, article.Content)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (r *ArticlePostgres) GetAllById(userId int) ([]dto.Article, error) {
	op := "repository.article_postgres.etAll"
	var articles []dto.Article
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", articlesTable)
	err := r.db.Select(&articles, query, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return articles, nil
}

func (r *ArticlePostgres) GetArticleById(userId, articleId int) (dto.Article, error) {
	op := "repository.article_postgres.getArticleById"
	var article dto.Article
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 and id=$2", articlesTable)
	err := r.db.Get(&article, query, userId, articleId)
	if err != nil {
		return article, fmt.Errorf("%s: %w", op, err)
	}
	return article, nil
}

func (r *ArticlePostgres) DeleteArticleById(userId, articleId int) error {
	op := "repository.article_postgres.deleteArticleById"
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 and id=$2", articlesTable)
	_, err := r.db.Exec(query, userId, articleId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}

func (r *ArticlePostgres) UpdateArticleById(userId, articleId int, input dto.UpdateArticle) error {
	op := "repository.article_postgres.updateArticleById"
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValue = append(setValue, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Content != nil {
		setValue = append(setValue, fmt.Sprintf("content=$%d", argId))
		args = append(args, *input.Content)
		argId++
	}

	setQuery := strings.Join(setValue, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id=$%d and id=$%d", articlesTable, setQuery, argId, argId+1)
	args = append(args, userId, articleId)
	log.Printf("updateQuery: %s", query)
	log.Printf("args: %v", args)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}

func (r *ArticlePostgres) GetLastArticles(count int) ([]dto.Article, error) {
	op := "repository.article_postgres.getLastArticles"
	var articles []dto.Article
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC LIMIT $1", articlesTable)
	err := r.db.Select(&articles, query, count)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return articles, nil
}
