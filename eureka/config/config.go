package config

import (
	"eureka"
	"strings"
)

const (
	EurekaAcceptFull    = "full"
	EurekaAcceptCompact = "compact"
	DEFAULT_EUREKA_URL  = "http://localhost:8761/eureka/"
	DEFAULT_EUREKA_ZONE = "defaultZone"
)

type Config struct {
	Eureka Eureka
}
type Eureka struct {
	Client   EurekaClientConfig
	Instance EurekaInstanceConfig
}
type EurekaTransportConfig struct {
	SessionedClientReconnectIntervalSeconds           int
	RetryableClientQuarantineRefreshPercentage        float64
	BootstrapResolverRefreshIntervalSeconds           int
	ApplicationsResolverDataStalenessThresholdSeconds int
	AsyncResolverRefreshIntervalMs                    int
	AsyncResolverWarmUpTimeoutMs                      int
	AsyncExecutorThreadPoolSize                       int
	ReadClusterVip                                    string
	BootstrapResolverForQuery                         bool
}

func NewEurekaTransportConfig() EurekaTransportConfig {
	return EurekaTransportConfig{
		SessionedClientReconnectIntervalSeconds:           1200,
		RetryableClientQuarantineRefreshPercentage:        0.66,
		BootstrapResolverRefreshIntervalSeconds:           300,
		ApplicationsResolverDataStalenessThresholdSeconds: 300,
		AsyncResolverRefreshIntervalMs:                    300000,
		AsyncResolverWarmUpTimeoutMs:                      5000,
		AsyncExecutorThreadPoolSize:                       5,
		BootstrapResolverForQuery:                         true,
	}
}

type EurekaInstanceConfig struct {
	Appname                          string
	AppGroupName                     string
	InstanceEnabledOnit              bool
	NonSecurePort                    int
	SecurePort                       int
	NonSecurePortEnabled             bool
	SecurePortEnabled                bool
	LeaseRenewalIntervalInSeconds    uint
	LeaseExpirationDurationInSeconds int
	VirtualHostName                  string
	InstanceId                       string
	SecureVirtualHostName            string
	ASGName                          string
	DataCenterInfo                   eureka.DataCenterInfo
	IpAddress                        string
	StatusPageUrlPath                string
	StatusPageUrl                    string
	HomePageUrlPath                  string
	HomePageUrl                      string
	HealthCheckUrlPath               string
	HealthCheckUrl                   string
	SecureHealthCheckUrl             string
	Namespace                        string
	Hostname                         string
	PreferIpAddress                  bool
	InitialStatus                    string
	MetadataMap                      map[string]string
}

func NewEurekaInstanceConfig() EurekaInstanceConfig {
	ins := EurekaInstanceConfig{
		Appname:                          "unknow",
		NonSecurePort:                    80,
		SecurePort:                       443,
		NonSecurePortEnabled:             true,
		LeaseRenewalIntervalInSeconds:    30,
		LeaseExpirationDurationInSeconds: 90,
		StatusPageUrlPath:                "/info",
		HomePageUrlPath:                  "/",
		HealthCheckUrlPath:               "/health",
		Namespace:                        "eureka",
		PreferIpAddress:                  true,
		InitialStatus:                    eureka.UP,
	}
	ins.VirtualHostName = ins.Appname

	return ins
}

type EurekaClientConfig struct {
	Transport                                     EurekaTransportConfig
	RegistryFetchIntervalSeconds                  int
	InstanceInfoReplicationIntervalSeconds        int
	InitialInstanceInfoReplicationIntervalSeconds int
	EurekaServiceUrlPollIntervalSeconds           int
	EurekaServerReadTimeoutSeconds                int
	EurekaServerConnectTimeoutSeconds             int
	BackupRegistryImpl                            string
	EurekaServerTotalConnections                  int
	EurekaServerTotalConnectionsPerHost           int
	EurekaServerURLContext                        string
	EurekaServerPort                              string
	EurekaServerDNSName                           string
	Region                                        string
	EurekaConnectionIdleTimeoutSeconds            int
	RegistryRefreshSingleVipAddress               string
	HeartbeatExecutorThreadPoolSize               int
	HeartbeatExecutorExponentialBackOffBound      int
	CacheRefreshExecutorThreadPoolSize            int
	CacheRefreshExecutorExponentialBackOffBound   int
	GZipContent                                   bool
	UseDnsForFetchingServiceUrls                  bool
	RegisterWithEureka                            bool
	PreferSameZoneEureka                          bool
	LogDeltaDiff                                  bool
	DisableDelta                                  bool
	FetchRemoteRegionsRegistry                    string
	FilterOnlyUpInstances                         bool
	FetchRegistry                                 bool
	DollarReplacement                             string
	EscapeCharReplacement                         string
	AllowRedirects                                bool
	OnDemandUpdateStatusChange                    bool
	EncoderName                                   string
	DecoderName                                   string
	ClientDataAccept                              string
	AvailabilityZones                             map[string]string
	ServiceUrl                                    map[string]string
}

func NewEurekaClientConfig() EurekaClientConfig {
	return EurekaClientConfig{
		Transport:                                     NewEurekaTransportConfig(),
		RegistryFetchIntervalSeconds:                  30,
		InstanceInfoReplicationIntervalSeconds:        30,
		InitialInstanceInfoReplicationIntervalSeconds: 40,
		EurekaServiceUrlPollIntervalSeconds:           300,
		EurekaServerReadTimeoutSeconds:                8,
		EurekaServerConnectTimeoutSeconds:             5,
		EurekaServerTotalConnections:                  200,
		EurekaServerTotalConnectionsPerHost:           50,
		Region: "cn-east-1",
		EurekaConnectionIdleTimeoutSeconds:          30,
		HeartbeatExecutorThreadPoolSize:             2,
		HeartbeatExecutorExponentialBackOffBound:    10,
		CacheRefreshExecutorThreadPoolSize:          2,
		CacheRefreshExecutorExponentialBackOffBound: 10,
		ServiceUrl: map[string]string{DEFAULT_EUREKA_ZONE: DEFAULT_EUREKA_URL},

		GZipContent:                  true,
		UseDnsForFetchingServiceUrls: false,
		RegisterWithEureka:           true,
		PreferSameZoneEureka:         true,
		AvailabilityZones:            make(map[string]string),
		FilterOnlyUpInstances:        true,
		FetchRegistry:                true,
		DollarReplacement:            "_-",
		EscapeCharReplacement:        "__",
		AllowRedirects:               false,
		OnDemandUpdateStatusChange:   true,
		ClientDataAccept:             EurekaAcceptFull,
	}
}

func (e *EurekaClientConfig) getAvailabilityZones(region string) []string {
	value, ok := e.AvailabilityZones[region]

	if !ok {
		value = DEFAULT_EUREKA_ZONE
	}

	return strings.Split(value, ",")
}

func (e *EurekaClientConfig) getEurekaServerServiceUrls(myZone string) []string {
	serviceUrls, ok := e.ServiceUrl[myZone]
	if !ok {
		serviceUrls = e.ServiceUrl[DEFAULT_EUREKA_ZONE]
	}
	if &serviceUrls == nil {
		return []string{"https://127.0.0.1:8761"}
	}
	return strings.Split(serviceUrls, ",")
}
