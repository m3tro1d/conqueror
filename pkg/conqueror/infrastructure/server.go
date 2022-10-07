package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
	subRouter.HandleFunc("/timetable", s.handleTimetable).Methods(http.MethodGet)

	return router
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index")
}

func (s *Server) handleTimetable(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Timetable")
}
