package main

import (
	"net/http"

	"github.com/sinmetal/gaego_handson_simple/backend"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/api/helloworld", backend.HelloWorldHandler)
	http.HandleFunc("/admin/appengine-env", backend.AppEngineEnvHandler)

	appengine.Main()
}
