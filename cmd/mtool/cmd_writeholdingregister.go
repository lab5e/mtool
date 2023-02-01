package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/lab5e/mtool/internal/util"
)

type writeHoldingRegisterCmd struct {
	Addr       uint16 `long:"addr" description:"address" required:"yes"`
	Value      uint16 `long:"value" description:"value"`
	Bytes      bool   `long:"bytes" description:"express value as bytes instead of int16"`
	OutputBase int    `long:"base" default:"10" description:"output base" choice:"2" choice:"8" choice:"10" choice:"16"`
}

func (wh *writeHoldingRegisterCmd) Execute([]string) error {
	res, err := client().WriteSingleRegister(mapAddr(wh.Addr), wh.Value)
	if err != nil {
		log.Fatalf("error writing holding register: %v", err)
	}

	var values []string
	if wh.Bytes {
		values = util.BytesToStringArray(res, wh.OutputBase)
	} else {
		values, err = util.BytesToInt16StringArray(res, wh.OutputBase)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("deviceID=%d addr=%d res={%s}\n", opt.DeviceID, wh.Addr, strings.Join(values, ", "))
	return nil
}
