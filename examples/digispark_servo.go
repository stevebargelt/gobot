// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/stevebargelt/gobot"
	"github.com/stevebargelt/gobot/drivers/gpio"
	"github.com/stevebargelt/gobot/platforms/digispark"
)

func main() {
	digisparkAdaptor := digispark.NewAdaptor()
	servo := gpio.NewServoDriver(digisparkAdaptor, "0")

	work := func() {
		gobot.Every(1*time.Second, func() {
			i := uint8(gobot.Rand(180))
			fmt.Println("Turning", i)
			servo.Move(i)
		})
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{digisparkAdaptor},
		[]gobot.Device{servo},
		work,
	)

	robot.Start()
}
