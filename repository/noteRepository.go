package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/couchbase/gocb.v1"
	"notebook/common"
	"notebook/models"
	"os"
	"strconv"
	"time"
)

var bucket *gocb.Bucket

type NoteRepository struct {
}

func init() {
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ = cluster.OpenBucket("default", "")
}

func (repo *NoteRepository) UpsertNote(note *models.Note) (err error) {

	key := strconv.Itoa(note.NoteBookId)
	//content := noteSnippet.Content

	if _, err := bucket.Upsert(key, note, 0); err != nil {
		common.Error.Println("Note could not be saved", err.Error())
		return err
	}
	return err
}

// https://developer.couchbase.com/documentation/server/4.5/sdk/go/start-using-sdk.html
func (repo *NoteRepository) AppendNote(noteSnippet *models.NoteSnippet) (err error) {

	key := strconv.Itoa(noteSnippet.NoteBookId)
	//content := noteSnippet.Content

	var note models.Note
	if _, err := bucket.Get(key, &note); err != nil {

		note = models.Note{
			PageUrl:          noteSnippet.PageUrl,
			Content:          noteSnippet.Content,
			LastModifiedDate: time.Now(),
		}
		common.Error.Println("Note does not exists creating new", err)
		if _, err := bucket.Insert(key, note, 0); err != nil {
			common.Error.Println("Note could not be added ", err)
			return err
		}

	} else {
		note.Content = note.Content + noteSnippet.Content
		note.LastModifiedDate = time.Now()
		common.Info.Println("Note exists, appending to it")
		if _, err := bucket.Upsert(strconv.Itoa(noteSnippet.NoteBookId), note, 0); err != nil {
			common.Error.Println("Note could not be saved", err.Error())
			return err
		}

	}
	return err
}

// Create a new noteSnippet
// Refer: https://stackoverflow.com/questions/7151261/append-to-a-file-in-go
func (repo *NoteRepository) AddNoteToFile(noteSnippet *models.NoteSnippet) (err error) {

	if _, err := os.Stat("/Users/a.nigam/Documents/workspace/gospace/src/notebook/temp.txt"); err != nil {
		common.Error.Println("File does not exist", err)
	}
	file, err := os.OpenFile("/Users/a.nigam/Documents/workspace/gospace/src/notebook/temp.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		common.Error.Println("NoteSnippet could not be created ", err)
		return
	}

	defer file.Close()

	if _, err = file.WriteString("\n"); err != nil {
		panic(err)
	}

	if _, err = file.WriteString(noteSnippet.PageUrl + "\n"); err != nil {
		panic(err)
	}
	if _, err = file.WriteString(noteSnippet.Content + "\n"); err != nil {
		panic(err)
	}

	return
}
