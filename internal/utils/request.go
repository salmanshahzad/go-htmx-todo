package utils

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func FormValue(r *http.Request, key string) string {
	return strings.TrimSpace(r.FormValue(key))
}

func IDParam(w http.ResponseWriter, r *http.Request, key string) int32 {
	id, err := strconv.Atoi(chi.URLParam(r, key))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Invalid ID"))
		return 0
	}
	return int32(id)
}
