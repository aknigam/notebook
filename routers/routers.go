package routers

import (
	"notebook/common"
	"notebook/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetRoutes(r *mux.Router) {


	r.HandleFunc("/notes/save", controllers.Note.AddNote).Methods("POST")

}

func DoNothing(writer http.ResponseWriter, r *http.Request) {

}

type appHandler struct {
	Handler func(http.ResponseWriter, *http.Request) *error
}

// https://blog.golang.org/error-handling-and-go
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/*
		should create the business objects instead of passing w and r
	*/

	if e := fn.Handler(w, r); e != nil { // e is *appError, not os.Error.
		common.Error.Println("Invalid customer id", e)
		err := *e
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
