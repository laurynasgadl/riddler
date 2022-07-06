package riddler

import (
	"math/rand"
	"time"

	"github.com/laurynasgadl/riddler/strings"
)

var R *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()),
)

// Generate returns a randomly generated string using the provided options
func String(opts ...strings.Option) string {
	return strings.NewStrings(R).Generate(opts...)
}
