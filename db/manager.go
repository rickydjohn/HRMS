package db

import (
	"time"

	"github.com/HRMS/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type db struct {
	db          *sqlx.DB
	uname       string
	pwd         string
	ip          string
	startOfYear string
}

func Begin(dbconn models.DB, startOfYear string) (Storage, error) {
	var d db
	d.uname = dbconn.Uname
	d.pwd = dbconn.Pwd
	d.ip = dbconn.IP
	d.startOfYear = startOfYear

	dsn := dbconn.Uname + ":" + dbconn.Pwd + "@tcp(" + dbconn.IP + ":3306)/HRMS?parseTime=true"

	conn := sqlx.MustConnect("mysql", dsn)
	conn.SetConnMaxIdleTime(time.Duration(1 * time.Minute))
	conn.SetConnMaxLifetime(time.Duration(1 * time.Minute))
	conn.SetMaxOpenConns(100)
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	d.db = conn
	return &d, nil
}

func (d *db) txHandler() *sqlx.Tx {
	err := d.db.Ping()
	if err != nil {
		conn := sqlx.MustConnect("mysql", d.uname+":"+d.pwd+"tcp@("+d.ip+":3306)/HRMS")
		if err = conn.Ping(); err != nil {
			panic("Unable to connect to DB: " + err.Error())
		}
		d.db = conn
	}

	tx, err := d.db.Beginx()
	if err != nil {
		panic("Unable to create a transaction: " + err.Error())
	}
	return tx
}

func (d *db) Auth(uname, pwd string) (models.User, error) {
	var m models.User
	var u DbUser
	tx := d.txHandler()
	err := tx.QueryRowx("select a.uid, b.fname, b.lname, b.empstatus, b.joining, b.deleted_at from Auth a join User b on a.uid=b.uid where a.uname=? and a.pwd=password(?) and b.deleted_at is null", uname, pwd).StructScan(&u)
	if err != nil {
		return m, err
	}
	m.Fname = u.Fname
	m.Lname = u.Lname
	m.EmpStatus = u.Empstatus
	m.UID = u.Uid
	return m, nil
}
