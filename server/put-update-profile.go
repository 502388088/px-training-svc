package server


import (
	"net/http"
	"github.com/502388088/px-training-svc/storage"
	"strconv"
)

func HandleUpdateProfile(w http.ResponseWriter, r *http.Request) {
	eid,_ := strconv.Atoi(r.URL.Query().Get("eid"))
	fullname := r.URL.Query().Get("fullname")
	storage.UpdateProfile(fullname, eid)
}