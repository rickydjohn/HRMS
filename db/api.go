package db

import (
	"encoding/json"

	"github.com/HRMS/models"
)

func (d *db) designations(val string) ([]byte, error) {
	var s models.Salary
	err := d.txHandler().QueryRowx("select b.basic, b.hra, b.lta, b.spa, b.others from Designations a join SalarySlab b on a.desgnID=b.desgnID where b.deleted_at is null and a.desgnID=?", val).StructScan(&s)
	if err != nil {
		return nil, err
	}
	bt, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bt, nil
}
