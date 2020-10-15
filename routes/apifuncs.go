package routes

import "net/http"

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
