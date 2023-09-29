package server_handler

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"product_move/internal/adapter/db_handler"
	"product_move/internal/common/method"
	"product_move/internal/controller"
)

type APIServer struct {
	listenPort uint16
	router     *mux.Router
	db         *db_handler.MySqlStore
}

func NewAPIServer(listenPort uint16, db *db_handler.MySqlStore) (*APIServer, error) {
	router := mux.NewRouter()
	return &APIServer{
		listenPort: listenPort,
		router:     router,
		db:         db,
	}, nil
}

func (s *APIServer) GetDB() *sql.DB {
	return s.db.GetDB()
}

func (s *APIServer) GetRouter() *mux.Router {
	return s.router
}

func (s *APIServer) Run() error {
	categoryController := controller.NewCategoryController(s.GetDB())
	s.AddRoutes("/", categoryController.IndexHandler, method.GET)
	log.Printf("Server starting")
	log.Printf("Open your browser on http://localhost:%d/", s.listenPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.listenPort), s.router)
	if err != nil {
		return err
	}
	return nil
}

func (s *APIServer) AddRoutes(path string, handler func(http.ResponseWriter, *http.Request), method method.Method) {
	s.router.HandleFunc(path, handler).Methods(string(method))
}
