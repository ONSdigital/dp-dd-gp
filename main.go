package main

import (
	"github.com/gorilla/pat"
	"net/http"
	"os"
	"github.com/ONSdigital/dp-dd-gp/config"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/dp-dd-gp/handlers"
)

func main() {
	config.Load()

	router := pat.New()
	for _, p := range config.PatientList {
		router.HandleFunc(p.Path, handlers.PatientHandler(p))
	}
	handlers.PatientList = config.PatientList
	router.HandleFunc("/", handlers.Status)

	server := http.Server{
		Addr: config.BIND_ADDR,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}
