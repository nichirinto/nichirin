package core

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewApp() *Nichirin {
	r := chi.NewRouter()

	return &Nichirin{
		Engine: r,
	}
}

func (r *Nichirin) Listen(addr string) {
	http.ListenAndServe(addr, r.Engine)
}
