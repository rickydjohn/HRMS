package models

type Storage interface {
	Auth(uname, pwd string) (User, error)
	BuildUser(uid int) (User, error)
	Edu(uid int) ([]Education, error)
	Bank(uid int) (Bank, error)
	UsedLeaves(uid int, fyear string) (Leaves, error)
}
