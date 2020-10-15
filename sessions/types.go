package sessions

import (
	"sync"
)

type session struct {
	Ip    string
	Uname string
	Uid   int
}

type SessInterface interface {
	Create(uuid, uname, ip string, uid int) bool
	IsValid(uuid, ip string) error
	Delete(uuid string) error
	GetUID(uuid string) (int, error)
	AllSessions() ([]byte, error)
}

type inSession struct {
	session map[string]session
	lock    sync.Mutex
}
