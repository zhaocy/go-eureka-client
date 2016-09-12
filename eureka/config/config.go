package config

import "strings"

const (
	EurekaAcceptFull    = "full"
	EurekaAcceptCompact = "compact"
	DEFAULT_EUREKA_URL  = "http://localhost:8761/eureka/"
	DEFAULT_EUREKA_ZONE = "defaultZone"
	UP                  = "UP"
	DOWN                = "DOWN"
	STARTING            = "STARTING"
)

type ApplicationConfig struct {
	Name    string `yaml:"name" json:"name" xml:"name"`
	Port    int    `yaml:"port" json:"port" xml:"port"`
	Secured bool   `yaml:"secured" json:"secured" xml:"secured"`
}
type EurekaConfig struct {
	Eureka      Eureka
	Application ApplicationConfig `yaml:"application" json:"application" xml:"application"`
}
type Eureka struct {
	Client   EurekaClientConfig
	Instance EurekaInstanceConfig
}
type EurekaTransportConfig struct {
	SessionedClientReconnectIntervalSeconds           int     `yaml:"sessionedClientReconnectIntervalSeconds" json:"sessionedClientReconnectIntervalSeconds" xml:"sessionedClientReconnectIntervalSeconds"`
	RetryableClientQuarantineRefreshPercentage        float64 `yaml:"retryableClientQuarantineRefreshPercentage" json:"retryableClientQuarantineRefreshPercentage" xml:"retryableClientQuarantineRefreshPercentage"`
	BootstrapResolverRefreshIntervalSeconds           int     `yaml:"bootstrapResolverRefreshIntervalSeconds" json:"bootstrapResolverRefreshIntervalSeconds" xml:"bootstrapResolverRefreshIntervalSeconds"`
	ApplicationsResolverDataStalenessThresholdSeconds int     `yaml:"applicationsResolverDataStalenessThresholdSeconds" json:"applicationsResolverDataStalenessThresholdSeconds" xml:"applicationsResolverDataStalenessThresholdSeconds"`
	AsyncResolverRefreshIntervalMs                    int     `yaml:"asyncResolverRefreshIntervalMs" json:"asyncResolverRefreshIntervalMs" xml:"asyncResolverRefreshIntervalMs"`
	AsyncResolverWarmUpTimeoutMs                      int     `yaml:"asyncResolverWarmUpTimeoutMs" json:"asyncResolverWarmUpTimeoutMs" xml:"asyncResolverWarmUpTimeoutMs"`
	AsyncExecutorThreadPoolSize                       int     `yaml:"asyncExecutorThreadPoolSize" json:"asyncExecutorThreadPoolSize" xml:"asyncExecutorThreadPoolSize"`
	ReadClusterVip                                    string  `yaml:"readClusterVip" json:"readClusterVip" xml:"readClusterVip"`
	BootstrapResolverForQuery                         bool    `yaml:"bootstrapResolverForQuery" json:"bootstrapResolverForQuery" xml:"bootstrapResolverForQuery"`
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
	Appname                          string `yaml:"appname" json:"appname" xml:"appname"`
	AppGroupName                     string `yaml:"appGroupName" json:"appGroupName" xml:"appGroupName"`
	InstanceEnabledOnit              bool   `yaml:"instanceEnabledOnit" json:"instanceEnabledOnit" xml:"instanceEnabledOnit"`
	NonSecurePort                    int    `yaml:"nonSecurePort" json:"nonSecurePort" xml:"nonSecurePort"`
	SecurePort                       int    `yaml:"securePort" json:"securePort" xml:"securePort"`
	NonSecurePortEnabled             bool   `yaml:"nonSecurePortEnabled" json:"nonSecurePortEnabled" xml:"nonSecurePortEnabled"`
	SecurePortEnabled                bool   `yaml:"securePortEnabled" json:"securePortEnabled" xml:"securePortEnabled"`
	LeaseRenewalIntervalInSeconds    uint   `yaml:"leaseRenewalIntervalInSeconds" json:"leaseRenewalIntervalInSeconds" xml:"leaseRenewalIntervalInSeconds"`
	LeaseExpirationDurationInSeconds int    `yaml:"leaseExpirationDurationInSeconds" json:"leaseExpirationDurationInSeconds" xml:"leaseExpirationDurationInSeconds"`
	VirtualHostName                  string `yaml:"virtualHostName" json:"virtualHostName" xml:"virtualHostName"`
	InstanceId                       string `yaml:"instanceId" json:"instanceId" xml:"instanceId"`
	SecureVirtualHostName            string `yaml:"secureVirtualHostName" json:"secureVirtualHostName" xml:"secureVirtualHostName"`
	ASGName                          string `yaml:"aSGName" json:"aSGName" xml:"aSGName"`
	//	DataCenterInfo                   DataCenterInfoConfig `yaml:"dataCenterInfo" json:"dataCenterInfo" xml:"dataCenterInfo"`
	IpAddress            string            `yaml:"ipAddress" json:"ipAddress" xml:"ipAddress"`
	StatusPageUrlPath    string            `yaml:"statusPageUrlPath" json:"statusPageUrlPath" xml:"statusPageUrlPath"`
	StatusPageUrl        string            `yaml:"statusPageUrl" json:"statusPageUrl" xml:"statusPageUrl"`
	HomePageUrlPath      string            `yaml:"homePageUrlPath" json:"homePageUrlPath" xml:"homePageUrlPath"`
	HomePageUrl          string            `yaml:"homePageUrl" json:"homePageUrl" xml:"homePageUrl"`
	HealthCheckUrlPath   string            `yaml:"healthCheckUrlPath" json:"healthCheckUrlPath" xml:"healthCheckUrlPath"`
	HealthCheckUrl       string            `yaml:"healthCheckUrl" json:"healthCheckUrl" xml:"healthCheckUrl"`
	SecureHealthCheckUrl string            `yaml:"secureHealthCheckUrl" json:"secureHealthCheckUrl" xml:"secureHealthCheckUrl"`
	Namespace            string            `yaml:"namespace" json:"namespace" xml:"namespace"`
	Hostname             string            `yaml:"hostname" json:"hostname" xml:"hostname"`
	PreferIpAddress      bool              `yaml:"preferIpAddress" json:"preferIpAddress" xml:"preferIpAddress"`
	InitialStatus        string            `yaml:"initialStatus" json:"initialStatus" xml:"initialStatus"`
	MetadataMap          map[string]string `yaml:"metadataMap" json:"metadataMap" xml:"metadataMap"`
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
		InitialStatus:                    UP,
	}
	ins.VirtualHostName = ins.Appname

	return ins
}

type EurekaClientConfig struct {
	Transport                                     EurekaTransportConfig `yaml:"transport" json:"transport" xml:"transport"`
	RegistryFetchIntervalSeconds                  int                   `yaml:"registryFetchIntervalSeconds" json:"registryFetchIntervalSeconds" xml:"registryFetchIntervalSeconds"`
	InstanceInfoReplicationIntervalSeconds        int                   `yaml:"instanceInfoReplicationIntervalSeconds" json:"instanceInfoReplicationIntervalSeconds" xml:"instanceInfoReplicationIntervalSeconds"`
	InitialInstanceInfoReplicationIntervalSeconds int                   `yaml:"initialInstanceInfoReplicationIntervalSeconds" json:"initialInstanceInfoReplicationIntervalSeconds" xml:"initialInstanceInfoReplicationIntervalSeconds"`
	EurekaServiceUrlPollIntervalSeconds           int                   `yaml:"eurekaServiceUrlPollIntervalSeconds" json:"eurekaServiceUrlPollIntervalSeconds" xml:"eurekaServiceUrlPollIntervalSeconds"`
	EurekaServerReadTimeoutSeconds                int                   `yaml:"eurekaServerReadTimeoutSeconds" json:"eurekaServerReadTimeoutSeconds" xml:"eurekaServerReadTimeoutSeconds"`
	EurekaServerConnectTimeoutSeconds             int                   `yaml:"eurekaServerConnectTimeoutSeconds" json:"eurekaServerConnectTimeoutSeconds" xml:"eurekaServerConnectTimeoutSeconds"`
	BackupRegistryImpl                            string                `yaml:"backupRegistryImpl" json:"backupRegistryImpl" xml:"backupRegistryImpl"`
	EurekaServerTotalConnections                  int                   `yaml:"eurekaServerTotalConnections" json:"eurekaServerTotalConnections" xml:"eurekaServerTotalConnections"`
	EurekaServerTotalConnectionsPerHost           int                   `yaml:"eurekaServerTotalConnectionsPerHost" json:"eurekaServerTotalConnectionsPerHost" xml:"eurekaServerTotalConnectionsPerHost"`
	EurekaServerURLContext                        string                `yaml:"eurekaServerURLContext" json:"eurekaServerURLContext" xml:"eurekaServerURLContext"`
	EurekaServerPort                              string                `yaml:"eurekaServerPort" json:"eurekaServerPort" xml:"eurekaServerPort"`
	EurekaServerDNSName                           string                `yaml:"eurekaServerDNSName" json:"eurekaServerDNSName" xml:"eurekaServerDNSName"`
	Region                                        string                `yaml:"region" json:"region" xml:"region"`
	EurekaConnectionIdleTimeoutSeconds            int                   `yaml:"eurekaConnectionIdleTimeoutSeconds" json:"eurekaConnectionIdleTimeoutSeconds" xml:"eurekaConnectionIdleTimeoutSeconds"`
	RegistryRefreshSingleVipAddress               string                `yaml:"registryRefreshSingleVipAddress" json:"registryRefreshSingleVipAddress" xml:"registryRefreshSingleVipAddress"`
	HeartbeatExecutorThreadPoolSize               int                   `yaml:"heartbeatExecutorThreadPoolSize" json:"heartbeatExecutorThreadPoolSize" xml:"heartbeatExecutorThreadPoolSize"`
	HeartbeatExecutorExponentialBackOffBound      int                   `yaml:"heartbeatExecutorExponentialBackOffBound" json:"heartbeatExecutorExponentialBackOffBound" xml:"heartbeatExecutorExponentialBackOffBound"`
	CacheRefreshExecutorThreadPoolSize            int                   `yaml:"cacheRefreshExecutorThreadPoolSize" json:"cacheRefreshExecutorThreadPoolSize" xml:"cacheRefreshExecutorThreadPoolSize"`
	CacheRefreshExecutorExponentialBackOffBound   int                   `yaml:"cacheRefreshExecutorExponentialBackOffBound" json:"cacheRefreshExecutorExponentialBackOffBound" xml:"cacheRefreshExecutorExponentialBackOffBound"`
	GZipContent                                   bool                  `yaml:"gZipContent" json:"gZipContent" xml:"gZipContent"`
	UseDnsForFetchingServiceUrls                  bool                  `yaml:"useDnsForFetchingServiceUrls" json:"useDnsForFetchingServiceUrls" xml:"useDnsForFetchingServiceUrls"`
	RegisterWithEureka                            bool                  `yaml:"registerWithEureka" json:"registerWithEureka" xml:"registerWithEureka"`
	PreferSameZoneEureka                          bool                  `yaml:"preferSameZoneEureka" json:"preferSameZoneEureka" xml:"preferSameZoneEureka"`
	LogDeltaDiff                                  bool                  `yaml:"logDeltaDiff" json:"logDeltaDiff" xml:"logDeltaDiff"`
	DisableDelta                                  bool                  `yaml:"disableDelta" json:"disableDelta" xml:"disableDelta"`
	FetchRemoteRegionsRegistry                    string                `yaml:"fetchRemoteRegionsRegistry" json:"fetchRemoteRegionsRegistry" xml:"fetchRemoteRegionsRegistry"`
	FilterOnlyUpInstances                         bool                  `yaml:"filterOnlyUpInstances" json:"filterOnlyUpInstances" xml:"filterOnlyUpInstances"`
	FetchRegistry                                 bool                  `yaml:"fetchRegistry" json:"fetchRegistry" xml:"fetchRegistry"`
	DollarReplacement                             string                `yaml:"dollarReplacement" json:"dollarReplacement" xml:"dollarReplacement"`
	EscapeCharReplacement                         string                `yaml:"escapeCharReplacement" json:"escapeCharReplacement" xml:"escapeCharReplacement"`
	AllowRedirects                                bool                  `yaml:"allowRedirects" json:"allowRedirects" xml:"allowRedirects"`
	OnDemandUpdateStatusChange                    bool                  `yaml:"onDemandUpdateStatusChange" json:"onDemandUpdateStatusChange" xml:"onDemandUpdateStatusChange"`
	EncoderName                                   string                `yaml:"encoderName" json:"encoderName" xml:"encoderName"`
	DecoderName                                   string                `yaml:"decoderName" json:"decoderName" xml:"decoderName"`
	ClientDataAccept                              string                `yaml:"clientDataAccept" json:"clientDataAccept" xml:"clientDataAccept"`
	AvailabilityZones                             map[string]string     `yaml:"availabilityZones" json:"availabilityZones" xml:"availabilityZones"`
	ServiceUrl                                    map[string]string     `yaml:"serviceUrl" json:"serviceUrl" xml:"serviceUrl"`
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

func (e *EurekaClientConfig) GetAvailabilityZones(region string) []string {
	value, ok := e.AvailabilityZones[region]

	if !ok {
		value = DEFAULT_EUREKA_ZONE
	}

	return strings.Split(value, ",")
}

func (e *EurekaClientConfig) GetEurekaServerServiceUrls(myZone string) []string {
	serviceUrls, ok := e.ServiceUrl[myZone]
	if !ok {
		serviceUrls = e.ServiceUrl[DEFAULT_EUREKA_ZONE]
	}
	if &serviceUrls == nil {
		return []string{"https://127.0.0.1:8761"}
	}
	return strings.Split(serviceUrls, ",")
}
