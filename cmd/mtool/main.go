// Package main implements a modbus command line tool
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goburrow/modbus"
	"github.com/jessevdk/go-flags"
)

//lint:file-ignore SA5008 Ignore duplicate struct tags
var opt struct {
	Device     string `long:"device" env:"MTOOL_DEVICE" description:"serial device" required:"yes"`
	Baud       int    `long:"baud" env:"MTOOL_BAUD" default:"9600" description:"baud rate"`
	DataBits   int    `long:"databits" env:"MTOOL_DATABITS" default:"8" description:"data bits"`
	Parity     string `long:"parity" env:"MTOOL_PARITY" default:"N" description:"parity" choice:"N" choice:"E" choice:"O"`
	StopBits   int    `long:"stop" env:"MTOOL_STOPBITS" default:"1" description:"stop bits" choice:"1" choice:"2"`
	DeviceID   byte   `long:"id" default:"1" description:"device id" required:"yes"`
	OutputBase int    `long:"base" env:"MTOOL_OUTPUT_BASE" default:"10" description:"output base" choice:"2" choice:"8" choice:"10" choice:"16"`

	ReadInputRegister    readInputRegisterCmd    `command:"ri" description:"read input register"`
	ReadHoldingRegister  readHoldingRegisterCmd  `command:"rh" description:"read holding register"`
	WriteHoldingRegister writeHoldingRegisterCmd `command:"wh" description:"write holding register"`
	ReadDiscreteInput    readDiscreteInputCmd    `command:"rd" description:"read discrete input"`
	ReadCoils            readCoilsCmd            `command:"rc" description:"read coils"`
}

func main() {
	// parse and execute command line
	p := flags.NewParser(&opt, flags.Default)
	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			switch flagsErr.Type {
			case flags.ErrHelp:
				os.Exit(0)

			case flags.ErrCommandRequired:
				os.Exit(1)

			case flags.ErrUnknownCommand:
				os.Exit(1)

			case flags.ErrRequired:
				os.Exit(1)

			case flags.ErrUnknownFlag:
				os.Exit(1)

			case flags.ErrMarshal:
				os.Exit(1)

			case flags.ErrExpectedArgument:
				os.Exit(1)

			default:
				fmt.Printf("%v [%d]\n", err, flagsErr.Type)
				os.Exit(0)
			}
		}
		os.Exit(1)
	}
}

func client() modbus.Client {
	handler := modbus.NewRTUClientHandler(opt.Device)
	handler.BaudRate = opt.Baud
	handler.DataBits = opt.DataBits
	handler.Parity = opt.Parity
	handler.StopBits = opt.StopBits
	handler.SlaveId = opt.DeviceID

	err := handler.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return modbus.NewClient(handler)
}
