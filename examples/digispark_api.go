// +build example
//
// Do not build by default.

package main

import (
	"github.com/stevebargelt/gobot"
	"github.com/stevebargelt/gobot/api"
	"github.com/stevebargelt/gobot/drivers/gpio"
	"github.com/stevebargelt/gobot/platforms/digispark"
)

func main() {
	master := gobot.NewMaster()
	api.NewAPI(master).Start()

	digisparkAdaptor := digispark.NewAdaptor()
	led := gpio.NewLedDriver(digisparkAdaptor, "0")

	robot := gobot.NewRobot("digispark",
		[]gobot.Connection{digisparkAdaptor},
		[]gobot.Device{led},
	)

	master.AddRobot(robot)

	master.Start()
}
