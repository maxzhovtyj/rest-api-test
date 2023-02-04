package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-test/internal/apperrors"
	"rest-api-test/internal/config"
	"rest-api-test/internal/handlers"
	"rest-api-test/pkg/logging"
)

const (
	usersUrl = "/users"
	userUrl  = "/user/:id"
)

type handler struct {
	logger  *logging.Logger
	cfg     *config.Config
	service *Service
}

func NewUserHandler(logger *logging.Logger, cfg *config.Config, service *Service) handlers.Handler {
	return &handler{
		logger:  logger,
		cfg:     cfg,
		service: service,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersUrl, apperrors.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersUrl, apperrors.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userUrl, apperrors.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userUrl, apperrors.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userUrl, apperrors.Middleware(h.PartUpdateUser))
	router.HandlerFunc(http.MethodDelete, userUrl, apperrors.Middleware(h.DeleteUser))
	router.HandlerFunc(http.MethodGet, "/invoice", apperrors.Middleware(h.InvoicePdf))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	marshal, err := json.Marshal(map[string]any{
		"id": 1,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(marshal)
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("get user by uuid")))
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	user, err := h.service.CreateOne(context.Background(), CreateUserDTO{
		Email:    "",
		Username: "",
		Password: "",
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("create user %s", user.Username)))
	return nil
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("update user")))
	return nil
}
func (h *handler) PartUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("partially update user")))
	return nil
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("delete user")))
	return nil
}

func (h *handler) InvoicePdf(w http.ResponseWriter, r *http.Request) error {
	return nil
}
