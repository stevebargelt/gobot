/*
Package firmata provides the Gobot adaptor for microcontrollers that support the Firmata protocol.

Installing:

	go get -d -u github.com/stevebargelt/gobot/... && go get github.com/stevebargelt/gobot/platforms/firmata

Example:

	package main

	import (
		"time"

		"github.com/stevebargelt/gobot"
		"github.com/stevebargelt/gobot/drivers/gpio"
		"github.com/stevebargelt/gobot/platforms/firmata"
	)

	func main() {
		firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
		led := gpio.NewLedDriver(firmataAdaptor, "13")

		work := func() {
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		}

		robot := gobot.NewRobot("bot",
			[]gobot.Connection{firmataAdaptor},
			[]gobot.Device{led},
			work,
		)

		robot.Start()
	}

For further information refer to firmata readme:
https://github.com/hybridgroup/gobot/blob/master/platforms/firmata/README.md
*/
package firmata // import "github.com/stevebargelt/gobot/platforms/firmata"
