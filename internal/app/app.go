package app

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"

	"github.com/salmanshahzad/go-htmx-todo/internal/database"
	"github.com/salmanshahzad/go-htmx-todo/internal/utils"
)

type Application struct {
	assets *fs.FS
	db     *database.Queries
	env    *utils.Environment
	rdb    *redis.Client
	router *chi.Mux
	sm     *scs.SessionManager
	views  *template.Template
}

func NewApplication(assets *fs.FS, db *database.Queries, env *utils.Environment, rdb *redis.Client, sm *scs.SessionManager, views *template.Template) *Application {
	app := Application{
		assets: assets,
		db:     db,
		env:    env,
		rdb:    rdb,
		router: chi.NewRouter(),
		sm:     sm,
		views:  views,
	}

	app.router.Use(middleware.Logger)
	app.router.Use(middleware.Recoverer)
	app.router.Use(middleware.GetHead)
	app.router.Use(sm.LoadAndSave)
	app.router.Get("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(*assets))).ServeHTTP)
	app.router.Mount("/", app.newHomeRouter())
	app.router.Mount("/session", app.newSessionRouter())
	app.router.Mount("/todos", app.newTodoRouter())
	app.router.Mount("/user", app.newUserRouter())
	return &app
}

func (app *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w, r)
}

func (app *Application) GracefulShutdown() {
	app.rdb.Close()
	log.Println("Disconnected from Redis")
}
