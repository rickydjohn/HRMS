package db

import (
	"database/sql"

	"github.com/HRMS/models"
)

func (d *db) HRAdmin() (models.HRAdmin, error) {
	var t models.HRAdmin
	rows, err := d.db.Queryx("select tid, name, manager from Team")
	if err == sql.ErrNoRows {
		return t, nil
	} else if err != nil {
		return t, err
	}
	for rows.Next() {
		var tmp models.Team
		if err := rows.StructScan(&tmp); err != nil {
			return t, err
		}
		t.Teams = append(t.Teams, tmp)
	}

	drows, err := d.db.Queryx("select desgnID, name, grade, level from Designations")
	if err == sql.ErrNoRows {
		return t, nil
	} else if err != nil {
		return t, err
	}

	for drows.Next() {
		var tmp models.Designation
		if err := drows.StructScan(&tmp); err != nil {
			return t, err
		}
		t.Designations = append(t.Designations, tmp)
	}
	return t, nil
}
