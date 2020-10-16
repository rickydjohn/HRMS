package routes

import (
	"net/http"

	"github.com/HRMS/db"
	"github.com/HRMS/models"
)

func (r *App) education(rw http.ResponseWriter, req *http.Request) {
	var u models.Webstruct
	cookie, _ := req.Cookie("ssid")
	uid, _ := r.sessionCtrl.GetUID(cookie.Value)
	user, err := r.crud.BuildUser(uid, db.BANKEDU)

	if err != nil {
		r.log.Error("Error occurred while building user info", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}

	u.AppSpec = r.AppSpec
	u.IsAuth = true
	u.User = user
	if err := r.template.ExecuteTemplate(rw, "edu.html", u); err != nil {
		r.log.Error("Error while processing request: ", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
}
