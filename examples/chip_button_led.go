// +build example
//
// Do not build by default.

package main

import (
	"github.com/stevebargelt/gobot"
	"github.com/stevebargelt/gobot/drivers/gpio"
	"github.com/stevebargelt/gobot/platforms/chip"
)

func main() {
	chipAdaptor := chip.NewAdaptor()
	button := gpio.NewButtonDriver(chipAdaptor, "XIO-P6")
	led := gpio.NewLedDriver(chipAdaptor, "XIO-P7")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			led.On()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			led.Off()
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{chipAdaptor},
		[]gobot.Device{button, led},
		work,
	)

	robot.Start()
}
