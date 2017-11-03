package models

type NoteSnippet struct {
	Id      int    `json:"id"`
	PageUrl string `json:"pageUrl"`
	Content string `json:"note"`
	NoteId  int    `json:"noteId"` // the identifier of the note to which this will be added
}
