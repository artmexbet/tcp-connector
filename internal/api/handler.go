package api

import (
	"context"
	"encoding/json"
	"net/http"
	"tcp-conntector/internal/checker"
	"time"
)

type CheckRequest struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

// PortChecker defines the behavior required by the handler.
type PortChecker interface {
	Check(ctx context.Context, ip string, port int) checker.PortStatus
}

type Handler struct {
	checker PortChecker
}

func NewHandler(c PortChecker) *Handler {
	return &Handler{
		checker: c,
	}
}

func (h *Handler) HandleCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CheckRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.IP == "" || req.Port == 0 {
		http.Error(w, "IP and Port are required", http.StatusBadRequest)
		return
	}

	// Context with timeout
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	result := h.checker.Check(ctx, req.IP, req.Port)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		// In a real app, we might want to log this error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
