package transport

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"conqueror/pkg/conqueror/infrastructure"
)

func NewServer(dependencyContainer *infrastructure.DependencyContainer) *Server {
	return &Server{
		dependencyContainer: dependencyContainer,
	}
}

type Server struct {
	dependencyContainer *infrastructure.DependencyContainer
}

func (s *Server) GetRouter() *mux.Router {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.HandleFunc("/user", s.handleIndex).Methods(http.MethodPost)
	subRouter.HandleFunc("/subject", s.handleIndex).Methods(http.MethodPost)
	subRouter.HandleFunc("/subject", s.handleIndex).Methods(http.MethodPut)

	return router
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {

}

func serverError(err error, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}
