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
	"discovery/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ConfigurationSet struct {
	machineIp                         string
	machineId                         string
	machineFogNetId                   string
	initServers                       []string
	pollTime                          uint
	listeningPort                     uint
	pollTimeout                       uint
	machineDeadPollsRemovingThreshold uint
	runningEnvironment                string
	defaultIface                      string
}

type ConfigurationSetExp struct {
	MachineIp                         string   `json:"machine_ip" bson:"machine_ip"`
	MachineId                         string   `json:"machine_id" bson:"machine_id"`
	MachineFogNetId                   string   `json:"machine_fog_net_id" bson:"machine_fog_net_id"`
	InitServers                       []string `json:"init_servers" bson:"init_servers"`
	PollTime                          uint     `json:"poll_time" bson:"poll_time"`
	ListeningPort                     uint     `json:"listening_port" bson:"listening_port"`
	PollTimeout                       uint     `json:"poll_timeout" bson:"poll_timeout"`
	MachineDeadPollsRemovingThreshold uint     `json:"machine_dead_polls_removing_threshold" bson:"machine_dead_polls_removing_threshold"`
	RunningEnvironment                string   `json:"running_environment" bson:"running_environment"`
	DefaultIface                      string   `json:"default_iface" bson:"default_iface"`
}

/*
 * Sample configuration file
 *
 * {
 *   "machine_ip": "192.168.99.102",
 *   "machine_id": "p2pfogc2n0",
 *   "init_servers": ["192.168.99.100"]
 * }
 *
 */

type ConfigError struct{}

func (ConfigError) Error() string {
	return "Configuration Error"
}

/*
 * Getters
 */

func (c ConfigurationSet) GetMachineIp() string {
	return c.machineIp
}
func (c ConfigurationSet) GetMachineId() string {
	return c.machineId
}
func (c ConfigurationSet) GetMachineFogNetId() string {
	return c.machineFogNetId
}
func (c ConfigurationSet) GetInitServers() []string {
	return c.initServers
}
func (c ConfigurationSet) GetPollTime() uint {
	return c.pollTime
}
func (c ConfigurationSet) GetListeningPort() uint {
	return c.listeningPort
}
func (c ConfigurationSet) GetPollTimeout() uint {
	return c.pollTimeout
}
func (c ConfigurationSet) GetMachineDeadPollsRemovingThreshold() uint {
	return c.machineDeadPollsRemovingThreshold
}
func (c ConfigurationSet) GetRunningEnvironment() string {
	return c.runningEnvironment
}
func (c ConfigurationSet) GetDefaultIface() string {
	return c.defaultIface
}

// GetConfiguration returns the configuration with exported fields
func (c ConfigurationSet) GetConfiguration() *ConfigurationSetExp {
	conf := &ConfigurationSetExp{}
	copyAllFieldsToExp(Configuration, conf)
	return conf
}

/*
 * Setters
 */

func (c *ConfigurationSet) SetMachineIp(ip string) {
	c.machineIp = ip
}
func (c *ConfigurationSet) SetMachineId(id string) {
	c.machineId = id
}
func (c *ConfigurationSet) SetMachineFogNetId(id string) {
	c.machineFogNetId = id
}
func (c *ConfigurationSet) SetInitServers(servers []string) {
	c.initServers = servers
}
func (c *ConfigurationSet) SetPollTime(time uint) {
	c.pollTime = time
}
func (c *ConfigurationSet) SetListeningPort(port uint) {
	c.listeningPort = port
}
func (c *ConfigurationSet) SetPollTimeout(port uint) {
	c.pollTimeout = port
}
func (c *ConfigurationSet) SetMachineDeadPollsRemovingThreshold(thr uint) {
	c.machineDeadPollsRemovingThreshold = thr
}
func (c *ConfigurationSet) SetDefaultIface(s string) {
	c.defaultIface = s
}

// SetConfiguration updates the entire configuration
func (c *ConfigurationSet) SetConfiguration(exp *ConfigurationSetExp) {
	copyAllFieldsToUnExp(exp, c)
}

/*
 * Utils
 */

func ReadConfigurationFromEnv() {

}

func ReadConfigFile() (*ConfigurationSet, bool) {
	conf := GetDefaultExpConfiguration()
	confValid := ConfigurationSet{}
	noConfigurationFile := false

	file, err := ioutil.ReadFile(GetConfigFilePath())
	if err != nil {
		log.Log.Info("Cannot read configuration file at %s", GetConfigFilePath())
		noConfigurationFile = true
	} else {
		err = json.Unmarshal(file, &conf)
		if err != nil {
			log.Log.Errorf("Cannot decode configuration file, maybe not valid json: %s", err.Error())
			noConfigurationFile = true
		}
	}

	// update fields
	if conf.RunningEnvironment == "" ||
		(conf.RunningEnvironment != RunningEnvironmentDevelopment && conf.RunningEnvironment != RunningEnvironmentProduction) {
		conf.RunningEnvironment = RunningEnvironmentDevelopment
	}
	copyAllFieldsToUnExp(conf, &confValid)

	// update config field
	Configuration = &confValid

	// check fields
	if conf.MachineId == "" || conf.MachineIp == "" {
		log.Log.Warningf("Configuration file does not contain MachineId or MachineIp. Will try to get ip from \"%s\"", Configuration.GetDefaultIface())
		// get ip from machine
		ip, err := utils.GetInternalIP(Configuration.GetDefaultIface())
		if err != nil {
			return &confValid, noConfigurationFile
		}
		confValid.machineIp = ip
		// generate machine id
		confValid.machineId = fmt.Sprintf("p2pfaas-%s", confValid.machineIp)
		log.Log.Infof("Got from machine ip: %s and id: %s", confValid.machineIp, confValid.machineId)
	}

	return &confValid, noConfigurationFile
}

func GetDefaultExpConfiguration() *ConfigurationSetExp {
	conf := &ConfigurationSetExp{
		MachineIp:                         "",
		MachineId:                         "",
		MachineFogNetId:                   "",
		InitServers:                       []string{},
		PollTime:                          DefaultPollTime,
		ListeningPort:                     DefaultListeningPort,
		PollTimeout:                       DefaultPollTimeoutTime,
		MachineDeadPollsRemovingThreshold: DefaultMachineDeadPollsRemovingThreshold,
		DefaultIface:                      DefaultIfaceName,
		RunningEnvironment:                os.Getenv(EnvRunningEnvironment),
	}
	return conf
}

func copyAllFieldsToExp(from *ConfigurationSet, to *ConfigurationSetExp) {
	to.MachineIp = from.machineIp
	to.MachineId = from.machineId
	to.MachineFogNetId = from.machineFogNetId
	to.InitServers = from.initServers
	to.PollTime = from.pollTime
	to.ListeningPort = from.listeningPort
	to.PollTimeout = from.pollTimeout
	to.MachineDeadPollsRemovingThreshold = from.machineDeadPollsRemovingThreshold
	to.DefaultIface = from.defaultIface
	to.RunningEnvironment = from.runningEnvironment
}

func copyAllFieldsToUnExp(from *ConfigurationSetExp, to *ConfigurationSet) {
	to.machineIp = from.MachineIp
	to.machineId = from.MachineId
	to.machineFogNetId = from.MachineFogNetId
	to.initServers = from.InitServers
	to.pollTime = from.PollTime
	to.listeningPort = from.ListeningPort
	to.pollTimeout = from.PollTimeout
	to.machineDeadPollsRemovingThreshold = from.MachineDeadPollsRemovingThreshold
	to.defaultIface = from.DefaultIface
	to.runningEnvironment = from.RunningEnvironment
}
