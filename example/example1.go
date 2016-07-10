package main

import (
	"encoding/xml"
	"net/http"

	"github.com/tietang/go-eureka-client/eureka"
)

func main() {
	client := eureka.CreateEurekaClientByYaml("application.yml")
	data, _ := xml.Marshal(client.InstanceInfo)
	println(string(data))
	client.Start()
	http.ListenAndServe(":8900", nil)
}
