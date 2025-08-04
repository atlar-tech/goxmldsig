package dsig

import (
	"time"
)

type Clock interface {
	Now() time.Time
}
