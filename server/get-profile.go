package server

import (
	"github.com/502388088/px-training-svc/storage"
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
)

func HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	eid,_ := strconv.Atoi(r.URL.Query().Get("eid"))

	profile := storage.GetProfile(eid)
	fmt.Println(fmt.Sprintf("num profile: %d", len(profile)))
	outJSON, _ := json.Marshal(profile)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(outJSON))
}
