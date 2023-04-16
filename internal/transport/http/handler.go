package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CommentService interface{}

type Handler struct {
	Service CommentService
	Router  *mux.Router
	Server  *http.Server
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr:    ":8080",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
