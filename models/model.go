package models

import "time"

type NoteSnippet struct {
	Id         int    `json:"id"`
	PageUrl    string `json:"pageUrl"`
	Content    string `json:"note"`
	NoteId     int    `json:"noteId"`     // the identifier of the note to which this will be added
	NoteBookId int    `json:"noteBookId"` // the identifier of the note to which this will be added
}

type Note struct {
	Id               int       `json:"id"`
	PageUrl          string    `json:"pageUrl"` // list of urls
	Tag              string    `json:"tag"`     // list of tags
	Content          string    `json:"note"`
	NoteId           int       `json:"noteId"`     // the identifier of the note to which this will be added
	NoteBookId       int       `json:"noteBookId"` // the identifier of the note to which this will be added
	LastModifiedDate time.Time `json:"lastModifiedDate"`
}
