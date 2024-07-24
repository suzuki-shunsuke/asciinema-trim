package controller

import (
	"fmt"
)

type Controller struct {
	Writer Writer
}

type Param struct {
	CastFile string
}

func New(_ *Param) (*Controller, error) {
	return &Controller{
		Writer: &simpleWriter{},
	}, nil
}

type Writer interface {
	Write(s string)
}

type simpleWriter struct{}

func (w *simpleWriter) Write(s string) {
	fmt.Println(s) //nolint:forbidigo
}
