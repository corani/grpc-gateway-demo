package api

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
)

//go:embed helloworld/hello_world.swagger.json
var spec []byte

func RegisterOpenAPIHandler(r *mux.Router, path string) {
	spec, err := openapi3.NewLoader().LoadFromData(spec)
	if err != nil {
		log.Fatalf("failed to load: %v", err)
	}

	r.NewRoute().Path(path).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(spec); err != nil {
			log.Printf("failed to encode: %v", err)
		}
	})
}
