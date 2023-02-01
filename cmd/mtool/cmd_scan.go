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
	Type        string        `long:"type" default:"holding" description:"what to use for scan" choice:"holding" choice:"input" choice:"coil" choice:"discrete"`
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
	for i := byte(0); i < 255; i++ {
		handler.SlaveId = i + 1

		var res []byte
		switch s.Type {
		case "holding":
			res, err = client.ReadHoldingRegisters(s.Addr, 1)
		case "discrete":
			res, err = client.ReadDiscreteInputs(s.Addr, 1)
		case "input":
			res, err = client.ReadInputRegisters(s.Addr, 1)
		case "coil":
			res, err = client.ReadCoils(s.Addr, 1)
		}

		if err != nil {
			if err.Error() == "serial: timeout" {
				log.Printf("id=[%3d] timeout", handler.SlaveId)
				continue
			}
			log.Printf("id=[%3d]: got response [%v] from addr %d", handler.SlaveId, err, s.Addr)
			continue
		}
		log.Printf("id=[%3d] got response %v from addr %d", handler.SlaveId, util.BytesToStringArray(res, 16), s.Addr)
	}
	return nil
}
