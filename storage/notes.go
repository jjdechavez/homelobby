package storage

import "github.com/jmoiron/sqlx"

type (
	NoteStorage struct {
		DB *sqlx.DB
	}

	Note struct {
		Id     string
		Detail string
	}

	NoteInput struct {
		Detail string
	}
)

func InitNoteStorage(db *sqlx.DB) *NoteStorage {
	return &NoteStorage{DB: db}
}

func (storage *NoteStorage) CreateNoteTable() {
	schema := `CREATE TABLE IF NOT EXISTS notes (
    id INTEGER NOT NULL PRIMARY KEY,
    detail TEXT
  );`

	_, err := storage.DB.Exec(schema)
	if err != nil {
		panic(err)
	}
}

func (storage *NoteStorage) CreateNote(noteInput *NoteInput) (int, error) {
	statement := "INSERT INTO notes (detail) VALUES (?)"
	res, err := storage.DB.Exec(statement, noteInput.Detail)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (storage *NoteStorage) GetAllNotes() ([]Note, error) {
	notes := []Note{}
	err := storage.DB.Select(&notes, "SELECT id, detail from notes")
	if err != nil {
		return notes, err
	}

	return notes, nil
}

func (storage *NoteStorage) GetNoteById(noteId string) (Note, error) {
	note := Note{}
	err := storage.DB.Get(&note, "SELECT id, detail from notes WHERE id = ?", noteId)
	if err != nil {
		return note, err
	}

	return note, nil
}

func (storage *NoteStorage) UpdateNoteById(noteId string, detail string) error {
	_, err := storage.DB.Exec("UPDATE notes SET detail = ? WHERE id = ?", detail, noteId)
	if err != nil {
		return err
	}

	return nil
}
