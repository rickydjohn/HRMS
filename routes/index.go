package routes

import (
	"net/http"

	"github.com/HRMS/models"
	"github.com/google/uuid"
)

func (r *App) index(rw http.ResponseWriter, req *http.Request) {
	var w models.Webstruct
	//w.CompanyName = r.AppSpec.Company
	w.AppSpec = r.AppSpec
	r.log.Info("Request received ")
	rw.Header().Add("content-type", "text/html")
	cookie, err := req.Cookie("ssid")
	if err == http.ErrNoCookie {
		if err := r.template.ExecuteTemplate(rw, "index.html", w); err != nil {
			r.log.Error("Error occurred while processing request: ", err)
			r.template.ExecuteTemplate(rw, "error.html", w)
			return
		}
	} else {
		uid, err := r.sessionCtrl.GetUID(cookie.Value)
		if err != nil {
			r.log.Error("Unable to find session ID:", uid, ". Must be from an older session. Routing to index page")
			if err := r.template.ExecuteTemplate(rw, "index.html", w); err != nil {
				r.log.Error(err)
				return
			}
		}
		user, err := r.crud.BuildUser(uid)
		if err != nil {
			r.log.Error("Unable to build user info from DB module: ", err)
			if err := r.template.ExecuteTemplate(rw, "error.html", w); err != nil {
				r.log.Error(err)
				return
			}
		}
		w.User = user
		if err := r.template.ExecuteTemplate(rw, "home.html", w); err != nil {
			r.log.Error("Unable to build user info from DB module: ", err)
			if err := r.template.ExecuteTemplate(rw, "error.html", w); err != nil {
				r.log.Error(err)
				return
			}
		}
	}
	r.log.Info("Completed handling request")
	return
}

func (r *App) login(rw http.ResponseWriter, req *http.Request) {
	var u models.Webstruct
	uname := req.FormValue("uname")
	pwd := req.FormValue("passwd")
	v, err := r.crud.Auth(uname, pwd)
	if err != nil {
		r.log.Errorf("Error while processing request :", err)
		rw.WriteHeader(http.StatusUnauthorized)
		r.template.ExecuteTemplate(rw, "index.html", r)
		return
	}
	ssid := uuid.New().String()
	r.sessionCtrl.Create(ssid, uname, req.RemoteAddr, v.UID)
	http.SetCookie(rw, &http.Cookie{Name: "ssid", Value: ssid, MaxAge: 86400})

	user, err := r.crud.BuildUser(v.UID)
	if err != nil {
		r.log.Error("Error occurred while building user info: ", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
	u.AppSpec = r.AppSpec
	u.IsAuth = true
	u.User = user
	if err := r.template.ExecuteTemplate(rw, "home.html", u); err != nil {
		r.log.Error("Error while processing request: ", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
	return
}

//logout function sets cookies to empty and session is deleted
func (r *App) logout(rw http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("ssid")
	if err != nil {
		r.log.Error("Error fetching cooking :", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
	if err := r.sessionCtrl.Delete(c.Value); err != nil {
		r.log.Error("Error deleting cookie :", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
	http.SetCookie(rw, &http.Cookie{Name: "ssid", Value: "", MaxAge: -1})
	if err := r.template.ExecuteTemplate(rw, "index.html", r); err != nil {
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
}

func (r *App) home(rw http.ResponseWriter, req *http.Request) {
	var u models.Webstruct
	cookie, _ := req.Cookie("ssid")
	uid, _ := r.sessionCtrl.GetUID(cookie.Value)
	user, err := r.crud.BuildUser(uid)
	if err != nil {
		r.log.Error("Error occurred while building user info", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
	u.AppSpec = r.AppSpec
	u.IsAuth = true
	u.User = user
	if err := r.template.ExecuteTemplate(rw, "home.html", u); err != nil {
		r.log.Error("Error while processing request: ", err)
		r.template.ExecuteTemplate(rw, "error.html", r)
		return
	}
}
