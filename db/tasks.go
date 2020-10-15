package db

import (
	"database/sql"

	"github.com/HRMS/models"
)

func (d *db) BuildUser(uid int) (models.User, error) {
	var u models.User
	basic, err := d.basic(uid)
	if err != nil {
		return u, err
	}
	address, err := d.address(uid)
	if err != nil {
		return u, err
	}
	contact, err := d.contact(uid)
	if err != nil {
		return u, err
	}
	u.Contact = contact
	u.Fname = basic.Fname
	u.Lname = basic.Lname
	u.UID = basic.Uid
	u.EmpStatus = basic.Empstatus
	u.Joining = string(basic.Joining)
	u.Address.District = address.District
	u.Address.House = address.House
	u.Address.Landmark = address.Landmark
	u.Address.State = address.State
	u.Address.Street = address.Street
	u.Address.Zipcode = address.Zipcode
	return u, nil
}

func (d *db) basic(uid int) (basic, error) {
	var b basic
	err := d.txHandler().QueryRowx("select uid, fname, lname, empstatus, convert(date(joining), char) as joining from User where uid=?", uid).StructScan(&b)
	if err != nil && err != sql.ErrNoRows {
		return b, err
	}
	return b, nil
}

func (d *db) address(uid int) (address, error) {
	var b address
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

func (d *db) UsedLeaves(uid int, fyear string) (models.Leaves, error) {
	var b models.Leaves
	rows, err := d.txHandler().Queryx("select type,start, end, status, reason from Leaves where uid=? and start >= date(?) and end < date(date_add(date(?), interval 1 year)) and status in ('approved', 'pending') order by updated_at desc", uid, fyear, fyear)
	if err != nil && err != sql.ErrNoRows {
		return b, err
	}
	for rows.Next() {
		var t models.Leave
		if err := rows.StructScan(&t); err != nil {
			return b, err
		}
		b.LeaveHistory = append(b.LeaveHistory, t)
	}

	err = d.txHandler().QueryRowx("select ifnull(sum(datediff(end, start)),0) as used from Leaves where uid=? and status='approved' and start >= date(?) and end < date(date_add(date(?), interval 1 year))", uid, fyear, fyear).Scan(&b.TotalLeaves)
	if err != nil && err != sql.ErrNoRows {
		return b, err
	}
	return b, nil
}
