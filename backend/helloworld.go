package backend

import (
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/appengine/log"
)

// HelloWorldHandler is /api/helloworld request を処理するハンドラ
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Message string    `json:"message"`
		Time    time.Time `json:"time"`
	}{
		Message: "Hello Google App Engine Standard Go",
		Time:    time.Now(),
	}
	log.Infof(r.Context(), "%+v\n", res)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Errorf(r.Context(), "failed json encode. %+v\n", err)
	}
}
