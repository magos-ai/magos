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
			if err := <-done; err != nil {
				t.Fatalf("after %v: %v, want exit 0", tc.sig, err)
			}
			rest, _ := io.ReadAll(out)
			if !strings.HasSuffix(string(rest), tc.role+" stopped.\n") {
				t.Errorf("output after the first line %q, want it to end with %q", rest, tc.role+" stopped.\n")
			}
		})
	}
}

func TestExitCodes(t *testing.T) {
	for _, tc := range []struct {
		args []string
		want int
	}{
		{nil, 2},
		{[]string{"nope"}, 2},
		{[]string{"serve", "extra"}, 2},
		{[]string{"help"}, 0},
	} {
		if got := run(tc.args, io.Discard, io.Discard); got != tc.want {
			t.Errorf("magos %q: exit %d, want %d", tc.args, got, tc.want)
		}
	}
}
