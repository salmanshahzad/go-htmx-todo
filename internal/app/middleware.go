package app

import (
	"net/http"
)

func (app *Application) verifyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := app.sm.GetInt32(r.Context(), "userId")
		if userId == 0 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			app.sm.RenewToken(r.Context())
			next.ServeHTTP(w, r)
		}
	})
}

func (app *Application) verifyNoAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := app.sm.GetInt32(r.Context(), "userId")
		if userId == 0 {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/todos", http.StatusSeeOther)
		}
	})
}
