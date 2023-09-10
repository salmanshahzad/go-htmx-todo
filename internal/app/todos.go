package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/salmanshahzad/go-htmx-todo/internal/database"
	"github.com/salmanshahzad/go-htmx-todo/internal/utils"
)

func (app *Application) newTodoRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(app.verifyAuth)
	r.Get("/", app.handleTodosView)
	r.Post("/", app.handleCreateTodo)
	r.Put("/{id}/complete", app.handleCompleteTodo)
	r.Put("/{id}/name", app.handleUpdateTodo)
	r.Delete("/{id}", app.handleDeleteTodo)
	return r
}

func (app *Application) handleTodosView(w http.ResponseWriter, r *http.Request) {
	userId := app.sm.GetInt32(r.Context(), "userId")
	data := make(map[string]any)
	todos, err := app.db.GetTodos(r.Context(), userId)
	if err != nil {
		data["Error"] = "There was an error getting your todos"
	} else {
		data["Todos"] = todos
	}

	app.renderLayout(w, "Todos", "todos.html", data)
}

func (app *Application) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	userId := app.sm.GetInt32(r.Context(), "userId")

	name := utils.FormValue(r, "name")
	if len(name) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	todo, err := app.db.CreateTodo(r.Context(), database.CreateTodoParams{
		Name:   name,
		UserID: userId,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("There was an error adding the todo"))
		return
	}

	app.views.ExecuteTemplate(w, "todoItem.html", todo)
}

func (app *Application) handleCompleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := utils.IDParam(w, r, "id")
	if todoId == 0 {
		return
	}

	completed := utils.FormValue(r, "completed") == "true"
	todo, err := app.db.UpdateTodoCompleted(r.Context(), database.UpdateTodoCompletedParams{
		ID:        todoId,
		Completed: completed,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("There was an error updating the todo"))
		return
	}

	app.views.ExecuteTemplate(w, "todoItem.html", todo)
}

func (app *Application) handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoId := utils.IDParam(w, r, "id")
	if todoId == 0 {
		return
	}

	name := utils.FormValue(r, "name")
	if len(name) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Todo cannot be empty"))
		return
	}

	todo, err := app.db.UpdateTodoName(r.Context(), database.UpdateTodoNameParams{
		ID:   todoId,
		Name: name,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("There was an error updating the todo"))
		return
	}

	app.views.ExecuteTemplate(w, "todoItem.html", todo)
}

func (app *Application) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := utils.IDParam(w, r, "id")
	if todoId == 0 {
		return
	}

	if err := app.db.DeleteTodo(r.Context(), todoId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("There was an error deleting the todo"))
		return
	}
}
