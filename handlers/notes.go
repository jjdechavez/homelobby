package handlers

import (
	"net/http"

	"github.com/jjdechavez/homelobby/views"
)

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	var notesView *views.View
	data := map[string]interface{}{"name": "Notes", "msg": "hello world"}

	notesView = views.InitView(r, "app", "views/notes.html")
	notesView.Render(w, data)
}
