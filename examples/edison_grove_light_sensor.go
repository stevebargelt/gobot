// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/stevebargelt/gobot"
	"github.com/stevebargelt/gobot/drivers/aio"
	"github.com/stevebargelt/gobot/platforms/intel-iot/edison"
)

func main() {
	board := edison.NewAdaptor()
	sensor := aio.NewGroveLightSensorDriver(board, "0")

	work := func() {
		sensor.On(sensor.Event("data"), func(data interface{}) {
			fmt.Println("sensor", data)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{board},
		[]gobot.Device{sensor},
		work,
	)

	robot.Start()
}
