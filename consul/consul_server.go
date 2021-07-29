package main

import (
	"fmt"
	"github.com/gorilla/mux"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"net"
	"net/http"
)

var count int64

func consulCheck(w http.ResponseWriter, r *http.Request) {
	s := "consulCheck" + fmt.Sprint(count) + "remote:" + r.RemoteAddr + " " + r.URL.String()
	fmt.Println(s)
	fmt.Fprintln(w, s)
	count++
}
func registerServer() {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error:", err)
	}
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "test1"
	registration.Name = "testName"
	registration.Port = 9527
	registration.Tags = []string{"v1000"}
	registration.Address = localIP()
	checkPort := 9527
	registration.Check = &consulapi.AgentServiceCheck{

		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "30s",
	}
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalln("register server error:", err)
	}

}
func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("you are visiting health check api"))
}
func main() {
	//router := mux.NewRouter()
	//router.HandleFunc("/people",Handler )
	//err:=http.ListenAndServe(":9527",router)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	registerServer()
	router := mux.NewRouter()
	router.HandleFunc("/people", Handler)
	router.HandleFunc("/check", consulCheck)
	err := http.ListenAndServe(fmt.Sprintf(":9527"), router)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
