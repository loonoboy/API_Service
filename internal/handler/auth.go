package handler

import (
	"API_Service"
	resp "API_Service/internal/response"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type Response struct {
	resp.Response
	Id int `json:"id"`
}

func (h *Handler) signUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.auth.signUp"

		log := h.log
		log.With(
			zap.String("op", op),
			zap.String("required_id", middleware.GetReqID(r.Context())),
		)

		var input API_Service.User

		err := render.DecodeJSON(r.Body, &input)

		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")
			render.JSON(w, r, resp.Error("empty request"))
		}
		if err != nil {
			log.Error("failed to decode request body", zap.Error(err))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}
		log.Info("request body decoded", zap.Any("request", input))

		if err := validator.New().Struct(input); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", zap.Error(err))

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		id, err := h.service.Authorization.CreateUser(input)
		if err != nil {
			log.Error("failed to create user", zap.Error(err))
			render.JSON(w, r, resp.Error("failed to create user"))

			return
		}

		log.Info("url added", zap.Int("id", id))

		responseOK(w, r, id)
	}
}

func (h *Handler) signIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func responseOK(w http.ResponseWriter, r *http.Request, id int) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Id:       id,
	})
}
