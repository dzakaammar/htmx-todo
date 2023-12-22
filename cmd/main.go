package main

import (
	"net/http"

	htmxtodo "github.com/dzakaammar/htmx-todo"
	"github.com/dzakaammar/htmx-todo/handler"
	"github.com/dzakaammar/htmx-todo/static"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	app := chi.NewRouter()

	app.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := htmxtodo.InjectNameToCtx(r.Context(), "FooBar")

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	app.Use(cors.AllowAll().Handler)
	static.Handler(app)

	app.Get("/", handler.IndexPage)
	app.Get("/todos", handler.GetAllTodos)
	app.Post("/todos", handler.CreateTodo)
	app.Delete("/todos/{id}", handler.DeleteTodo)
	app.Put("/todos/{id}", handler.UpdateTodo)
	app.Get("/sse", handler.SSE())

	if err := http.ListenAndServe(":8000", app); err != nil {
		panic(err)
	}
}
