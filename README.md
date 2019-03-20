Go Eureka Client
================
[![Build Status](https://travis-ci.org/tietang/go-eureka-client.svg?branch=master)](https://travis-ci.org/tietang/go-eureka-client)
[![GoDoc Documentation](http://godoc.org/github.com/tietang/go-eureka-client?status.png)](<https://godoc.org/github.com/tietang/go-eureka-client>)
[![Sourcegraph](https://sourcegraph.com/github.com/tietang/go-eureka-client/-/badge.svg)](https://sourcegraph.com/github.com/tietang/go-eureka-client?badge)
[![CircleCI](https://circleci.com/gh/tietang/go-eureka-client.svg?style=svg)](https://circleci.com/gh/tietang/go-eureka-client)


Based and fork on code from https://github.com/ArthurHlt/go-eureka-client .
## Install
```shell
go get -u github.com/tietang/go-eureka-client/eureka 
```

## Getting started

```go

import (
	"github.com/tietang/go-eureka-client/eureka"
	"time"
)

func main() {
	config := eureka.Config{
		// default timeout is one second
		DialTimeout: time.Second,
	}
	client := eureka.NewClientByConfig([]string{
		"http://127.0.0.1:8761/eureka", //From a spring boot based eureka server
		// add others servers here
	}, config)
	instance := eureka.NewInstanceInfo("test.com", "test", "69.172.200.235", 80, 30, false) //Create a new instance to register
	instance.Metadata = &eureka.MetaData{
		Map: make(map[string]string),
	}
	instance.Metadata.Map["foo"] = "bar"                  //add metadata for example
	client.RegisterInstance("myapp", instance)            // Register new instance in your eureka(s)
	applications, _ := client.GetApplications()           // Retrieves all applications from eureka server(s)
	client.GetApplication(instance.App)                   // retrieve the application "test"
	client.GetInstance(instance.App, instance.HostName)   // retrieve the instance from "test.com" inside "test"" app
	client.SendHeartbeat(instance.App, instance.HostName) // say to eureka that your app is alive (here you must send heartbeat before 30 sec)
}
```

**Note:**
- `appId` here is the name of the app
- `instanceId` is the hostname of the app
- When calling `RegisterInstance` the `appId` is needed but not used by eureka, this is not the appId but a whatever value

All these strange behaviour come from Eureka.

## Create Client from a config file

### by ini config file

You can create from a ini file with this form (here we called it `app.ini`):

```ini
[application]
name : go-example
port : 19002


;[management]
;contextPath: /demo

[eureka.instance]
preferIpAddress : true
leaseRenewalIntervalInSeconds : 10s
statusPageUrlPath : ${management.contextPath}/info
healthCheckUrlPath : ${management.contextPath}/health

[eureka.client]
registerWithEureka : true
fetchRegistry : true
serviceUrl.defaultZone : http://172.16.1.248:8761/eureka/

```
And to read it:

```go
import (
    "fmt"
    "encoding/json"
    "github.com/tietang/go-eureka-client/eureka"
    "github.com/tietang/props/ini"
    "github.com/tietang/props/kvs"
)

func main() {
    iniconf := ini.NewIniFileConfigSource("app.ini")
    conf := kvs.NewDefaultCompositeConfigSource(iniconf)
    client := eureka.NewClient(conf)
    client.Start()
    c := make(chan int, 1)
    x := <-c
}

```

### by yaml config file

You can create from a yaml file with this form (here we called it `app.yaml`):

```yaml

application:
  name: go-example
  port: 19002

eureka:
  instance:
    preferIpAddress: true
    leaseRenewalIntervalInSeconds: 10
    statusPageUrlPath: ${management.contextPath}/info
    healthCheckUrlPath: ${management.contextPath}/health
    metadataMap:
      instanceId: ${spring.application.name}:${spring.application.instance_id:${server.port}}

  client:
    registerWithEureka: true
    fetchRegistry: true
    serviceUrl:
      defaultZone: http://localhost:8761/eureka/
```
And to load it:

```go

	client := eureka.CreateEurekaClientByYaml("application.yml")
	client.Start()
	
```



### by json config file

You can create from a json file with this form (here we called it `config.json`):

```json
{
  "config": {
    "certFile": "",
    "keyFile": "",
    "caCertFiles": null,
    "timeout": 1000000000,
    "consistency": ""
  },
  "cluster": {
    "leader": "http://127.0.0.1:8761/eureka",
    "machines": [
      "http://127.0.0.1:8761/eureka"
    ]
  }
}
```

And to load it:

```go
client := NewClientFromFile("config.json")
```
