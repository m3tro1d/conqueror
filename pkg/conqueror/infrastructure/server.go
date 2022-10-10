package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"conqueror/pkg/conqueror/domain"
)

func NewServer(dependencyContainer *DependencyContainer) *Server {
	return &Server{
		dependencyContainer: dependencyContainer,
	}
}

type Server struct {
	dependencyContainer *DependencyContainer
}

func (s *Server) GetRouter() *mux.Router {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.HandleFunc("", s.handleIndex).Methods(http.MethodGet)

	return router
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	repo, err := s.dependencyContainer.userRepository()
	if err != nil {
		serverError(err, w, r)
	}

	err = repo.Store(domain.NewUser(
		1,
		"test",
		"test",
		"test",
	))
	if err != nil {
		serverError(err, w, r)
	}
}

func serverError(err error, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}
