package routes

import (
	"html/template"
	"net/http"

	"github.com/HRMS/db"
	"github.com/HRMS/models"
	"github.com/HRMS/sessions"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//App for router package
type App struct {
	AppSpec     models.App
	template    *template.Template
	handler     *mux.Router
	log         *logrus.Entry
	sessionCtrl sessions.SessInterface
	crud        db.Storage
}

//New is the starting function
func New(log *logrus.Entry, CN models.App, session sessions.SessInterface, crud db.Storage) *App {
	log.Info("Creating a route object")
	t := template.Must(template.ParseGlob("web/static/*.html"))
	handler := mux.NewRouter().StrictSlash(true)
	r := &App{
		AppSpec:     CN,
		template:    t,
		handler:     handler,
		log:         log,
		sessionCtrl: session,
		crud:        crud,
	}

	handler.Methods("GET").HandlerFunc(r.index).Path("/")
	handler.Methods("POST").HandlerFunc(r.login).Path("/login")
	handler.Methods("GET").HandlerFunc(r.mwareLoggedIn(r.logout)).Path("/logout")
	handler.Methods("GET").HandlerFunc(r.mwareLoggedIn(r.home)).Path("/home")
	handler.Methods("GET").HandlerFunc(r.mwareLoggedIn(r.education)).Path("/education")
	handler.Methods("GET").HandlerFunc(r.mwareLoggedIn(r.leaves)).Path("/leaves")
	handler.Methods("GET").HandlerFunc(r.mwareLoggedIn(r.hradmin)).Path("/hradmin")

	api := handler.PathPrefix("/api").Subrouter().StrictSlash(true)
	api.Methods("GET").Path("/sessions").HandlerFunc(r.sessionAPI)
	api.Methods("GET").Path("/designation/{id:[0-9]+}").HandlerFunc(r.designations)

	fs := http.FileServer(http.Dir("."))
	handler.PathPrefix("/web").Handler(fs)

	log.Info("Created a route object")
	return r
}

//Begin is for starting the router module
func (r *App) Begin() error {
	defer func() {
		if err := recover(); err != nil {
			r.log.Errorf("=========ERROR=============\n%v\n", err)
		}
	}()
	if err := http.ListenAndServe("0.0.0.0:8080", r.handler); err != nil {
		return err
	}
	return nil
}
