package controllers

import (
	"encoding/json"
	"net/http"
	"notebook/common"
	"notebook/models"
	"notebook/repository"
)

type NoteController struct {
	repo repository.NoteRepository
}

var Note NoteController

func (controller *NoteController) UpsertNote(w http.ResponseWriter, r *http.Request) {

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var note models.Note
	json.Unmarshal(body, &note)

	err := controller.repo.UpsertNote(&note)
	if err != nil {
		common.Error.Println("Note could not be created", err)
		return
	}

	common.Info.Println("Call successful", &note)

	w.WriteHeader(http.StatusOK)
	// location header should also be set as per the REST standards
	return
}

func (controller *NoteController) AppendNote(w http.ResponseWriter, r *http.Request) {

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var noteSnippet models.NoteSnippet
	json.Unmarshal(body, &noteSnippet)

	err := controller.repo.AppendNote(&noteSnippet)
	if err != nil {
		common.Error.Println("Note could not be created", err)
		return
	}

	common.Info.Println("Call successful", &noteSnippet)

	w.WriteHeader(http.StatusOK)
	// location header should also be set as per the REST standards
	return
}
