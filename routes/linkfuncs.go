package routes

import (
	"net/http"

	"github.com/HRMS/models"
)

func (r *App) education(rw http.ResponseWriter, req *http.Request) {
	var u models.Webstruct
	cookie, _ := req.Cookie("ssid")
	uid, _ := r.sessionCtrl.GetUID(cookie.Value)
	user, err := r.crud.Edu(uid)

	if err != nil {
		r.log.Error("Error occurred while building user info", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
	bank, err := r.crud.Bank(uid)
	if err != nil {
		r.log.Error("Error occurred while building user info", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}

	u.AppSpec = r.AppSpec
	u.IsAuth = true
	u.User.Education = user
	u.User.Bank = bank
	if err := r.template.ExecuteTemplate(rw, "edu.html", u); err != nil {
		r.log.Error("Error while processing request: ", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
}

func (r *App) leaves(rw http.ResponseWriter, req *http.Request) {
	var u models.Webstruct
	u.AppSpec = r.AppSpec
	uid, err := r.getuidfromreq(req)
	if err != nil {
		r.log.Error("Error while processing request: ", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
	leaves, err := r.crud.UsedLeaves(uid, r.AppSpec.YearStart)
	u.User.Leaves = leaves
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
