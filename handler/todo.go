package handler

import (
	"net/http"
	"sync/atomic"

	htmxtodo "github.com/dzakaammar/htmx-todo"
	"github.com/dzakaammar/htmx-todo/view/component"
	"github.com/dzakaammar/htmx-todo/view/page"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

var todos = []*htmxtodo.Todo{
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

var counter atomic.Int32

func init() {
	_ = counter.Add(3)
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

	todos = append(todos, &htmxtodo.Todo{
		ID:     int(counter.Add(1)),
		Title:  r.Form.Get("title"),
		IsDone: false,
	})

	w.Header().Set("HX-Trigger", "newTodos")
	c := component.TodoForm()
	_ = c.Render(r.Context(), w)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := cast.ToInt(chi.URLParam(r, "id"))

	todos = lo.Filter(todos, func(item *htmxtodo.Todo, index int) bool {
		return item.ID != id
	})

	render.Status(r, http.StatusOK)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := cast.ToInt(chi.URLParam(r, "id"))

	var target *htmxtodo.Todo
	for _, todo := range todos {
		if todo.ID == id {
			todo.IsDone = !todo.IsDone
			target = todo
		}
	}

	btn := component.DoneButton(target)
	_ = btn.Render(r.Context(), w)
}
