package db

import (
	"database/sql"

	"github.com/HRMS/models"
)

func (d *db) BuildUser(uid int, funcName Queries) (models.User, error) {
	basic, err := d.basic(uid)
	if err != nil {
		return models.User{}, err
	}
	f, err := d.funcmap(funcName)
	if err != nil {
		return models.User{}, err
	}
	if err := f(&basic); err != nil {
		return models.User{}, err
	}
	return basic, nil
}

func (d *db) personal(m *models.User) error {
	address, err := d.address(m.UID)
	if err != nil {
		return err
	}
	contact, err := d.contact(m.UID)
	if err != nil {
		return err
	}
	m.Address = address
	m.Contact = contact
	return nil
}

func (d *db) basic(uid int) (models.User, error) {
	var b models.User
	err := d.txHandler().QueryRowx("select a.uid, a.fname, a.lname, a.empstatus, convert(date(a.joining), char) as joining, ifnull(c.name, '') as role  from User a left join Role b on a.uid=b.uid left join KnownRoles c on b.roleID=c.roleId where a.uid=?", uid).StructScan(&b)
	if err != nil && err != sql.ErrNoRows {
		return b, err
	}
	return b, nil
}

func (d *db) address(uid int) (models.Address, error) {
	var b models.Address
	err := d.txHandler().QueryRowx("select house,street,district,state,zipcode,landmark from Address where uid=?", uid).StructScan(&b)
	if err != nil && err != sql.ErrNoRows {
		return b, err
	}
	return b, nil
}

func (d *db) contact(uid int) (models.Contact, error) {
	var b models.Contact
	err := d.txHandler().QueryRowx("select phone,email,ename,ephone from Contact where uid=?", uid).StructScan(&b)
	if err != nil && err != sql.ErrNoRows {
		return b, err
	}
	return b, nil
}
