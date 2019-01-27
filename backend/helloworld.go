package backend

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/appengine/datastore"
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

	key, err := putToDatastore(r.Context(), res.Message)
	if err != nil {
		log.Errorf(r.Context(), "failed put to datastore. err=%+v", err)
		http.Error(w, "failed put to datastore", http.StatusInternalServerError)
		return
	}
	log.Infof(r.Context(), "Sample Key : %+v", key)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Errorf(r.Context(), "failed json encode. %+v\n", err)
	}
}

type Sample struct {
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

func putToDatastore(ctx context.Context, message string) (*datastore.Key, error) {
	now := time.Now()
	key := datastore.NewKey(ctx, "Sample", now.String(), 0, nil)
	e := Sample{
		Message:   message,
		CreatedAt: now,
	}
	return datastore.Put(ctx, key, &e)
}
