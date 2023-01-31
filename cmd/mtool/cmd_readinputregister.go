package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lab5e/mtool/internal/util"
)

type readInputRegisterCmd struct {
	Addrs       []uint16      `long:"addr" description:"address" required:"yes"`
	Count       uint16        `long:"count" default:"1" description:"count"`
	RepeatEvery time.Duration `long:"repeat" default:"0" description:"repeat interval, if zero no repeat"`
	JSON        bool          `long:"json" description:"format data as JSON"`
	Bytes       bool          `long:"bytes" description:"list values as bytes instead of int16"`
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

			var values []string

			// select between bytes and int16
			if ri.Bytes {
				values = util.BytesToStringArray(res, opt.OutputBase)
			} else {
				values, err = util.BytesToInt16StringArray(res, opt.OutputBase)
				if err != nil {
					return err
				}
			}

			d := dataPoint{
				Time:      time.Now().UTC(),
				ValueType: "input_register",
				DeviceID:  opt.DeviceID,
				Addr:      addr,
				Count:     ri.Count,
				Data:      values,
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
