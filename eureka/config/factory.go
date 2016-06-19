package config

import (
	"eureka"
	"fmt"
	"strconv"
	"tietang/utils"
	"time"
)

const (
	DataCenterNameNetflix = "Netflix"
	DataCenterNameAmazon  = "Amazon"
	DataCenterNameMyOwn   = "MyOwn"
	//
	DataCenterNameMyOwnClass = "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo"

	//
	StatusUp           = "UP"             //Ready to receive traffic
	StatusDown         = "DOWN"           // Do not send traffic- healthcheck callback failed
	StatusStarting     = "STARTING"       //Just about starting- initializations to be done - do not send traffic
	StatusOutOfService = "OUT_OF_SERVICE" // Intentionally shutdown for traffic
	StatusUnknown      = "UNKNOWN"
	//
	DEFAULT_LEASE_RENEWAL_INTERVAL = 30
	DEFAULT_LEASE_DURATION         = 90
)

func CreateDataCenterInfo(dataCenterInfo *eureka.DataCenterInfo) *eureka.DataCenterInfo {
	if dataCenterInfo == nil {
		dataCenterInfo = &eureka.DataCenterInfo{}
	}
	if dataCenterInfo.Name == "" {
		dataCenterInfo.Name = DataCenterNameMyOwn
	}

	if dataCenterInfo.Class == "" {
		dataCenterInfo.Class = DataCenterNameMyOwnClass
	}
	return dataCenterInfo
}

func CreateInstanceInfo(config EurekaInstanceConfig) *eureka.InstanceInfo {
	dataCenterInfo := CreateDataCenterInfo(&config.DataCenterInfo)

	leaseInfo := eureka.NewLeaseInfo(config.LeaseRenewalIntervalInSeconds)
	leaseInfo.DurationInSecs = config.LeaseExpirationDurationInSeconds

	ip, _ := utils.GetExternalIP()
	hostName := ip
	appName := config.Appname
	instanceInfo := &eureka.InstanceInfo{
		HostName:       hostName,
		App:            config.Appname,
		AppName:        config.Appname,
		AppGroupName:   config.AppGroupName,
		IpAddr:         ip,
		Status:         StatusStarting,
		DataCenterInfo: dataCenterInfo,
		LeaseInfo:      leaseInfo,
		Metadata:       nil,
	}
	port := config.NonSecurePort
	scheme := "http"
	isSsl := config.SecurePortEnabled
	if isSsl {
		port = config.SecurePort
	}
	stringPort := ":" + strconv.Itoa(port)

	portj := &eureka.Port{
		Port:    port,
		Enabled: true,
	}
	if isSsl {
		scheme = "https"
		instanceInfo.SecureVipAddress = appName //protocol + "://" + hostName + stringPort
		instanceInfo.SecurePort = portj
	} else {
		instanceInfo.VipAddress = appName // protocol + "://" + hostName + stringPort
		instanceInfo.Port = portj
	}
	instanceInfo.StatusPageUrl = scheme + "://" + hostName + stringPort + "/info"
	instanceInfo.HealthCheckUrl = scheme + "://" + hostName + stringPort + "/health"
	instanceInfo.HomePageUrl = scheme + "://" + hostName + stringPort + "/"

	instanceInfo.Metadata = &eureka.MetaData{
		Map: make(map[string]string),
	}
	kv := config.MetadataMap
	for k, v := range kv {
		instanceInfo.Metadata.Map[k] = v
	}
	//	instanceInfo.Metadata.Map["foo"] = "bar" //add metadata for example
	instanceId := fmt.Sprintf("%s:%s:%d", instanceInfo.HostName, appName, instanceInfo.Port.Port)
	instanceInfo.Metadata.Map["instanceId"] = instanceId
	instanceInfo.InstanceId = instanceId

	return instanceInfo
}

func CreateEurekaClient(config EurekaClientConfig) *eureka.Client {
	zones := config.getAvailabilityZones(config.Region)
	machines := make([]string, 1)
	for _, zone := range zones {
		machinesForZone := config.getEurekaServerServiceUrls(zone)
		machines = append(machines, machinesForZone...)
	}
	c := eureka.Config{
		// default timeout is one second
		DialTimeout: time.Second,
	}
	client := eureka.NewClientByConfig(machines, c)
	return client
}
