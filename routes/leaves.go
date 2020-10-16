package routes

import (
	"net/http"

	"github.com/HRMS/db"
	"github.com/HRMS/models"
)

func (r *App) leaves(rw http.ResponseWriter, req *http.Request) {
	var u models.Webstruct
	u.AppSpec = r.AppSpec
	uid, err := r.getuidfromreq(req)
	if err != nil {
		r.log.Error("Error while processing request: ", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
	user, err := r.crud.BuildUser(uid, db.LEAVES)
	u.User = user
	if err != nil {
		r.log.Error("Error while processing request: ", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}

	if err := r.template.ExecuteTemplate(rw, "leaves.html", u); err != nil {
		r.log.Error("Error while processing request: ", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}

}
