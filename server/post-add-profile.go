package server

import (
	"net/http"
	"github.com/502388088/px-training-svc/storage"
	"github.com/502388088/px-training-svc/model"
	"strconv"
)

func HandleAddProfile(w http.ResponseWriter, r *http.Request) {
	eid,_ := strconv.Atoi(r.URL.Query().Get("eid"))
	fullname := r.URL.Query().Get("fullname")
	position := r.URL.Query().Get("position")

	profile := model.Profile{
		Eid: eid,
		Fullname: fullname,
		Position: position,
	}
	storage.AddProfile(profile)
}