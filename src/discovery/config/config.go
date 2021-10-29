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

package config

import (
	"discovery/log"
)

// application
const Version = "0.0.3b"
const UserAgentMachine = "Machine"

const DataPath = "/data"
const ConfigurationFileName = "p2p_faas-discovery.json"

const GetParamIp = "p2pfaas-machine-ip"
const GetParamName = "p2pfaas-machine-name"
const GetParamGropuName = "p2pfaas-machine-group-name"

// default parameters
const DefaultListeningPort = 19000
const DefaultPollTime = 120      // seconds
const DefaultPollTimeoutTime = 5 // seconds
const DefaultIfaceName = "eth0"

// MachineDeadPollsRemovingThreshold tells the number of times we need to poll the machine for removing it from the db
const DefaultMachineDeadPollsRemovingThreshold = 20

// env
const EnvRunningEnvironment = "P2PFAAS_DEV_ENV"
const RunningEnvironmentProduction = "production"
const RunningEnvironmentDevelopment = "development"

// Configuration general parsed configuration
var Configuration *ConfigurationSet

// ConfigurationReadFromFile Set if configuration has been read from file or not
var ConfigurationReadFromFile = false

func init() {
	// init configuration
	_, noConfigurationFile := ReadConfigFile()
	if noConfigurationFile {
		log.Log.Info("Configuration file not present, creating...")
		// save to fs
		SaveConfigurationToConfigFile()
	} else {
		log.Log.Info("Loaded configuration file")
		ConfigurationReadFromFile = true
	}

	log.Log.Info("Starting in %s environment", Configuration.GetRunningEnvironment())
}

func Start() {

}
