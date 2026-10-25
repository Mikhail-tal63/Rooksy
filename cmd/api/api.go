package api

import (
	"log"
	"net/http"

	"github.com/Mikhail-Tal63/Rooksy/internal/config"
	"github.com/Mikhail-Tal63/Rooksy/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type API struct {
	DB   *pgxpool.Pool
	Port string
}

func NewAPIServer(DB *pgxpool.Pool, Port string) *API {
	return &API{
		DB:   DB,
		Port: Port,
	}
}

func (s *API) Run(cfg *config.Config) error {

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	/*protected router*/
	protacted := apiRouter.PathPrefix("").Subrouter()
	protacted.Use(middleware.AuthMiddleWare(cfg.JwtSecret))
	log.Printf("App is listening on %s", s.Port)
	addr := ":" + s.Port
	return http.ListenAndServe(addr, router)
}
