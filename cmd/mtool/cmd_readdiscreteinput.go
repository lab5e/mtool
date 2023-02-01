package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lab5e/mtool/internal/util"
)

type readDiscreteInputCmd struct {
	Addrs       []uint16      `long:"addr" description:"address" required:"yes"`
	Count       uint16        `long:"count" default:"1" description:"count"`
	RepeatEvery time.Duration `long:"repeat" default:"0" description:"repeat interval, if zero no repeat"`
	JSON        bool          `long:"json" description:"format data as JSON"`
	OutputBase  int           `long:"base" default:"2" description:"output base" choice:"2" choice:"8" choice:"10" choice:"16"`
}

func (rd *readDiscreteInputCmd) Execute([]string) error {
	client := client()

	for {
		for _, addr := range rd.Addrs {
			res, err := client.ReadInputRegisters(mapAddr(addr), rd.Count)
			if err != nil {
				return err
			}
			if err != nil {
				log.Fatal(err)
			}

			d := dataPoint{
				Time:      time.Now().UTC(),
				ValueType: "discrete_input",
				DeviceID:  opt.DeviceID,
				Addr:      addr,
				Count:     rd.Count,
				Data:      util.BytesToStringArray(res, rd.OutputBase),
			}

			// output JSON or string
			if rd.JSON {
				fmt.Println(d.JSON())
			} else {
				fmt.Println(d.String())
			}
		}

		if rd.RepeatEvery != 0 {
			time.Sleep(rd.RepeatEvery)
			continue
		}
		return nil
	}
}
