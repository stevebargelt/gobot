package spi

import (
	"testing"

	"github.com/stevebargelt/gobot"
	"github.com/stevebargelt/gobot/drivers/aio"
	"github.com/stevebargelt/gobot/gobottest"
)

var _ gobot.Driver = (*MCP3202Driver)(nil)

// must implement the AnalogReader interface
var _ aio.AnalogReader = (*MCP3202Driver)(nil)

func initTestMCP3202Driver() *MCP3202Driver {
	d := NewMCP3202Driver(&TestConnector{})
	return d
}

func TestMCP3202DriverStart(t *testing.T) {
	d := initTestMCP3202Driver()
	gobottest.Assert(t, d.Start(), nil)
}

func TestMCP3202DriverHalt(t *testing.T) {
	d := initTestMCP3202Driver()
	d.Start()
	gobottest.Assert(t, d.Halt(), nil)
}

func TestMCP3202DriverRead(t *testing.T) {
	d := initTestMCP3202Driver()
	d.Start()

	// TODO: actual read test
}
