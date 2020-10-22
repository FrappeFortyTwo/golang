package main

import	"gobot.io/x/gobot/"

// LedDriver represents a digital led
type LedDriver struct {
	pin  string
	name string
	connection DigitalWriter
	high bool
}
