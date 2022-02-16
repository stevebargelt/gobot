// +build example
//
// Do not build by default.

/*
 How to run
 Pass serial port to use as the first param:

	go run examples/firmata_curie_imu_shock_detect.go /dev/ttyACM0
*/

package main

import (
	"log"
	"os"
	"time"

	"github.com/stevebargelt/gobot"
	"github.com/stevebargelt/gobot/drivers/gpio"
	"github.com/stevebargelt/gobot/platforms/firmata"
	"github.com/stevebargelt/gobot/platforms/intel-iot/curie"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	led := gpio.NewLedDriver(firmataAdaptor, "13")
	imu := curie.NewIMUDriver(firmataAdaptor)

	work := func() {
		imu.On("Shock", func(data interface{}) {
			log.Println("Shock", data)
		})

		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})

		imu.EnableShockDetection(true)
	}

	robot := gobot.NewRobot("curieBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{imu, led},
		work,
	)

	robot.Start()
}
