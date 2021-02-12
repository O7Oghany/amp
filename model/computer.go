package model

import "time"

type Computer struct {
	FetchAll ListComputers
}

type ListComputers struct {
	Version  string `json:"version"`
	Metadata `json:"metadata"`
	Data []ListComputersData `json:"data"`
}

type ListComputersData struct {
	ConnectorGUID      string `json:"connector_guid"`
	Hostname           string `json:"hostname"`
	WindowsProcessorID string `json:"windows_processor_id"`
	Active             bool   `json:"active"`
	Links   ComputerLinks      `json:"links"`
	ConnectorVersion string    `json:"connector_version"`
	OperatingSystem  string    `json:"operating_system"`
	InternalIps      []string  `json:"internal_ips"`
	ExternalIP       string    `json:"external_ip"`
	GroupGUID        string    `json:"group_guid"`
	InstallDate      time.Time `json:"install_date"`
	NetworkAddresses []ComputerNetworkAddresses `json:"network_addresses"`
	Policy ComputerPolicy`json:"policy"`
	LastSeen  time.Time     `json:"last_seen"`
	Faults    []interface{} `json:"faults"`
	Isolation ComputerIsolation`json:"isolation"`
	Orbital struct {
		Status string `json:"status"`
	} `json:"orbital"`
}


type ComputerLinks  struct {
	Computer   string `json:"computer"`
	Trajectory string `json:"trajectory"`
	Group      string `json:"group"`
}
type ComputerNetworkAddresses struct {
	Mac string `json:"mac"`
	IP  string `json:"ip"`
}
type ComputerPolicy struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
}

type ComputerIsolation struct {
	Available bool   `json:"available"`
	Status    string `json:"status"`
}