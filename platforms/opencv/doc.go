/*
Package opencv contains the Gobot drivers for opencv.

Installing:

This package requires `opencv` to be installed on your system

Then you can install the package with:

	go get github.com/stevebargelt/gobot && go install github.com/stevebargelt/gobot/platforms/opencv

Example:

	package main

	import (
		cv "gobot.io/x/go-opencv/opencv"
		"github.com/stevebargelt/gobot"
		"github.com/stevebargelt/gobot/platforms/opencv"
	)

	func main() {
		window := opencv.NewWindowDriver()
		camera := opencv.NewCameraDriver(0)

		work := func() {
			camera.On(camera.Event("frame"), func(data interface{}) {
				window.ShowImage(data.(*cv.IplImage))
			})
		}

		robot := gobot.NewRobot("cameraBot",
			[]gobot.Device{window, camera},
			work,
		)

		robot.Start()
	}

For further information refer to opencv README:
https://github.com/hybridgroup/gobot/blob/master/platforms/opencv/README.md
*/
package opencv // import "github.com/stevebargelt/gobot/platforms/opencv"
