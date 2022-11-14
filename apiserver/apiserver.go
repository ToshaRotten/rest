package apiserver

import (
	"fmt"
	"net/http"

	"github.com/ToshaRotten/rest/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer default struct contains configuration
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New APIServer
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start server
func (s *APIServer) Start() error {
	err := s.configurateLogger()
	if err != nil {
		return err
	}

	s.configurateRouter()

	err = s.configurateStore()
	if err != nil {
		return err
	}

	s.logger.Info("Starting API Server")

	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *APIServer) configurateLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configurateStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *APIServer) configurateRouter() {
	s.router.HandleFunc("/home", s.handleHome())
}

func (s *APIServer) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "home")
	}
}
