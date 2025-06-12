package handler

import (
	"API_Service/internal/dto"
	resp "API_Service/internal/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ResponseCreateArticle struct {
	resp.Response
	Id int `json:"id"`
}

type ResponseGetAll struct {
	resp.Response
	Data []dto.Article `json:"data"`
}

type ResponseGetArticleById struct {
	resp.Response
	Data dto.Article `json:"data"`
}

func (h *Handler) createArticle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.article.createArticle"
		log := h.log.With(
			zap.String("op", op),
			zap.String("request_id", middleware.GetReqID(r.Context())),
		)
		userId := r.Context().Value(userCtxKey)
		if userId == nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("user id not found"))
			return
		}
		var input dto.Article
		if err := render.DecodeJSON(r.Body, &input); err != nil {
			log.Error("parse request body failed", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, resp.Error("parse request body failed"))
			return
		}

		id, err := h.service.Article.CreateArticle(userId.(int), input)
		if err != nil {
			log.Error("create article failed", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("create article failed"))
			return
		}

		render.JSON(w, r, ResponseCreateArticle{
			Response: resp.OK(),
			Id:       id,
		})
	}

}

func (h *Handler) getAllArticles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.article.getAllArticles"
		log := h.log.With(
			zap.String("op", op),
			zap.String("request_id", middleware.GetReqID(r.Context())),
		)
		articles, err := h.service.Article.GetAllArticles()
		if err != nil {
			log.Error("get articles failed", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("get articles failed"))
			return
		}
		render.JSON(w, r, ResponseGetAll{
			Response: resp.OK(),
			Data:     articles,
		})
	}
}

func (h *Handler) getAllArticlesById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.article.getAllArticlesById"
		log := h.log.With(
			zap.String("op", op),
			zap.String("request_id", middleware.GetReqID(r.Context())),
		)
		userId := r.Context().Value(userCtxKey)
		if userId == nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("user id not found"))
			return
		}
		articles, err := h.service.Article.GetAllById(userId.(int))
		if err != nil {
			log.Error("get articles failed", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("get articles failed"))
			return
		}
		render.JSON(w, r, ResponseGetAll{
			Response: resp.OK(),
			Data:     articles,
		})
	}

}

func (h *Handler) getArticleById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.article.getAllArticles"
		log := h.log.With(
			zap.String("op", op),
			zap.String("request_id", middleware.GetReqID(r.Context())),
		)
		userId := r.Context().Value(userCtxKey)
		if userId == nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("user id not found"))
			return
		}

		articleID, err := strconv.Atoi(chi.URLParam(r, "articleID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, resp.Error("invalid article id"))
			return
		}
		article, err := h.service.Article.GetArticleById(userId.(int), articleID)
		if err != nil {
			log.Error("get article failed", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("get article failed"))
			return
		}
		render.JSON(w, r, ResponseGetArticleById{
			Response: resp.OK(),
			Data:     article,
		})
	}
}

func (h *Handler) updateArticleById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.article.updateArticleById"
		log := h.log.With(
			zap.String("op", op),
			zap.String("request_id", middleware.GetReqID(r.Context())),
		)
		userId := r.Context().Value(userCtxKey)
		if userId == nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("user id not found"))
			return
		}

		articleID, err := strconv.Atoi(chi.URLParam(r, "articleID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, resp.Error("invalid article id"))
			return
		}

		var input dto.UpdateArticle
		if err := render.DecodeJSON(r.Body, &input); err != nil {
			log.Error("parse request body failed", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, resp.Error("parse request body failed"))
			return
		}

		if err = h.service.Article.UpdateArticleById(userId.(int), articleID, input); err != nil {
			log.Error("update article failed", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("update article  failed"))
			return
		}
		render.JSON(w, r, resp.OK())
	}
}

func (h *Handler) deleteArticleById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.article.getAllArticles"
		log := h.log.With(
			zap.String("op", op),
			zap.String("request_id", middleware.GetReqID(r.Context())),
		)
		userId := r.Context().Value(userCtxKey)
		if userId == nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("user id not found"))
			return
		}

		articleID, err := strconv.Atoi(chi.URLParam(r, "articleID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, resp.Error("invalid article id"))
			return
		}
		err = h.service.Article.DeleteArticleById(userId.(int), articleID)
		if err != nil {
			log.Error("delete article failed", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("delete article failed"))
			return
		}
		render.JSON(w, r, resp.OK())
	}
}
