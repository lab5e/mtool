package main

import (
	"fmt"
	"log"
)

type writeCoilsCmd struct {
	Addr  uint16 `long:"addr" description:"address" required:"yes"`
	Value uint16 `long:"value" description:"value"`
}

func (wc *writeCoilsCmd) Execute([]string) error {
	res, err := client().WriteSingleCoil(wc.Addr, wc.Value)
	if err != nil {
		log.Fatalf("error writing coils: %v", err)
	}

	fmt.Printf("slave=%d addr=%d res={%s}\n", opt.SlaveID, wc.Addr, bytesToString(res, opt.OutputBase))
	return nil
}
