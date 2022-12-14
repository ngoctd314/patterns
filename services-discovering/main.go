package main

import (
	"github.com/ArthurHlt/go-eureka-client/eureka"
)

func main() {

	client := eureka.NewClient([]string{
		"http://127.0.0.1:8761/eureka", //From a spring boot based eureka server
		// add others servers here
	})
	instance := eureka.NewInstanceInfo("test.com", "test", "69.172.200.235", 80, 30, false) //Create a new instance to register
	instance.Metadata = &eureka.MetaData{
		Map: make(map[string]string),
	}
	instance.Metadata.Map["foo"] = "bar"                  //add metadata for example
	client.RegisterInstance("myapp", instance)            // Register new instance in your eureka(s)
	client.GetApplication(instance.App)                   // retrieve the application "test"
	client.GetInstance(instance.App, instance.HostName)   // retrieve the instance from "test.com" inside "test"" app
	client.SendHeartbeat(instance.App, instance.HostName) // say to eureka that your app is alive (here you must send heartbeat before 30 sec)
	_, _ = client.GetApplications()                       // Retrieves all applications from eureka server(s)
}
