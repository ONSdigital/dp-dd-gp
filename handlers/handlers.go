package handlers

import (
	"net/http"
	"github.com/ONSdigital/go-ns/log"
	"html/template"
	"github.com/ONSdigital/dp-dd-gp/model"
)

var PatientList []model.Patient

func Status(w http.ResponseWriter, r *http.Request) {
	var statuses []*model.Status
	for _, patient := range PatientList {
		statuses = append(statuses, &model.Status{Patient: patient, Status: checkup(patient)})
	}

	t, err := template.ParseFiles("templates/all-status.html")
	if err != nil {
		log.Error(err, nil)
	}
	t.Execute(w, statuses)
}

func checkup(p model.Patient) int {
	resp, err := http.Get(p.URL)
	if err != nil {
		log.Error(err, nil)
		return 500
	}
	return resp.StatusCode
}

// Create a handler for the specific patient.
func PatientHandler(p model.Patient) http.HandlerFunc {
	log.Debug("patientHandler " + p.Name, nil)
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		t, err := template.ParseFiles("templates/status.html")
		if err != nil {
			log.Error(err, nil)
		}
		t.Execute(w, &model.Status{Patient: p, Status: checkup(p)})
	})
}
