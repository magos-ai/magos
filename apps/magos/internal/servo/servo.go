// Package servo is the Servo of Magos, the role of `magos servo`. It will run
// the harness sessions and work with files (docs/adr/0001). It talks to the
// core only through the control protocol, so it imports no package of the
// core.
package servo

import "context"

// Run runs the Servo until ctx is done. The Servo does no work yet.
func Run(ctx context.Context) {
	<-ctx.Done()
}
