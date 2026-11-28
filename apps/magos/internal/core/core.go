// Package core is the core of Magos, the role of `magos serve`. It will keep
// the workspaces, the projects, the grants, the threads and the runs
// (docs/adr/0001). It talks to the Servo only through the control protocol,
// so it imports no package of the Servo.
package core

import "context"

// Run runs the core until ctx is done. The core does no work yet.
func Run(ctx context.Context) {
	<-ctx.Done()
}
