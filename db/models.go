package db

import (
	"errors"

	"github.com/HRMS/models"
)

type DbUser struct {
	Uid        int
	Fname      string
	Lname      string
	Empstatus  string
	Joining    []uint8
	Deleted_at []uint8
}

var ErrNoUser = errors.New("User does not exist")
var ErrNoQuery = errors.New("Unknown function")

type queryFunc func(*models.User) error
type apidata func(val string) ([]byte, error)
type Queries string

const (
	PERSONAL         Queries = "personal"
	BANKEDU                  = "bankedu"
	LEAVES                   = "leaves"
	PAYROLL                  = "payroll"
	HRADMIN                  = "hradmin"
	ITADMIN                  = "itadmin"
	API_DESIGNATIONS         = "designations"
)

type Storage interface {
	Auth(uname, pwd string) (models.User, error)
	BuildUser(uid int, fName Queries) (models.User, error)
	ApiFuncs(val string, fName Queries) ([]byte, error)
	HRAdmin() (models.HRAdmin, error)
}
