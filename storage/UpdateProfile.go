package storage

import "github.com/502388088/px-training-svc/postgres"

func UpdateProfile(fullname string, eid int) (err error) {
	tx, err := postgres.Database.Begin()
	defer tx.Rollback()

	stmtDataRows, _ := tx.Prepare("UPDATE poeple.profile SET fullname = $1 WHERE eid = $2")
	defer stmtDataRows.Close()
	stmtDataRows.Exec(fullname, eid)
	tx.Commit()
	return
}

