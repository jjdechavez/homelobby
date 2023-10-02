package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jjdechavez/homelobby/storage"
	"github.com/jjdechavez/homelobby/views"
)

type NotesHandler struct {
	Storage *storage.NoteStorage
}

func InitNotesHandler(storage *storage.NoteStorage) *NotesHandler {
	return &NotesHandler{Storage: storage}
}

func (n *NotesHandler) NoteRoutes() chi.Router {
	router := chi.NewRouter()
	router.Get("/", n.IndexHandler)
  router.Get("/create", n.CreateNoteHandler)
  router.Post("/", n.StoreNoteHandler)

	return router
}

func (n *NotesHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	var notesView *views.View
	data := map[string]interface{}{"name": "Notes", "msg": "hello world"}

	notesView = views.InitView(r, "app", "views/notes.html")
	notesView.Render(w, data)
}

func (n *NotesHandler) CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notesView *views.View
	data := map[string]interface{}{"name": "Notes", "msg": "hello world"}

	notesView = views.InitView(r, "app", "views/create_note.html")
	notesView.Render(w, data)
}

func (n *NotesHandler) StoreNoteHandler(w http.ResponseWriter, r *http.Request) {
	detail := r.FormValue("detail")
  note := &storage.NoteInput{
    Detail: detail,
  }
	n.Storage.CreateNote(note)

	http.Redirect(w, r, "/notes", http.StatusSeeOther)
}
