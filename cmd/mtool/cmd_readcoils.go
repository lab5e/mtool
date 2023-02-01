package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lab5e/mtool/internal/util"
)

type readCoilsCmd struct {
	Addrs       []uint16      `long:"addr" description:"address" required:"yes"`
	Count       uint16        `long:"count" default:"1" description:"count"`
	RepeatEvery time.Duration `long:"repeat" default:"0" description:"repeat interval, if zero no repeat"`
	JSON        bool          `long:"json" description:"format data as JSON"`
	OutputBase  int           `long:"base" default:"2" description:"output base" choice:"2" choice:"8" choice:"10" choice:"16"`
}

func (rc *readCoilsCmd) Execute([]string) error {
	client := client()
	for {
		for _, addr := range rc.Addrs {
			res, err := client.ReadCoils(mapAddr(addr), rc.Count)
			if err != nil {
				return err
			}
			if err != nil {
				log.Fatal(err)
			}

			d := dataPoint{
				Time:      time.Now().UTC(),
				ValueType: "read_coil",
				DeviceID:  opt.DeviceID,
				Addr:      addr,
				Count:     rc.Count,
				Data:      util.BytesToStringArray(res, rc.OutputBase),
			}

			// output JSON or string
			if rc.JSON {
				fmt.Println(d.JSON())
			} else {
				fmt.Println(d.String())
			}
		}

		if rc.RepeatEvery != 0 {
			time.Sleep(rc.RepeatEvery)
			continue
		}
		return nil
	}
}
