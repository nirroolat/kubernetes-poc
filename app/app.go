package app

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Register mysql driver

	"github.com/gorilla/mux"
	"scm.bluebeam.com/stu/golang-template/apis"
	"scm.bluebeam.com/stu/golang-template/repositories"
	"scm.bluebeam.com/stu/golang-template/services"
)

// App struct
type App struct {
	Router *mux.Router
}

// Initialize app and construct router
func (a *App) Initialize(dsn string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	a.Router = buildRouter(db)
}

// Run app
func (a *App) Run(addr string) {
	http.ListenAndServe(addr, a.Router)
}

func buildRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Register Controllers
	usersRepository := &repositories.UsersRepository{DB: db}
	usersService := &services.UsersService{UsersPersister: usersRepository}
	apis.RegisterUsersResource(r, usersService)
	return r
}
