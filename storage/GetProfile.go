package storage

import (
	"github.com/502388088/px-training-svc/model"
	"github.com/502388088/px-training-svc/postgres"
	"fmt"
)

func GetProfile(eid int) (returnProfile []model.Profile){
	err := postgres.Database.Select(&returnProfile, "SELECT * FROM people.profile WHERE eid = $1", eid)
	if err != nil {
		fmt.Println(err.Error())
	}
	return

}
