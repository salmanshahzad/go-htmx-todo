package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/salmanshahzad/go-htmx-todo/internal/database"
	"github.com/salmanshahzad/go-htmx-todo/internal/utils"
)

func (app *Application) newUserRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", app.handleCreateUser)
	return r
}

func (app *Application) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	username := utils.FormValue(r, "username")
	password := utils.FormValue(r, "password")
	confirmPassword := utils.FormValue(r, "confirmPassword")

	data := map[string]any{
		"IsSignUp":        true,
		"Username":        username,
		"Password":        password,
		"ConfirmPassword": confirmPassword,
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
	if len(confirmPassword) == 0 {
		data["ConfirmPasswordError"] = "Confirm Password is required"
		hasError = true
	}
	if password != confirmPassword {
		data["ConfirmPasswordError"] = "Passwords do not match"
		hasError = true
	}

	if hasError {
		app.views.ExecuteTemplate(w, "auth.html", data)
		return
	}

	count, err := app.db.CountUsersWithUsername(r.Context(), username)
	if err != nil {
		data["Error"] = "An unexpected error occurred"
		app.views.ExecuteTemplate(w, "auth.html", data)
		return
	}

	if count > 0 {
		data["UsernameError"] = "Username already exists"
		app.views.ExecuteTemplate(w, "auth.html", data)
		return
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		data["Error"] = "An unexpected error occurred"
		app.views.ExecuteTemplate(w, "auth.html", data)
		return
	}

	userId, err := app.db.CreateUser(r.Context(), database.CreateUserParams{
		Username: username,
		Password: hashedPassword,
	})
	if err != nil {
		data["Error"] = "An unexpected error occurred"
		app.views.ExecuteTemplate(w, "auth.html", data)
		return
	}

	app.sm.Put(r.Context(), "userId", userId)
	w.Header().Add("HX-Location", "/todos")
}
