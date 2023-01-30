package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type dataPoint struct {
	Time      time.Time `json:"time"`
	ValueType string    `json:"valueType"`
	SlaveID   byte      `json:"slaveId"`
	Addr      uint16    `json:"addr"`
	Data      string    `json:"data"`
}

func (d *dataPoint) JSON() string {
	jsonData, err := json.Marshal(d)
	if err != nil {
		log.Fatalf("error marshalling json: %v", err)
	}
	return string(jsonData)
}

func (d *dataPoint) String() string {
	return fmt.Sprintf("%s valueType=%s slaveID=%d addr=%d data={%s}", d.Time.Format(time.RFC3339), d.ValueType, d.SlaveID, d.Addr, d.Data)
}
