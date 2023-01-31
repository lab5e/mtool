package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lab5e/mtool/internal/util"
)

type readHoldingRegisterCmd struct {
	Addrs       []uint16      `long:"addr" description:"address" required:"yes"`
	Count       uint16        `long:"count" default:"1" description:"count"`
	Int16       bool          `long:"int16" descriiption:"display values as 16 bit ints"`
	RepeatEvery time.Duration `long:"repeat" default:"0" description:"repeat interval, if zero no repeat"`
	JSON        bool          `long:"json" description:"format data as JSON"`
	Bytes       bool          `long:"bytes" description:"list values as bytes instead of int16"`
}

func (rh *readHoldingRegisterCmd) Execute([]string) error {
	client := client()

	for {
		for _, addr := range rh.Addrs {
			res, err := client.ReadHoldingRegisters(addr, rh.Count)

			if err != nil {
				return err
			}
			if err != nil {
				log.Fatal(err)
			}

			var values []string

			// select between bytes and int16
			if rh.Bytes {
				values = util.BytesToStringArray(res, opt.OutputBase)
			} else {
				values, err = util.BytesToInt16StringArray(res, opt.OutputBase)
				if err != nil {
					return err
				}
			}

			d := dataPoint{
				Time:      time.Now().UTC(),
				ValueType: "holding_register",
				DeviceID:  opt.DeviceID,
				Addr:      addr,
				Count:     rh.Count,
				Data:      values,
			}

			// output JSON or string
			if rh.JSON {
				fmt.Println(d.JSON())
			} else {
				fmt.Println(d.String())
			}

			if rh.RepeatEvery != 0 {
				time.Sleep(rh.RepeatEvery)
				continue
			}
			return nil
		}
	}
}
