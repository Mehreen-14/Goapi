package main

import (
	"fmt"
	"net/http"

	"github.com/Mehreen-14/Goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Server starting...")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}