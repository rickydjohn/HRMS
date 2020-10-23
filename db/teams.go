package db

import (
	"database/sql"

	"github.com/HRMS/models"
)

func (d *db) ListTeams() ([]models.Team, error) {
	var t []models.Team
	rows, err := d.txHandler().Queryx("select tid, name, manager from Team")
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
		t = append(t, tmp)
	}
	return t, nil
}
