package handlers

import (
	"log"
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
	router.Get("/{id}/edit", n.EditNoteHandler)
	router.Put("/{id}", n.UpdateNoteHandler)

	return router
}

type IndexNoteHandlerResponse struct {
	Name  string
	Notes []storage.Note
}

func (n *NotesHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	var notesView *views.View

	data, err := n.Storage.GetAllNotes()
	if err != nil {
		log.Println("NoteIndexHandler Error: ", err)
	}

	response := IndexNoteHandlerResponse{
		Name:  "Notes",
		Notes: data,
	}

	notesView = views.InitView(r, "app", "views/notes.html")
	notesView.Render(w, map[string]interface{}{"Name": response.Name, "Notes": response.Notes})
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

func (n *NotesHandler) EditNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notesView *views.View
	noteId := chi.URLParam(r, "id")

	note, err := n.Storage.GetNoteById(noteId)
	if err != nil {
		log.Println("NoteEditHandler Error: ", err)
	}
	data := map[string]interface{}{"name": "Edit Note"}
	data["Note"] = note

	notesView = views.InitView(r, "app", "views/edit_note.html")
	notesView.Render(w, data)
}

func (n *NotesHandler) UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	noteId := chi.URLParam(r, "id")
	detail := r.FormValue("detail")

	err := n.Storage.UpdateNoteById(noteId, detail)
	if err != nil {
		log.Println("UpdateNoteHandler Error: ", err)
	}

	http.Redirect(w, r, "/notes", http.StatusSeeOther)
}
