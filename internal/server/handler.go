package server

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /calculate", h.calculate) // TODO: implement calculation handler

	// Serve static UI
	fs := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fs)
}

func (h *Handler) calculate(w http.ResponseWriter, r *http.Request) {
	// TODO: implement calculation logic
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not implemented yet"))
}
