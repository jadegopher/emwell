package handlers

import (
	"net/http"

	"emwell/internal/core/link"
	"emwell/internal/logger"
)

type Handler struct {
	logger      logger.ILogger
	linkService *link.Service
}

func NewHandler(logger logger.ILogger, linkService *link.Service) *Handler {
	return &Handler{
		logger:      logger,
		linkService: linkService,
	}
}

func (h *Handler) GetEmotionalStatistics(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error try again"))
	}
	pwd := r.Form.Get("password")

	chart, err := h.linkService.GetByPassword(r.Context(), pwd)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error try again"))
		return
	}

	w.WriteHeader(200)
	_, _ = w.Write(chart)
}
