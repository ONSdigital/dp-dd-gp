package main

import (
	"github.com/gorilla/pat"
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"github.com/ONSdigital/dp-dd-gp/patients"
	"github.com/ONSdigital/go-ns/log"
)

func main() {
	fmt.Println("Starting stuff!")
	file, err := os.Open("config.json")
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}

	var pList []patients.Patient
	err = json.NewDecoder(file).Decode(&pList)

	if err != nil {
		log.Error(err, nil)
		fmt.Println("Failed to load config.")
	}

	fmt.Printf("List %+v", pList)

	router := pat.New()

	for _, p := range pList {
		router.HandleFunc(p.Path, patientHandler(p))
	}

	server := http.Server{
		Addr: ":22000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}


// Create a handler for the specific patient.
func patientHandler(p patients.Patient) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		resp, err := http.Get(p.Addr)
		if err != nil || resp.StatusCode != 200 {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("%s is down", p.Name)))
		} else {
			w.WriteHeader(200)
			w.Write([]byte(fmt.Sprintf("%s is OK", p.Name)))
		}
	})
}
