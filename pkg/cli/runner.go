package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/suzuki-shunsuke/asciinema-trim/pkg/controller"
)

func showCLIHelp() {
	fmt.Fprint(os.Stderr, `asciinema-trim - Trim asciinema Record File

https://github.com/suzuki-shunsuke/asciinema-trim

Usage
  asciinema-trim --help  - Show this help
  asciinema-trim --version  - Show asciinema-trim version
  asciinema-trim <record file>  - Output trimmed record file
`)
}

type cliArgs struct {
	Help     bool
	Version  bool
	CastFile string
}

func parseArgs() *cliArgs {
	cArgs := &cliArgs{}
	flag.BoolVar(&cArgs.Help, "help", false, "show help message")
	flag.BoolVar(&cArgs.Version, "version", false, "show version")
	flag.Parse()
	return cArgs
}

type Runner struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	LDFlags *LDFlags
}

type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

func (runner *Runner) Run(ctx context.Context) error {
	cArgs := parseArgs()
	args := flag.Args()

	if cArgs.Help {
		showCLIHelp()
		return nil
	}

	if cArgs.Version {
		fmt.Fprintln(os.Stderr, runner.LDFlags.Version+" ("+runner.LDFlags.Commit+")")
		return nil
	}

	if len(args) == 0 {
		return errors.New("record file is needed")
	}
	param := &controller.Param{
		CastFile: args[0],
	}
	ctrl, err := controller.New(param)
	if err != nil {
		return fmt.Errorf("initialize a controller: %w", err)
	}
	return ctrl.Build(ctx, param) //nolint:wrapcheck
}
