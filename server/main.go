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
	router.HandleFunc("/ping", s.pingHandler).Methods("GET")
	router.HandleFunc("/apps", s.createAppHandler).Methods("POST")
	router.HandleFunc("/apps/{appId}/nodes", s.createNodeHandler).Methods("POST")
	router.HandleFunc("/apps/{appId}/deployments", s.createDeploymentHandler).Methods("POST")
	router.HandleFunc("/apps/{appId}/deployments/{deploymentId}", s.updateDeploymentHandler).Methods("PUT")

	if err := http.ListenAndServe(":"+s.config.Port, router); err != nil {
		panic(err)
	}
}

type createAppRequest struct {
	Name string `json:"name"`
}

func (s *Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG"))
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

func (s *Server) createNodeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["appId"]
	node, err := s.app.CreateNode(appID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.writeJSON(w, node)
}

type updateDeploymentRequest struct {
	NodeID string `json:"nodeId"`
	Status string `json:"status"`
}

func (s *Server) updateDeploymentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["appId"]
	deploymentID := vars["deploymentId"]

	var payload updateDeploymentRequest
	if ok := s.readJSON(r, &payload); !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	deployment, err := s.app.UpdateDeployment(appID, deploymentID, payload.NodeID, payload.Status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.writeJSON(w, deployment)
}

type createDeploymentRequest struct {
	GitURL string `json:"gitURL"`
}

func (s *Server) createDeploymentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["appId"]

	var payload createDeploymentRequest
	if ok := s.readJSON(r, &payload); !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	deployment, err := s.app.CreateDeployment(appID, payload.GitURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.writeJSON(w, deployment)
}
