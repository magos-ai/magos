package main

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
	"time"
)

// TestMain runs main when a test starts this binary as the program magos.
func TestMain(m *testing.M) {
	if os.Getenv("MAGOS_RUN_MAIN") == "1" {
		main()
		return
	}
	os.Exit(m.Run())
}

// A role runs until Ctrl-C or the SIGTERM of launchd, and then exits with 0.
// The role and the signal do not depend on each other in run, so each one
// is in one case.
func TestRoleRunsUntilSignal(t *testing.T) {
	for _, tc := range []struct {
		command, role string
		sig           os.Signal
	}{
		{"serve", "The core", os.Interrupt},
		{"servo", "The Servo", syscall.SIGTERM},
	} {
		t.Run(tc.command+" "+tc.sig.String(), func(t *testing.T) {
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			defer r.Close()
			cmd := exec.Command(os.Args[0], tc.command)
			cmd.Env = append(os.Environ(), "MAGOS_RUN_MAIN=1")
			cmd.Stderr = w
			if err := cmd.Start(); err != nil {
				t.Fatal(err)
			}
			w.Close()
			done := make(chan error, 1)
			go func() { done <- cmd.Wait() }()
			defer cmd.Process.Kill()

			r.SetReadDeadline(time.Now().Add(5 * time.Second))
			out := bufio.NewReader(r)
			if line, err := out.ReadString('\n'); line != tc.role+" runs.\n" {
				t.Fatalf("first line %q (%v), want %q", line, err, tc.role+" runs.\n")
			}
			select {
			case err := <-done:
				t.Fatalf("magos %s stopped before the signal: %v", tc.command, err)
			case <-time.After(100 * time.Millisecond):
			}
			if err := cmd.Process.Signal(tc.sig); err != nil {
				t.Fatal(err)
			}
			select {
			case err := <-done:
				if err != nil {
					t.Fatalf("after %v: %v, want exit 0", tc.sig, err)
				}
			case <-time.After(5 * time.Second):
				t.Fatalf("magos %s still runs 5 s after %v", tc.command, tc.sig)
			}
			// The deadline of the first line must not cut this read. A new
			// one still ends it when another process holds the pipe open.
			r.SetReadDeadline(time.Now().Add(5 * time.Second))
			rest, err := io.ReadAll(out)
			if want := "Stop it with Ctrl-C.\n" + tc.role + " stopped.\n"; err != nil || string(rest) != want {
				t.Errorf("output after the first line %q (%v), want %q", rest, err, want)
			}
		})
	}
}

// A call that runs no role exits at once. With no command, magos prints the
// usage. Any other wrong call says what is wrong, and then the next step.
func TestCallsWithoutARole(t *testing.T) {
	const wantUsage = "Usage: magos <command>\n\nCommands:\n" +
		"  serve  Run the core.\n  servo  Run the Servo.\n  help   Show this help.\n"
	for _, tc := range []struct {
		args           []string
		want           int
		stdout, stderr string
	}{
		{nil, 2, "", wantUsage},
		{[]string{"nope"}, 2, "", "Unknown command \"nope\".\nRun \"magos help\" for the commands.\n"},
		{[]string{"nope", "extra"}, 2, "", "Unknown command \"nope\".\nRun \"magos help\" for the commands.\n"},
		{[]string{"serve", "extra"}, 2, "", "\"magos serve\" takes no arguments.\nRun \"magos serve\".\n"},
		{[]string{"help"}, 0, wantUsage, ""},
		{[]string{"-h"}, 0, wantUsage, ""},
		{[]string{"--help"}, 0, wantUsage, ""},
	} {
		var stdout, stderr strings.Builder
		got := run(tc.args, &stdout, &stderr)
		if got != tc.want || stdout.String() != tc.stdout || stderr.String() != tc.stderr {
			t.Errorf("magos %q: exit %d, stdout %q, stderr %q; want %d, %q, %q",
				tc.args, got, stdout.String(), stderr.String(), tc.want, tc.stdout, tc.stderr)
		}
	}
}
