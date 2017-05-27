package server

import (
	"github.com/gorilla/mux"
	"github.com/metaops/metaops-api/app"
	"github.com/metaops/metaops-api/config"
	"net/http"
)

type Server struct {
	app    *app.App
	config *config.ServerConfig
}

func New(a *app.App, serverConfig *config.ServerConfig) *Server {

	return &Server{
		app:    a,
		config: serverConfig,
	}
}

func (s *Server) Init() {
	router := mux.NewRouter()
	router.HandleFunc("/apps", s.createAppHandler).Methods("POST")

	if err := http.ListenAndServe(":"+s.config.Port, router); err != nil {
		panic(err)
	}
}

type createAppRequest struct {
	Name string `json:"name"`
}

func (s *Server) createAppHandler(w http.ResponseWriter, r *http.Request) {
	var payload createAppRequest
	if ok := s.readJSON(r, &payload); !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userApp, err := s.app.CreateApp(payload.Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.writeJSON(w, userApp)
}
