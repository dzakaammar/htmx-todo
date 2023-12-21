package main

import (
	"net/http"

	"github.com/dzakaammar/htmx-todo/handler"
	"github.com/dzakaammar/htmx-todo/static"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	app := chi.NewRouter()
	app.Use(cors.AllowAll().Handler)
	static.Handler(app)

	app.Get("/", handler.IndexPage)
	app.Get("/todos", handler.GetAllTodos)
	app.Post("/todos", handler.CreateTodo)
	app.Delete("/todos/{id}", handler.DeleteTodo)
	app.Put("/todos/{id}", handler.UpdateTodo)
	handler.WS(app)

	if err := http.ListenAndServe(":8000", app); err != nil {
		panic(err)
	}
}
