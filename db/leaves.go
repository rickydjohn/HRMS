package db

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/HRMS/models"
)

func (d *db) leaves(m *models.User) error {
	l, err := d.leaveData(m.UID)
	if err != nil {
		return err
	}
	m.Leaves = l
	return nil
}

func (d *db) leaveData(uid int) (models.Leaves, error) {
	fyear := strconv.Itoa(time.Now().Year()) + "-" + d.startOfYear + "-01"
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
