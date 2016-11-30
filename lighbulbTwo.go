package main

import (
	"fmt"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"io/ioutil"
	"log"
	"net/http"
)

func sendRequestToLighbulb(value int) {
	resp, err := http.Get(fmt.Sprintf("http://192.168.1.94/digital/13/%v", value))
	if err != nil {
		log.Println("Error:", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("Response:", string(body))
}

func turnLightOn() {
	log.Println("Turn Light On")
	sendRequestToLighbulb(1)
}

func turnLightOff() {
	log.Println("Turn Light Off")
	sendRequestToLighbulb(0)
}

func brightnessChanged(value int) {
	log.Println("Brightness changed to: ", value)
	resp, err := http.Get(fmt.Sprintf("http://192.168.1.94/analog/13/%v", int(value*(255/100))))
	if err != nil {
		log.Println("Error:", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("Response:", string(body))
}

func main() {
	lightBulbInfo := accessory.Info{
		Name:         "Second Lightbulb",
		Manufacturer: "Vlad Somov",
	}

	acc := accessory.NewLightbulb(lightBulbInfo)

	acc.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			turnLightOn()
		} else {
			turnLightOff()
		}
	})

	acc.Lightbulb.Brightness.OnValueRemoteUpdate(func(value int) {
		brightnessChanged(value)
	})

	bulb, err := hc.NewIPTransport(hc.Config{Pin: "32232232"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		bulb.Stop()
	})

	bulb.Start()
}
