// +build example
//
// Do not build by default.

package main

import (
	"time"

	"github.com/stevebargelt/gobot"
	"github.com/stevebargelt/gobot/platforms/audio"
)

func main() {
	e := audio.NewAdaptor()
	laser := audio.NewDriver(e, "./examples/laser.mp3")

	work := func() {
		gobot.Every(2*time.Second, func() {
			laser.Play()
		})
	}

	robot := gobot.NewRobot("soundBot",
		[]gobot.Connection{e},
		[]gobot.Device{laser},
		work,
	)

	robot.Start()
}
