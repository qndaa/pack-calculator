package server

import (
	"encoding/json"
	"net/http"

	"github.com/qndaa/pack-calculator/internal/model/dto"
	"github.com/qndaa/pack-calculator/internal/usecase/interfaces"
)

type Handler struct {
	calculatorUseCase interfaces.Calculator
}

func NewHandler(calculatorUseCase interfaces.Calculator) *Handler {
	return &Handler{
		calculatorUseCase: calculatorUseCase,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /calculate", h.calculate)
	mux.HandleFunc("GET /packs", h.getPacks)

	// Serve static UI
	fs := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fs)
}

func (h *Handler) calculate(w http.ResponseWriter, r *http.Request) {
	var req dto.CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.calculatorUseCase.Calculate(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) getPacks(w http.ResponseWriter, r *http.Request) {
	result, err := h.calculatorUseCase.GetPacks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
