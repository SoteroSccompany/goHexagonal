package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Soter-Tec/go-hexagonal/adapters/web/handler"
	"github.com/Soter-Tec/go-hexagonal/aplication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service aplication.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandler(r, n, w.Service)
	http.Handle("/", r)

	server := http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":9000",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

}
