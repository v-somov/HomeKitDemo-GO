package main

import (
	_ "fmt"
	owm "github.com/briandowns/openweathermap"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"log"
	"os"
	"time"
)

func getTemperature() float64 {
	w, err := owm.NewCurrent("C", "EN")
	if err != nil {
		log.Fatal(err)
	}

	w.CurrentByName("YOUR CITY")
	return w.Main.Temp
}

func main() {
	_ = os.Setenv("OWM_API_KEY", "OPEN_WEATHER_MAP_API_KEY")

	info := accessory.Info{
		Name:         "Outside Weather",
		Manufacturer: "Vlad Somov",
	}

	// When creating NewTemperatureSensor you should pass (device info, current temp, min temp, max temp and step value)
	acc := accessory.NewTemperatureSensor(info, getTemperature(), -100, 100, 0.1)
	t, err := hc.NewIPTransport(hc.Config{Pin: "YOUR_8_DIGIT_ACCESSORY_PASSWORD"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			temperature := getTemperature()
			log.Println("Current Temp:", temperature)
			acc.TempSensor.CurrentTemperature.SetValue(temperature)
			time.Sleep(5 * time.Second)
		}
	}()

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
