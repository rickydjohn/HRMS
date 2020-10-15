package routes

import (
	"html/template"
	"net/http"

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
	crud        models.Storage
}

//New is the starting function
func New(log *logrus.Entry, CN models.App, session sessions.SessInterface, crud models.Storage) *App {
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

	api := handler.PathPrefix("/api").Subrouter().StrictSlash(true)
	api.Methods("GET").Path("/sessions").HandlerFunc(r.sessionAPI)

	fs := http.FileServer(http.Dir("."))
	handler.PathPrefix("/web").Handler(fs)

	log.Info("Created a route object")
	return r
}

//Begin is for starting the router module
func (r *App) Begin() error {
	if err := http.ListenAndServe("0.0.0.0:8080", r.handler); err != nil {
		return err
	}
	return nil
}
