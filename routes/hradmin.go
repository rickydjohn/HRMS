package routes

import (
	"net/http"

	"github.com/HRMS/db"
	"github.com/HRMS/models"
)

func (r *App) hradmin(rw http.ResponseWriter, req *http.Request) {
	var u models.Webstruct
	u.AppSpec = r.AppSpec
	uid, err := r.getuidfromreq(req)
	if err != nil {
		r.log.Error("Error while processing request: ", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
	user, err := r.crud.BuildUser(uid, db.PERSONAL)
	if err != nil {
		r.log.Error("Error while processing request: ", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
	u.User = user
	teams, err := r.crud.ListTeams()
	if err != nil {
		r.log.Error("Error while processing request: ", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
	u.Teams = teams
	if user.Role == "admin" || user.Role == "hradmin" {
		if err := r.template.ExecuteTemplate(rw, "hradmin.html", u); err != nil {
			r.log.Error("Error while processing request: ", err)
			r.template.ExecuteTemplate(rw, "error.html", r)
			return
		}
	} else {
		http.Redirect(rw, req, "/", http.StatusTemporaryRedirect)
		return
	}
}
