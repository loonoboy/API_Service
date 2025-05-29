package handler

import (
	"API_Service/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/auth", func(auth chi.Router) {
		auth.Post("/signup", h.signUp)
		auth.Post("/signin", h.signIn)
	})
	router.Route("/api", func(api chi.Router) {
		api.Route("/article", func(article chi.Router) {
			article.Post("/", h.createArticle)
			article.Get("/", h.getAllArticles)
			article.Get("/:id", h.getArticleById)
			article.Put("/:id", h.updateArticle)
			article.Delete("/:id", h.deleteArticle)
		})
	})

	return router
}
