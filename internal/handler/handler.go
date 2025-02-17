package handler

import (
	"net/http"

	"github.com/duckcoding00/multiple-file/internal/service"
)

type Handler struct {
	File interface {
		Upload(w http.ResponseWriter, r *http.Request)
	}
}

func NewHandler() Handler {
	s := service.NewService()
	return Handler{
		File: &FileHandler{
			s: s,
		},
	}
}
