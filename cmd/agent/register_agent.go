package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Jorrit05/DYNAMOS/pkg/etcd"
	"github.com/Jorrit05/DYNAMOS/pkg/lib"
)

func registerAgent() {
	// Prepare agent configuration data
	// var service lib.MicroServiceData = lib.UnmarshalStackFile("/var/log/stack-files/agents.yaml")
	now := time.Now()
	agentConfig := lib.AgentDetails{
		Name:          serviceName,
		ActiveSince:   &now,
		ConfigUpdated: &now,
		RoutingKey:    serviceName + "-in",
	}

	// Serialize agent configuration data as JSON
	configData, err := json.Marshal(agentConfig)
	if err != nil {
		log.Fatal(err)
	}

	go etcd.PutEtcdWithLease(etcdClient, fmt.Sprintf("/agents/%s", agentConfig.Name), string(configData))
}

func updateAgent(agentConfig *lib.AgentDetails) {

	// Update the ActiveSince field
	now := time.Now()
	agentConfig.ConfigUpdated = &now

	// Serialize agent configuration data as JSON
	configData, err := json.Marshal(agentConfig)
	if err != nil {
		log.Fatal(err)
	}

	go etcd.PutEtcdWithLease(etcdClient, fmt.Sprintf("/agents/%s", agentConfig.Name), string(configData))

}
