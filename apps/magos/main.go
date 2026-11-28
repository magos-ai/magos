// Command magos is the Go program of Magos. Each command runs one role: the
// core or the Servo (docs/adr/0001).
package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/magos-ai/magos/apps/magos/internal/core"
	"github.com/magos-ai/magos/apps/magos/internal/servo"
)

const usage = `Usage: magos <command>

Commands:
  serve  Run the core.
  servo  Run the Servo.
  help   Show this help.
`

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

// run runs the command in args and returns the exit code.
func run(args []string, stdout, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprint(stderr, usage)
		return 2
	}
	var name string
	var role func(context.Context)
	switch args[0] {
	case "serve":
		name, role = "The core", core.Run
	case "servo":
		name, role = "The Servo", servo.Run
	case "help", "-h", "--help":
		fmt.Fprint(stdout, usage)
		return 0
	default:
		fmt.Fprintf(stderr, "Unknown command %q.\nRun \"magos help\" for the commands.\n", args[0])
		return 2
	}

	// Ctrl-C stops the role, and so does the SIGTERM that launchd sends.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	fmt.Fprintf(stderr, "%s runs.\nStop it with Ctrl-C.\n", name)
	role(ctx)
	fmt.Fprintf(stderr, "%s stopped.\n", name)
	return 0
}
