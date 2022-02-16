package spi

import (
	"testing"

	"github.com/stevebargelt/gobot"
	"github.com/stevebargelt/gobot/drivers/aio"
	"github.com/stevebargelt/gobot/gobottest"
)

var _ gobot.Driver = (*MCP3008Driver)(nil)

// must implement the AnalogReader interface
var _ aio.AnalogReader = (*MCP3008Driver)(nil)

func initTestMCP3008Driver() *MCP3008Driver {
	d := NewMCP3008Driver(&TestConnector{})
	return d
}

func TestMCP3008DriverStart(t *testing.T) {
	d := initTestMCP3008Driver()
	gobottest.Assert(t, d.Start(), nil)
}

func TestMCP3008DriverHalt(t *testing.T) {
	d := initTestMCP3008Driver()
	d.Start()
	gobottest.Assert(t, d.Halt(), nil)
}

func TestMCP3008DriverRead(t *testing.T) {
	d := initTestMCP3008Driver()
	d.Start()

	// TODO: actual read test
}
