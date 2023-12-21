package handler

import (
	"net/http"

	htmxtodo "github.com/dzakaammar/htmx-todo"
	"github.com/dzakaammar/htmx-todo/view/component"
	"github.com/dzakaammar/htmx-todo/view/page"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/spf13/cast"
)

var ch chan struct{}

func init() {
	ch = make(chan struct{})
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	page := page.Index(htmxtodo.GetTodos())

	_ = page.Render(r.Context(), w)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	c := component.TodoList(htmxtodo.GetTodos())

	_ = c.Render(r.Context(), w)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	_ = htmxtodo.NewTodo(r.Form.Get("title"))
	ch <- struct{}{}

	c := component.TodoForm()
	_ = c.Render(r.Context(), w)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := cast.ToInt(chi.URLParam(r, "id"))

	htmxtodo.DeleteTodo(id)

	render.Status(r, http.StatusOK)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := cast.ToInt(chi.URLParam(r, "id"))

	target := htmxtodo.UpdateTodo(id)

	btn := component.DoneButton(target)
	_ = btn.Render(r.Context(), w)
}
