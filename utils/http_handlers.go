package utils

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bschaatsbergen/mock/model"
)

func StartServer(conf model.Config, port string) {
	r := mux.NewRouter()

	for _, endpoint := range conf.Endpoints {
		r.HandleFunc(endpoint.Resource, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(endpoint.Response))
		}).Methods(endpoint.Method)
	}

	address := ":" + port

	http.ListenAndServe(address, r)
}
