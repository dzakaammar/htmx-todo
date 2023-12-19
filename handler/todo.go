package handler

import (
	"net/http"

	htmxtodo "github.com/dzakaammar/htmx-todo"
	"github.com/dzakaammar/htmx-todo/view/component"
	"github.com/dzakaammar/htmx-todo/view/page"
)

var todos = []htmxtodo.Todo{
	{
		ID:     1,
		Title:  "Learn Go",
		IsDone: true,
	},
	{
		ID:     2,
		Title:  "Learn Templ",
		IsDone: false,
	},
	{
		ID:     3,
		Title:  "Learn HTMX",
		IsDone: false,
	},
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	page := page.Index(todos)

	_ = page.Render(r.Context(), w)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	c := component.TodoList(todos)

	_ = c.Render(r.Context(), w)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	todos = append(todos, htmxtodo.Todo{
		ID:     4,
		Title:  r.Form.Get("title"),
		IsDone: false,
	})

	w.Header().Set("HX-Trigger", "newTodos")
	c := component.TodoForm()
	_ = c.Render(r.Context(), w)
}
