/*
 * P2PFaaS - A framework for FaaS Load Balancing
 * Copyright (c) 2020. Gabriele Proietti Mattia <pm.gabriele@outlook.com>
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
 *
 */

package config

import (
	"discovery/log"
	"encoding/json"
	"io/ioutil"
)

func GetConfigFilePath() string {
	return GetDataPath() + "/" + ConfigurationFileName
}

func SaveConfigurationToConfigFile() error {
	// prepare configuration
	confExported := GetDefaultExpConfiguration()
	copyAllFieldsToExp(Configuration, confExported)

	// save configuration to file
	configJson, err := json.MarshalIndent(confExported, "", "  ")
	err = ioutil.WriteFile(GetConfigFilePath(), configJson, 0644)
	if err != nil {
		log.Log.Errorf("Cannot save configuration to file %s: %s", GetConfigFilePath(), err.Error())
		return err
	}

	return nil
}
