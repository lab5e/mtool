package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/lab5e/mtool/internal/util"
)

type writeCoilsCmd struct {
	Addr       uint16 `long:"addr" description:"address" required:"yes"`
	State      string `long:"state" description:"state" choice:"on" choice:"off" required:"yes"`
	Bytes      bool   `long:"bytes" description:"express value as bytes instead of int16"`
	OutputBase int    `long:"base" default:"2" description:"output base" choice:"2" choice:"8" choice:"10" choice:"16"`
}

// note that this has not been tested since we do not have any device to test against.
func (wc *writeCoilsCmd) Execute([]string) error {
	value := uint16(0)
	if wc.State == "on" {
		value = 0xff00
	}

	res, err := client().WriteSingleCoil(wc.Addr, value)
	if err != nil {
		log.Fatalf("error writing coil: %v", err)
	}

	var values []string
	if wc.Bytes {
		values = util.BytesToStringArray(res, wc.OutputBase)
	} else {
		values, err = util.BytesToInt16StringArray(res, wc.OutputBase)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("deviceID=%d addr=%d res={%s}\n", opt.DeviceID, wc.Addr, strings.Join(values, ", "))
	return nil
}
