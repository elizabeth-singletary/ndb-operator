/*
Copyright 2022-2023 Nutanix, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ndb_api

type DatabaseNode struct {
	Id               string         `json:"id"`
	Name             string         `json:"name"`
	DatabaseServerId string         `json:"dbServerId"`
	DbServer         DatabaseServer `json:"dbserver"`
}

type DatabaseServer struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	IPAddresses []string `json:"ipAddresses"`
}

type TimeMachineInfo struct {
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	SlaId            string            `json:"slaId"`
	Schedule         map[string]string `json:"schedule"`
	Tags             []string          `json:"tags"`
	AutoTuneLogDrive bool              `json:"autoTuneLogDrive"`
}

type ActionArgument struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Node struct {
	Properties []string `json:"properties"`
	VmName     string   `json:"vmName"`
}

type Property struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
}