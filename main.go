package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/profiler"
	"github.com/sinmetal/gaego_handson_simple/backend"
)

func startProfiler(projectID string) error {
	cfg := profiler.Config{
		ProjectID:      projectID,
		Service:        "gaego_handson_simple",
		ServiceVersion: "0.0.1",

		// For OpenCensus users:
		// To see Profiler agent spans in APM backend,
		// set EnableOCTelemetry to true
		// EnableOCTelemetry: true,
	}

	// Profiler initialization, best done as early as possible.
	return profiler.Start(cfg)
}

func main() {
	ctx := context.Background()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	var err error
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if metadata.OnGCE() {
		projectID, err = metadata.ProjectID()
		if err != nil {
			log.Fatal(err)
		}
		if err := startProfiler(projectID); err != nil {
			log.Fatal(err)
		}
	}

	ds, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	accessLogStore, err := backend.NewAccessLogStore(ctx, ds)
	if err != nil {
		log.Fatal(err)
	}

	h := &backend.Handlers{
		AccessLogStore: accessLogStore,
	}

	http.HandleFunc("/api/helloworld", h.HelloWorldHandler)
	http.HandleFunc("/admin/appengine-env", backend.AppEngineEnvHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), http.DefaultServeMux); err != nil {
		log.Printf("failed ListenAndServe err=%+v", err)
	}
}
