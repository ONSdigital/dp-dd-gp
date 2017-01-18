package config

import (
	"os"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/dp-dd-gp/model"
	"encoding/json"
	"fmt"
)

const bindAddrKey = "BIND_ADDR"
const configJsonPath = "config.json"

var BIND_ADDR = ":22000"
var PatientList []model.Patient

func init() {

	if bindAddrEnv := os.Getenv(bindAddrKey); len(bindAddrEnv) > 0 {
		BIND_ADDR = bindAddrEnv
	}

	file, err := os.Open(configJsonPath)
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}

	err = json.NewDecoder(file).Decode(&PatientList)

	if err != nil {
		log.Error(err, nil)
		fmt.Println("Failed to load config.")
	}
}

func Load() {
	// Will invoke init on first call.
	log.Debug("Config", log.Data{
		bindAddrKey: BIND_ADDR,
		"patients": PatientList,
	})
}
