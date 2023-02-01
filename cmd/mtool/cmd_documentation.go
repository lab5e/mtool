package main

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/lab5e/mtool"
	"github.com/yuin/goldmark"
)

type documentationCmd struct {
	Format string `long:"format" default:"markdown" description:"documentation format" choice:"markdown" choice:"html"`
}

func (d *documentationCmd) Execute([]string) error {
	switch d.Format {
	case "markdown":
		fmt.Println(string(mtool.ReadmeBytes))
		return nil

	case "html":
		var buf bytes.Buffer
		err := goldmark.Convert(mtool.ReadmeBytes, &buf)
		if err != nil {
			return err
		}
		fmt.Println(buf.String())
		return nil

	default:
		return errors.New("unknown format")
	}
}
