package storage

import (
	"github.com/502388088/px-training-svc/model"
	"github.com/502388088/px-training-svc/postgres"
)

func AddProfile(profileObject model.Profile) (err error) {
	tx, err := postgres.Database.Begin()
	defer tx.Rollback()

	stmtDataRows, _ := tx.Prepare("INSERT INTO people.profile (eid,fullname,position) VALUES ($1,$2,$3)")
	defer stmtDataRows.Close()
	stmtDataRows.Exec(profileObject.Eid, profileObject.Fullname, profileObject.Position)
	tx.Commit()
	return
}
