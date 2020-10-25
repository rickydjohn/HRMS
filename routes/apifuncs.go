package routes

import (
	"net/http"

	"github.com/HRMS/db"
	"github.com/gorilla/mux"
)

func (r *App) sessionAPI(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("content-type", "application/json")
	bt, err := r.sessionCtrl.AllSessions()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Write(bt)
	return
}

func (r *App) designations(rw http.ResponseWriter, req *http.Request) {
	val := mux.Vars(req)
	rw.Header().Add("content-type", "application/json")
	bt, err := r.crud.ApiFuncs(val["id"], db.API_DESIGNATIONS)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Write(bt)
	return

}
