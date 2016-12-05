package main

import (
	"fmt"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"io/ioutil"
	"log"
	"net/http"
)

var rootURL = "YOUR_IoT_DEVICE_IP_ADDRESS"

func sendRequestToCoffeMachine(value int) {
	resp, err := http.Get(fmt.Sprintf("%s/digital/13/%v", rootURL, value))
	if err != nil {
		log.Println("Error:", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("Response:", string(body))
}

func turnCoffeeOn() {
	log.Println("Turn Light On")
	sendRequestToCoffeMachine(1)
}

func turnCoffeeOff() {
	log.Println("Turn Light Off")
	sendRequestToCoffeMachine(0)
}

func main() {
	switchInfo := accessory.Info{
		Name:         "Coffee",
		SerialNumber: "12345678",
		Manufacturer: "Vlad Somov",
		Model:        "1",
	}
	acc := accessory.NewSwitch(switchInfo)

	t, err := hc.NewIPTransport(hc.Config{Pin: "12344321"}, acc.Accessory)

	if err != nil {
		log.Panic(err)
	}

	acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			turnCoffeeOn()
		} else {
			turnCoffeeOff()
		}
	})

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
