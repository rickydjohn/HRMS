package routes

import (
	"net/http"

	"github.com/HRMS/models"
)

func (r *App) buildWebStruct(uname string) models.Webstruct {
	var m models.Webstruct
	m.AppSpec = r.AppSpec
	m.IsAuth = true
	return m
}

func (r *App) mwareLoggedIn(f http.HandlerFunc) http.HandlerFunc {

	return func(rw http.ResponseWriter, req *http.Request) {
		c, err := req.Cookie("ssid")
		if err != nil {
			r.log.Info("Non logged in session: ", err)
			r.template.ExecuteTemplate(rw, "index.html", r)
			return
		}
		err = r.sessionCtrl.IsValid(c.Value, req.RemoteAddr)
		if err != nil {
			r.log.Error("Invalid session ID. Redirecting to home page", err)
			r.template.ExecuteTemplate(rw, "index.html", r)
			return
		}
		f(rw, req)
	}
}

func (r *App) getuidfromreq(req *http.Request) (int, error) {
	cookie, err := req.Cookie("ssid")
	if err != nil {
		return -1, err
	}
	uid, err := r.sessionCtrl.GetUID(cookie.Value)
	if err != nil {
		return -1, err
	}
	return uid, nil

}
