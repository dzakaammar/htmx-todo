package component_test

import (
	"context"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	htmxtodo "github.com/dzakaammar/htmx-todo"
	"github.com/dzakaammar/htmx-todo/view/component"
)

func TestDoneButton(t *testing.T) {
	todo := &htmxtodo.Todo{
		ID:     1,
		Title:  "Testing",
		IsDone: false,
	}

	r, w := io.Pipe()
	go func() {
		_ = component.DoneButton(todo).Render(context.Background(), w)
		w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	if actualTitle := doc.Find(".w-full").Text(); actualTitle != todo.Title {
		t.Fatalf("actual title %s should be equal to %s", actualTitle, todo.Title)
	}

	if id, ok := doc.Find(".contents").Attr("id"); !ok || !strings.Contains(id, strconv.Itoa(todo.ID)) {
		t.Fatalf("the div's id should contain todo ID: %s", id)
	}
}
