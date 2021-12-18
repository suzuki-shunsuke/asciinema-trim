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

func build(write func(string), file io.Reader) error {
	var delta float64
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	write(scanner.Text())
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "[") {
			idx := strings.Index(txt, ",")
			if idx == -1 {
				return errors.New(`"," isn't founc`)
			}
			t, err := strconv.ParseFloat(txt[1:idx], 64) //nolint:gomnd
			if err != nil {
				return fmt.Errorf("time must be float64: %w", err)
			}
			if delta != 0 {
				txt = "[" + strconv.FormatFloat(t-delta, 'f', -1, 64) + txt[idx:] //nolint:gomnd
			}
			write(txt)
			continue
		}
		a, err := strconv.ParseFloat(txt, 64) //nolint:gomnd
		if err != nil {
			return fmt.Errorf("trimmed time must be float64: %w", err)
		}
		if a <= 0 {
			return errors.New("trimmed time must be greater than 0")
		}
		delta += a
		continue
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("non EOF error: %w", err)
	}
	return nil
}

func (ctrl *Controller) build(ctx context.Context, param *Param, file io.Reader) error { //nolint:unparam
	return build(ctrl.write, file)
}

func (ctrl *Controller) write(s string) {
	ctrl.Writer.Write(s)
}
