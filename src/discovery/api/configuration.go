/*
 * P2PFaaS - A framework for FaaS Load Balancing
 * Copyright (c) 2019. Gabriele Proietti Mattia <pm.gabriele@outlook.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package api

import (
	config2 "discovery/config"
	"discovery/db"
	"discovery/errors"
	"discovery/log"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func GetConfiguration(w http.ResponseWriter, r *http.Request) {
	// if we do not have the machine ip we report 404
	if config2.Configuration.GetMachineIp() == "" {
		errors.ReplyWithError(w, errors.ConfigurationNotReady)
		return
	}

	config, err := json.Marshal(config2.Configuration.GetConfiguration())
	if err != nil {
		log.Log.Errorf("Cannot encode configuration to json")
		errors.ReplyWithError(w, errors.GenericError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(config))
}

func SetConfiguration(w http.ResponseWriter, r *http.Request) {
	defaultConfiguration := config2.GetDefaultExpConfiguration()
	currentConfiguration := config2.Configuration.GetConfiguration()
	reqBody, _ := ioutil.ReadAll(r.Body)

	var newConfiguration *config2.ConfigurationSetExp
	var err error
	// do the merge with the default configuration or existing
	if config2.ConfigurationReadFromFile {
		err = json.Unmarshal(reqBody, &currentConfiguration)
		newConfiguration = currentConfiguration
	} else {
		err = json.Unmarshal(reqBody, &defaultConfiguration)
		newConfiguration = defaultConfiguration
	}
	if err != nil {
		log.Log.Errorf("Cannot encode passed configuration")
		errors.ReplyWithError(w, errors.GenericError)
		return
	}

	// update existing configuration
	config2.Configuration.SetConfiguration(newConfiguration)
	// clean machines and update init servers
	err = db.MachineRemoveAll()
	if err != nil {
		log.Log.Errorf("Error while clearing nodes list: %s", err.Error())
	} else {
		log.Log.Infof("Nodes list cleared")
	}
	db.AddInitServers(config2.Configuration.GetInitServers())

	// save configuration to file
	configJson, err := json.Marshal(&newConfiguration)
	err = config2.SaveConfigurationToConfigFile()
	if err != nil {
		log.Log.Error("Cannot save configuration file to disk: %s", err.Error())
	} else {
		log.Log.Infof("Configuration updated with %s", configJson)
	}

	w.WriteHeader(200)
}
