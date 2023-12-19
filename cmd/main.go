package main

import (
	"net/http"

	"github.com/dzakaammar/htmx-todo/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	app := chi.NewRouter()

	app.Get("/", handler.IndexPage)
	app.Get("/todos", handler.GetAllTodos)
	app.Post("/todos", handler.CreateTodo)

	if err := http.ListenAndServe(":8000", app); err != nil {
		panic(err)
	}
}
