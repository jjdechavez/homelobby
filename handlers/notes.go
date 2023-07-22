package handlers

import (
	"net/http"

	"github.com/jjdechavez/homelobby/views"
)

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	var notesView *views.View
	notesView = views.NewView("app", "views/notes.html")
	notesView.Render(w, map[string]interface{}{"name": "Notes", "msg": "hello world"})
}
