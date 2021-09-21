package main

import (
	"encoding/json"
	"github.com/annakallo/travel-log-server/config"
	"github.com/annakallo/travel-log-server/data/countries"
	"github.com/annakallo/travel-log-server/server"
	"github.com/annakallo/travel-log-server/settings"
	"github.com/annakallo/travelog/log"
	"net/http"
)

const (
	LogPrefix = "travel-log"
)

func initializeConfigAndLogger() {
	conf := config.GetInstance()
	logger := log.GetInstance()
	logger.SetLogFile(conf.LogFile)
	logger.SetLevel(conf.LogLevel)
}

func UpdateTablesVersion() {
	settings.UpdateSettingsTable()
	countries.UpdateCountriesTable()
}

func main() {
	initializeConfigAndLogger()
	log.GetInstance().Error(LogPrefix, "Started application service")

	UpdateTablesVersion()

	r := server.NewRouter()
	e := http.ListenAndServe(":12345", r)
	// @TODO needs to be tested this error handling
	responseJson, _ := json.Marshal(e)
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "Error in listen and serve %s", string(responseJson))
	}
}
