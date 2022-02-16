package sprkplus

import (
	"github.com/stevebargelt/gobot"
	"github.com/stevebargelt/gobot/platforms/ble"
	"github.com/stevebargelt/gobot/platforms/sphero/ollie"
)

// Driver represents a Sphero SPRK+
type SPRKPlusDriver struct {
	*ollie.Driver
}

// NewDriver creates a Driver for a Sphero SPRK+
func NewDriver(a ble.BLEConnector) *SPRKPlusDriver {
	d := ollie.NewDriver(a)
	d.SetName(gobot.DefaultName("SPRKPlus"))

	return &SPRKPlusDriver{
		Driver: d,
	}
}
