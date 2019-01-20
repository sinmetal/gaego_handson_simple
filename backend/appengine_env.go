package backend

import (
	"encoding/json"
	"net/http"
	"os"

	"google.golang.org/appengine/log"
)

func AppEngineEnvHandler(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Service  string `json:"service"`
		Version  string `json:"version"`
		Instance string `json:"instance"`
	}{
		Service:  os.Getenv("GAE_SERVICE"),
		Version:  os.Getenv("GAE_VERSION"),
		Instance: os.Getenv("GAE_INSTANCE"),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Errorf(r.Context(), "failed json encode. %+v\n", err)
	}
}
