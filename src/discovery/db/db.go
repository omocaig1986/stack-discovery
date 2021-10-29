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

package db

import (
	"discovery/log"
	"discovery/types"
	"net"
	"time"
)

const DatabasePath = "db"
const DatabaseName = "discovery.db"

func init() {
	log.Log.Debugf("Starting DB module")

}

func Start() {

}

func AddInitServers(initServersArr []string) {
	initServersValid := 0

	for _, s := range initServersArr {
		// parse the IP
		ip := net.ParseIP(s)
		if ip == nil {
			continue
		}

		err := MachineAdd(&types.Machine{
			IP:         s,
			Alive:      true,
			DeadPolls:  0,
			LastUpdate: time.Now().Unix(),
		}, true)

		if err != nil {
			log.Log.Errorf("Could not add %s as init server: %s", s, err.Error())
		} else {
			initServersValid++
			log.Log.Debugf("Added " + s + " as init server")
		}
	}
	log.Log.Infof("Init DB with %d init servers", initServersValid)
}
