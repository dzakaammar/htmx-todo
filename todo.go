package htmxtodo

import (
	"sync/atomic"

	"github.com/samber/lo"
)

var todos = []*Todo{
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
	_ = counter.Add(int32(len(todos)))
}

type Todo struct {
	ID     int
	Title  string
	IsDone bool
}

func NewTodo(title string) *Todo {
	t := &Todo{
		ID:     int(counter.Add(1)),
		Title:  title,
		IsDone: false,
	}
	todos = append(todos, t)
	return t
}

func GetTodos() []*Todo {
	return todos
}

func DeleteTodo(id int) {
	todos = lo.Filter(todos, func(item *Todo, index int) bool {
		return item.ID != id
	})
}

func UpdateTodo(id int) *Todo {
	var target *Todo
	for _, todo := range todos {
		if todo.ID == id {
			todo.IsDone = !todo.IsDone
			target = todo
		}
	}

	return target
}
