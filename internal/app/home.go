package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) newHomeRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(app.verifyNoAuth)
	r.Get("/", app.handleHome)
	r.Get("/signin", app.handleSignInView)
	r.Get("/signup", app.handleSignUpView)
	return r
}

func (app *Application) handleHome(w http.ResponseWriter, r *http.Request) {
	app.renderLayout(w, "Go + HTMX Todo", "home.html", nil)
}

func (app *Application) handleSignInView(w http.ResponseWriter, r *http.Request) {
	app.renderLayout(w, "Go + HTMX Todo", "auth.html", map[string]any{
		"IsSignUp": false,
	})
}

func (app *Application) handleSignUpView(w http.ResponseWriter, r *http.Request) {
	app.renderLayout(w, "Go + HTMX Todo", "auth.html", map[string]any{
		"IsSignUp": true,
	})
}
