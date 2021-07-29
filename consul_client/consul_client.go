package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
)

func main() {
	var lastIndex uint64
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("api new client is failed,err:", err)
		return
	}
	services, metainfo, err := client.Health().Service("testName", "v1000", true, &api.QueryOptions{
		WaitIndex: lastIndex,
	})
	if err != nil {
		logrus.Warn("error retrieving instances form Consul:%v", err)
	}
	lastIndex = metainfo.LastIndex
	addrs := map[string]struct{}{}
	for _, service := range services {

		// service.Checks.AggregatedStatus()
		fmt.Println("service.Service.Address:", service.Service.Address, "service.Service.Port:", service.Service.Port)
		addrs[net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))] = struct{}{}
	}

}
