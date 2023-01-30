package main

import (
	"fmt"
	"log"
	"time"
)

type readInputRegisterCmd struct {
	Addrs       []uint16      `long:"addr" description:"address" required:"yes"`
	Count       uint16        `long:"count" default:"1" description:"count"`
	RepeatEvery time.Duration `long:"repeat" default:"0" description:"repeat interval, if zero no repeat"`
	JSON        bool          `long:"json" description:"format data as JSON"`
}

func (ri *readInputRegisterCmd) Execute([]string) error {
	client := client()

	for {
		for _, addr := range ri.Addrs {
			res, err := client.ReadInputRegisters(addr, ri.Count)
			if err != nil {
				return err
			}
			if err != nil {
				log.Fatal(err)
			}

			d := dataPoint{
				Time:      time.Now().UTC(),
				ValueType: "input_register",
				DeviceID:  opt.DeviceID,
				Addr:      addr,
				Data:      bytesToStringArray(res, opt.OutputBase),
			}

			// output JSON or string
			if ri.JSON {
				fmt.Println(d.JSON())
			} else {
				fmt.Println(d.String())
			}
		}

		if ri.RepeatEvery != 0 {
			time.Sleep(ri.RepeatEvery)
			continue
		}
		return nil
	}
}
