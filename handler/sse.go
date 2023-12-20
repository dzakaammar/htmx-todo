package handler

import (
	"io"
	"net/http"
	"strings"
	"time"
)

func SSE() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f, ok := w.(http.Flusher)
		if !ok {
			return
		}

		h := w.Header()
		h.Set("Content-Type", "text/event-stream")
		h.Set("Connection", "keep-alive")
		h.Set("Cache-Control", "no-cache")

	L:
		for {
			select {
			case <-r.Context().Done():
				break L
			case <-time.After(2 * time.Second):
				_, _ = io.WriteString(w, ": heart beat\n\n")
				f.Flush()
			case <-ch:
				str := strings.Builder{}
				str.WriteString("event: newTodos\n")
				str.WriteString("data: newTodos\n\n")
				_, _ = w.Write([]byte(str.String()))
				f.Flush()
			}
		}
	}
}
