package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qndaa/pack-calculator/internal/model/dto"
	"github.com/qndaa/pack-calculator/internal/usecase/interfaces"
)

type Handler struct {
	calculatorUseCase    interfaces.Calculator
	packRetrieverUseCase interfaces.PackRetriever
	packRemoverUseCase   interfaces.PackRemover
	packCreatorUseCase   interfaces.PackCreator
}

func NewHandler(
	calculatorUseCase interfaces.Calculator,
	packRetrieverUseCase interfaces.PackRetriever,
	packRemoverUseCase interfaces.PackRemover,
	packCreatorUseCase interfaces.PackCreator,
) *Handler {
	return &Handler{
		calculatorUseCase:    calculatorUseCase,
		packRetrieverUseCase: packRetrieverUseCase,
		packRemoverUseCase:   packRemoverUseCase,
		packCreatorUseCase:   packCreatorUseCase,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /calculate", h.calculate)
	mux.HandleFunc("GET /packs", h.getPacks)
	mux.HandleFunc("DELETE /packs/{size}", h.deletePack)
	mux.HandleFunc("POST /packs/{size}", h.createPack)

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
	result := h.packRetrieverUseCase.GetPacks(r.Context())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) deletePack(w http.ResponseWriter, r *http.Request) {
	sizeStr := r.PathValue("size")
	var size int
	if _, err := fmt.Sscanf(sizeStr, "%d", &size); err != nil {
		http.Error(w, "invalid pack size", http.StatusBadRequest)
		return
	}

	if err := h.packRemoverUseCase.RemovePack(r.Context(), size); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) createPack(w http.ResponseWriter, r *http.Request) {
	sizeStr := r.PathValue("size")
	var size int
	if _, err := fmt.Sscanf(sizeStr, "%d", &size); err != nil {
		http.Error(w, "invalid pack size", http.StatusBadRequest)
		return
	}

	if err := h.packCreatorUseCase.CreatePack(r.Context(), size); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
