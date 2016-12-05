#HomeKit Demo GO 
Examples of accessories which implements HomeKit Accessory Protocol(HAP) written in GO [https://github.com/brutella/hc](https://github.com/brutella/hc)

##Getting Started

1. [Install GO](https://golang.org/doc/install)
2. [Setup Go workspace](https://golang.org/doc/code.html#Organization)
3. Clone This Repo
  ```
  cd $GOPATH/src

  git clone https://github.com/vlad1803/HomeKitDemo-GO
  cd HomeKitDemo-GO

  # Download package along with its dependencies
  go get github.com/brutella/hc

  # Run accessory you like for example for lighbulb
  go run lighbulb.go
  ```
4. Setup Your Iot Device IP Address and correct pins to which your IoT aRest api is opened 
5. Pair your device with Home App or any app That uses HomeKit
