package server

import (
	"github.com/502388088/px-training-svc/storage"
	"net/http"
	"strconv"
)

func HandleDeleteProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	eid,_ := strconv.Atoi(r.URL.Query().Get("eid"))

	storage.DeleteProfile(eid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}