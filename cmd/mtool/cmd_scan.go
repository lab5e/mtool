package main

import (
	"log"
	"os"
	"time"

	"github.com/goburrow/modbus"
	"github.com/lab5e/mtool/internal/util"
)

type scanCmd struct {
	Addr        uint16        `long:"addr" default:"1" description:"which address to use for scan"`
	Type        string        `long:"type" default:"holding" description:"what to use for scan" choice:"holding" choice:"input" choice:"coil"`
	ScanTimeout time.Duration `long:"scan-timeout" default:"100ms" description:"timeout when scanning each id"`
}

func (s *scanCmd) Execute([]string) error {
	handler := modbus.NewRTUClientHandler(opt.Device)
	handler.BaudRate = opt.Baud
	handler.DataBits = opt.DataBits
	handler.Parity = opt.Parity
	handler.StopBits = opt.StopBits
	handler.Timeout = s.ScanTimeout
	handler.SlaveId = 1

	if opt.Debug {
		handler.Logger = log.New(os.Stderr, "", log.LstdFlags)
	}

	err := handler.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()

	client := modbus.NewClient(handler)

	log.Printf("scan will take %s", 255*s.ScanTimeout)
	for i := byte(1); i < 255; i++ {
		handler.SlaveId = i
		res, err := client.ReadHoldingRegisters(s.Addr, 1)
		if err != nil {
			if err.Error() == "serial: timeout" {
				continue
			}
			log.Printf("%d: %v", i, err)
			continue
		}
		log.Printf("Found device with id [%d] (read %v from addr %d)", i, util.BytesToStringArray(res, 16), s.Addr)
	}
	return nil
}
