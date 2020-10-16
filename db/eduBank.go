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
	m.Education = edu
	m.Bank = bank
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
