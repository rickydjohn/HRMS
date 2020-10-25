package db

import (
	"database/sql"

	"github.com/HRMS/models"
)

func (d *db) bankedu(m *models.User) error {
	edu, err := d.edu(m.UID)
	if err != nil {
		return err
	}
	bank, err := d.bank(m.UID)
	if err != nil {
		return err
	}
	empl, err := d.emphistory(m.UID)
	if err != nil {
		return err
	}
	m.Education = edu
	m.Bank = bank
	m.EmpHistory = empl
	return nil
}

func (d *db) edu(uid int) ([]models.Education, error) {
	var educ []models.Education
	rows, err := d.txHandler().Queryx("select id, institution, course, yop, mop from Education where deleted_at is NULL and uid=? order by yop desc", uid)
	defer rows.Close()
	if err != nil && err != sql.ErrNoRows {
		return educ, err
	}
	for rows.Next() {
		var t models.Education
		//if err := rows.Scan(&t.ID, &t.Institution, &t.Course, &t.Yop, &t.Mop); err != nil {
		if err := rows.StructScan(&t); err != nil {
			return educ, err
		}
		educ = append(educ, t)
	}
	return educ, nil
}

func (d *db) bank(uid int) (models.Bank, error) {
	var bank models.Bank
	if err := d.txHandler().QueryRowx("select pan, account, ifsc, name from Bank where uid=?", uid).StructScan(&bank); err != nil && err != sql.ErrNoRows {
		return bank, err
	}
	return bank, nil
}

func (d *db) emphistory(uid int) ([]models.EmpHistory, error) {
	var empl []models.EmpHistory
	rows, err := d.txHandler().Queryx("select company, fromMonth, fromYear, toMonth, toYear from EmpHistory where uid=? order by toYear desc", uid)
	defer rows.Close()
	if err != nil && err != sql.ErrNoRows {
		return empl, err
	}
	for rows.Next() {
		var t models.EmpHistory
		//if err := rows.Scan(&t.ID, &t.Institution, &t.Course, &t.Yop, &t.Mop); err != nil {
		if err := rows.StructScan(&t); err != nil {
			return empl, err
		}
		empl = append(empl, t)
	}
	return empl, nil
}
