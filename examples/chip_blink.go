// +build example
//
// Do not build by default.

package main

import (
	"time"

	"github.com/stevebargelt/gobot"
	"github.com/stevebargelt/gobot/drivers/gpio"
	"github.com/stevebargelt/gobot/platforms/chip"
)

func main() {
	chipAdaptor := chip.NewAdaptor()
	led := gpio.NewLedDriver(chipAdaptor, "XIO-P6")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{chipAdaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
