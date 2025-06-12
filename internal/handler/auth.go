package handler

import (
	"API_Service/internal/dto"
	resp "API_Service/internal/response"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type ResponsesignUpOk struct {
	resp.Response
	Id int `json:"id"`
}

type ResponsesignInok struct {
	resp.Response
	Token string `json:"token"`
}

type signInInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *Handler) signUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.auth.signUp"

		log := h.log.With(
			zap.String("op", op),
			zap.String("request_id", middleware.GetReqID(r.Context())),
		)

		var input dto.User

		err := render.DecodeJSON(r.Body, &input)

		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")
			render.JSON(w, r, resp.Error("empty request"))
			return
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

		log.Info("user added", zap.Int("id", id))

		render.JSON(w, r, ResponsesignUpOk{
			Response: resp.OK(),
			Id:       id,
		})
	}
}

func (h *Handler) signIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.auth.signIn"

		log := h.log
		log.With(
			zap.String("op", op),
			zap.String("required_id", middleware.GetReqID(r.Context())),
		)

		var input signInInput

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

		token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
		if err != nil {
			log.Error("failed to generate token", zap.Error(err))
			render.JSON(w, r, resp.Error("failed to generate token"))

			return
		}

		log.Info("token has been got", zap.String("token", token))

		render.JSON(w, r, ResponsesignInok{
			Response: resp.OK(),
			Token:    token,
		})
	}
}
