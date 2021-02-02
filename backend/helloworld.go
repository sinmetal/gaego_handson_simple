package backend

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Handlers struct {
	AccessLogStore *AccessLogStore
}

// HelloWorldHandler is /api/helloworld request を処理するハンドラ
func (h *Handlers) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res := struct {
		Message string    `json:"message"`
		Time    time.Time `json:"time"`
	}{
		Message: "Hello Google App Engine Standard Go",
		Time:    time.Now(),
	}
	log.Printf("%+v\n", res)

	key, err := h.AccessLogStore.Insert(ctx, &AccessLog{
		ID: uuid.New().String(),
	})
	if err != nil {
		log.Printf("failed put to datastore. err=%+v", err)
		http.Error(w, "failed put to datastore", http.StatusInternalServerError)
		return
	}
	log.Printf("AccessLog Key : %+v", key)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("failed json encode. %+v\n", err)
	}
}
