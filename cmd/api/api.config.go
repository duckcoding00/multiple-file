package api

import (
	"log"
	"net/http"

	"github.com/duckcoding00/multiple-file/internal/handler"
	"github.com/duckcoding00/multiple-file/lib/utils"
	"github.com/gorilla/mux"
)

type Application struct {
	router *mux.Router
	config AppConfig
}

type AppConfig struct {
	h    handler.Handler
	addr string
}

func NewApp(config AppConfig) *Application {
	return &Application{
		router: mux.NewRouter(),
		config: config,
	}
}

func (a *Application) RegisterRouter() {
	apiRouter := a.router.PathPrefix("/api/v1").Subrouter()

	// check
	apiRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteOk(w, http.StatusOK, "Connection OK")
	})

	apiRouter.HandleFunc("/upload", a.config.h.File.Upload).Methods("POST")
}

func (a *Application) Run() {
	log.Printf("server running on :%s", a.config.addr)
	if err := http.ListenAndServe(a.config.addr, a.router); err != nil {
		log.Fatalf("failed to starting server :%v", err)
	}
}
