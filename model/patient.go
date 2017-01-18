package model

type Patient struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	Path           string `json:"path"`
	HealthCheckURL string `json:"healthCheckURL"`
}

type Status struct {
	Patient Patient
	Status  int
}