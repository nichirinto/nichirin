package core

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/nichirinto/nichirin/framework/lib/logger"
	"net/http"
)

func Init() {
	_ = godotenv.Load()
}

func NewApp(addr string) *Nichirin {

	r := chi.NewRouter()

	return &Nichirin{
		Engine: r, Address: addr,
	}
}

func (s *Nichirin) StartServer() {
	logger.Log.Info("Server started")
	logger.Log.Info(s.Address)

	if err := http.ListenAndServe(s.Address, s.Engine); err != nil {
		logger.Log.Error(err.Error())
		panic(err)
	}
}
