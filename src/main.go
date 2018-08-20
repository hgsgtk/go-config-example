package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Khigashiguchi/go-config-example/src/config"
	_ "github.com/go-sql-driver/mysql" // mysql driverを使うため
	"os"
)

// NewDB return database global connection handle.
func NewDB(conf config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Name))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

type Handler struct {
	DB *sql.DB
}

// GetPostsHandler handle GET request to /posts.
func (h *Handler) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	// use database here.
}

func main() {
	var err error

	// Get configuration
	conf, err := config.NewConfig()
	if err != nil {
		panic(err.Error())
	}

	// Get database Handle
	db, err := NewDB(conf.DB)
	if err != nil {
		panic(err.Error())
	}

	// Router
	r := mux.NewRouter()
	h := Handler{DB: db}
	r.Methods("GET").Path("/posts").HandlerFunc(h.GetPostsHandler)
	r.Methods("GET").Path("/.healthcheck").HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Serve HTTP service
	fmt.Fprint(os.Stdout, ">> Start to listen http server post :8080\n")
	if err = http.ListenAndServe(":8080", r); err != nil {
		panic(err.Error())
	}
}
