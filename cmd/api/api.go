package api

import (
	"log"
	"net/http"

	"github.com/Mikhail-Tal63/Rooksy/internal/config"
	"github.com/Mikhail-Tal63/Rooksy/internal/middleware"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type API struct {
	DB   *mongo.Database
	Port string
}

func NewAPIServer(DB *mongo.Database, Port string) *API {
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
