package app

import (
	"bytes"
	"html/template"
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
	home := new(bytes.Buffer)
	app.views.ExecuteTemplate(home, "home.html", nil)

	app.views.ExecuteTemplate(w, "layout.html", map[string]any{
		"Content": template.HTML(home.String()),
		"Title":   "Go + HTMX",
	})
}

func (app *Application) handleSignInView(w http.ResponseWriter, r *http.Request) {
	signin := new(bytes.Buffer)
	app.views.ExecuteTemplate(signin, "auth.html", map[string]any{
		"IsSignUp": false,
	})

	app.views.ExecuteTemplate(w, "layout.html", map[string]any{
		"Content": template.HTML(signin.String()),
		"Title":   "Go + HTMX",
	})
}

func (app *Application) handleSignUpView(w http.ResponseWriter, r *http.Request) {
	signup := new(bytes.Buffer)
	app.views.ExecuteTemplate(signup, "auth.html", map[string]any{
		"IsSignUp": true,
	})

	app.views.ExecuteTemplate(w, "layout.html", map[string]any{
		"Content": template.HTML(signup.String()),
		"Title":   "Go + HTMX",
	})
}
