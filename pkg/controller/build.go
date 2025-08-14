package controller

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func (ctrl *Controller) Build(ctx context.Context, param *Param) error {
	file, err := os.Open(param.CastFile)
	if err != nil {
		return fmt.Errorf("open a file: %w", err)
	}
	defer file.Close()
	return ctrl.build(ctx, param, file)
}

type State struct {
	Prev       float64
	ActualPrev float64
	Speed      float64
	Delta      float64
}

var (
	errSpeedMustBeGreaterThanZero = errors.New("speed must be greater than zero")
	errCommaIsntFound             = errors.New(`"," isn't found`)
)

func parseLine(txt string, state *State) (string, error) {
	if strings.HasPrefix(txt, "[") {
		idx := strings.Index(txt, ",")
		if idx == -1 {
			return "", errCommaIsntFound
		}
		t, err := strconv.ParseFloat(txt[1:idx], 64)
		if err != nil {
			return "", fmt.Errorf("time must be float64: %w", err)
		}
		a := state.Prev + (t-state.ActualPrev-state.Delta)/state.Speed
		if a <= state.Prev {
			return "", fmt.Errorf("trimmed time must be lower than interval: %w", err)
		}
		state.Prev = a
		state.ActualPrev = t
		state.Delta = 0
		return "[" + strconv.FormatFloat(a, 'f', -1, 64) + txt[idx:], nil
	}
	if strings.HasPrefix(txt, "*") {
		s, err := strconv.ParseFloat(txt[1:], 64)
		if err != nil {
			return "", fmt.Errorf("speed must be float64: %w", err)
		}
		if s <= 0 {
			return "", errSpeedMustBeGreaterThanZero
		}
		state.Speed = s
		return "", nil
	}
	a, err := strconv.ParseFloat(txt, 64)
	if err != nil {
		return "", fmt.Errorf("trimmed time must be float64: %w", err)
	}
	state.Delta = a
	return "", nil
}

func build(write func(string), file io.Reader) error {
	lineNum := 1
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	write(scanner.Text())
	state := &State{
		Speed: 1,
	}
	for scanner.Scan() {
		lineNum++
		txt := scanner.Text()
		if s, err := parseLine(txt, state); err != nil {
			return fmt.Errorf("parse a line line_number=%d: %w", lineNum, err)
		} else if s != "" {
			write(s)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("non EOF error: %w", err)
	}
	return nil
}

func (ctrl *Controller) build(_ context.Context, _ *Param, file io.Reader) error {
	return build(ctrl.write, file)
}

func (ctrl *Controller) write(s string) {
	ctrl.Writer.Write(s)
}
