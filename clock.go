// Modified by Atlar in 2025: replaced the clockwork-based Clock struct with a Clock interface.
// Original: github.com/russellhaering/goxmldsig v1.2.0 (Apache-2.0).

package dsig

import (
	"time"
)

type Clock interface {
	Now() time.Time
}
