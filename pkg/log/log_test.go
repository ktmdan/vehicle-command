package log_test

import (
	"testing"

	"github.com/teslamotors/vehicle-command/pkg/log"
)

func TestSetLevelAcceptsConstants(t *testing.T) {
	// Exercises every exported constant and the setter. The internal logger has
	// no getter, so this is a smoke test that the wrapper compiles and runs.
	for _, lvl := range []log.Level{
		log.LevelNone,
		log.LevelError,
		log.LevelWarning,
		log.LevelInfo,
		log.LevelDebug,
	} {
		log.SetLevel(lvl)
	}
}
