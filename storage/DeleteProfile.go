package storage

import (
	"github.com/502388088/px-training-svc/postgres"
)

func DeleteProfile(eid int) {
	tx, _ := postgres.Database.Begin()
	stmtDataRows, _ := tx.Prepare("DELETE FROM people.profile WHERE eid = $1")
	defer stmtDataRows.Close()
	stmtDataRows.Exec(eid)
	tx.Commit()
	return
}