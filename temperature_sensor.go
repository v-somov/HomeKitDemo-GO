package main

import (
	"encoding/json"
	"fmt"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type TemperatureAPIResponse struct {
	Temperature float64
}

func parseTemperature(body []byte) (*TemperatureAPIResponse, error) {
	var s = new(TemperatureAPIResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func getTemperature() float64 {
	res, err := http.Get("YOUR_IoT_DEVICE_IP_ADRESS/temperature")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	s, err := parseTemperature([]byte(body))
	return s.Temperature
}

func main() {
	info := accessory.Info{
		Name:         "Temp Sensor",
		Manufacturer: "Vlad Somov",
	}

	// When creating NewTemperatureSensor you should pass (device info, current temp, min temp, max temp and step value)
	acc := accessory.NewTemperatureSensor(info, getTemperature(), 0, 100, 0.1)

	t, err := hc.NewIPTransport(hc.Config{Pin: "11192123"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			log.Println("Current Temp:", getTemperature())
			acc.TempSensor.CurrentTemperature.SetValue(getTemperature())
			time.Sleep(5 * time.Second)
		}
	}()

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
