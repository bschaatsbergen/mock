package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bschaatsbergen/mock/model"
)

func StartServer(conf model.Config, port string) {
	r := mux.NewRouter()

	for _, endpoint := range conf.Endpoints {
		route := endpoint
		r.HandleFunc(route.Resource, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(route.Statuscode)
			io.WriteString(w, route.Response)
		}).Methods(route.Method)
	}

	address := ":" + port

	fmt.Printf("📡 starting mock on %s, exit to stop\n", address)

	log.Fatal(http.ListenAndServe(address, r))
}
