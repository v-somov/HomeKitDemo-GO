package main

import (
	"fmt"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"io/ioutil"
	"log"
	"net/http"
)

//If you are using Arduino IDE you can see your IoT device Ip Address when entering serial monitor and waiting for device to connect to WiFi
var rootURL = "YOUR_IoT_DEVICE_IP_ADDRESS"

func sendRequestToLighbulb(value int) {
	// digital/5 means that you will be sending on\off (1 or 0) state to the 5th pin
	resp, err := http.Get(fmt.Sprintf("%s/digital/5/%v", rootURL, value))
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
	brightness := value * (255 / 100)
	// analog/5 means that you will be sending value from 0 to 255 to the 5th pin to change brightness
	resp, err := http.Get(fmt.Sprintf("%s/analog/5/%v", rootURL, int(brightness)))
	if err != nil {
		log.Println("Error:", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("Response:", string(body))
}

func main() {
	// here you provide all infromation about your IoT
	lightBulbInfo := accessory.Info{
		Name:         "Lightbulb",
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

	// This pin you will be using to add this device on your iOS device
	bulb, err := hc.NewIPTransport(hc.Config{Pin: "32191123"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		bulb.Stop()
	})

	bulb.Start()
}
