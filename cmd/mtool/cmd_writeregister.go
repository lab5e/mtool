package main

import (
	"fmt"
	"log"
)

type writeHoldingRegisterCmd struct {
	Addr  uint16 `long:"addr" description:"address" required:"yes"`
	Value uint16 `long:"value" description:"value"`
}

func (wh *writeHoldingRegisterCmd) Execute([]string) error {
	res, err := client().WriteSingleRegister(wh.Addr, wh.Value)
	if err != nil {
		log.Fatalf("error writing holding register: %v", err)
	}

	fmt.Printf("deviceID=%d addr=%d res={%s}\n", opt.DeviceID, wh.Addr, bytesToString(res, opt.OutputBase))
	return nil
}
