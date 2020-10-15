package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	_ "net/http/pprof"

	"github.com/HRMS/db"
	"github.com/HRMS/models"
	"github.com/HRMS/routes"
	"github.com/HRMS/sessions"
	"github.com/sirupsen/logrus"
)

func config(path string, log *logrus.Entry) *models.SvcConfig {
	log.Info("accessing config file at ", path)
	bt, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error occurred while reading file: ", err.Error)
		return nil
	}
	var conf models.SvcConfig
	if err := json.Unmarshal(bt, &conf); err != nil {
		log.Fatal("Error while parsing config file: ", err)
	}

	return &conf
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true,
		DisableTimestamp: true})
	logf := flag.String("c", "", "location of the config file")
	flag.Parse()
	c := config(*logf, log.WithField("Module", "Config"))
	session := sessions.Begin()
	db, err := db.Begin(c.DB)
	if err != nil {
		panic(err)
	}
	r := routes.New(log.WithField("module", "routes"), c.Application, session, db)
	if err := r.Begin(); err != nil {
		panic(err)
	}
}
