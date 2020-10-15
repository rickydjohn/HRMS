package db

import (
	"errors"
)

type DbUser struct {
	Uid        int
	Fname      string
	Lname      string
	Empstatus  string
	Joining    []uint8
	Deleted_at []uint8
}

type basic struct {
	Uid       int     `db:"uid"`
	Fname     string  `db:"fname"`
	Lname     string  `db:"lname"`
	Empstatus string  `db:"empstatus"`
	Joining   []uint8 `db:"joining"`
}

type address struct {
	Uid      int    `db:"uid"`
	House    string `db:"house"`
	Street   string `db:"street"`
	District string `db:"district"`
	State    string `db:"state"`
	Zipcode  int    `db:"zipcode"`
	Landmark string `db:"landmark"`
}

var ErrNoUser = errors.New("User does not exist")
