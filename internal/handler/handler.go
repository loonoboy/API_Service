package handler

import (
	"API_Service/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type Handler struct {
	log     *zap.Logger
	service *service.Service
}

func NewHandler(log *zap.Logger, service *service.Service) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.signUp())
		r.Get("/sign-in", h.signIn())
	})
	router.Route("/api", func(r chi.Router) {
		r.Use(h.userIdentity())
		r.Route("/article", func(r chi.Router) {
			r.Post("/", h.createArticle())
			r.Get("/", h.getAllArticles())
			r.Route("/{articleID}", func(r chi.Router) {
				r.Get("/", h.getArticleById())
				r.Put("/", h.updateArticleById())
				r.Delete("/", h.deleteArticleById())
			})

		})
	})

	return router
}
