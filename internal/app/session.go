package app

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/salmanshahzad/go-htmx-todo/internal/utils"
)

func (app *Application) newSessionRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", app.handleSignIn)
	r.Delete("/", app.handleSignOut)
	return r
}

func (app *Application) handleSignIn(w http.ResponseWriter, r *http.Request) {
	username := utils.FormValue(r, "username")
	password := utils.FormValue(r, "password")

	data := map[string]any{
		"IsSignUp": false,
		"Username": username,
		"Password": password,
	}
	hasError := false

	if len(username) == 0 {
		data["UsernameError"] = "Username is required"
		hasError = true
	}
	if len(password) == 0 {
		data["PasswordError"] = "Password is required"
		hasError = true
	}

	if hasError {
		app.views.ExecuteTemplate(w, "auth.html", data)
		return
	}

	user, err := app.db.GetUserByUsername(r.Context(), username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			data["Error"] = "Invalid username or password"
		} else {
			data["Error"] = "An unexpected error occurred"
		}
		app.views.ExecuteTemplate(w, "auth.html", data)
		return
	}

	validPassword, err := utils.VerifyPassword(password, user.Password)
	if err != nil {
		data["Error"] = "An unexpected error occurred"
		app.views.ExecuteTemplate(w, "auth.html", data)
		return
	}

	if !validPassword {
		data["Error"] = "Invalid username or password"
		app.views.ExecuteTemplate(w, "auth.html", data)
		return
	}

	app.sm.Put(r.Context(), "userId", user.ID)
	w.Header().Add("HX-Location", "/todos")
}

func (app *Application) handleSignOut(w http.ResponseWriter, r *http.Request) {
	if err := app.sm.Destroy(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("HX-Location", "/")
}
