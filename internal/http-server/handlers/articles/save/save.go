package save

import (
	resp "API_Service/internal/lib/api/response"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

type Request struct {
	UserId  string `json:"user_id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type Response struct {
	resp.Response
}

type ArticleSaver interface {
	SaveArticle(userId, title, content string) error
}

func New(log *zap.Logger, articleSaver ArticleSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.article.save.new"

		log = log.With(
			zap.String("op", op),
			zap.String("request_id", middleware.GetReqID(r.Context())),
		)
		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request dody", zap.Error(err))

			log.Info("body", zap.Any("body", r.Body))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		log.Info("request body decode", zap.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("failed to validate request", zap.Error(err))
			render.JSON(w, r, resp.ValidationError(validateErr))
			return
		}

		err = articleSaver.SaveArticle(req.UserId, req.Title, req.Content)
		if err != nil {
			log.Error("failed to save article", zap.Error(err))
			render.JSON(w, r, resp.Error("failed to save article"))
			return
		}
		log.Info("save article success")

		render.JSON(w, r, Response{
			Response: resp.OK(),
		})
	}

}
