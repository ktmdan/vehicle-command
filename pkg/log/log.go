// Package log exposes the library's global log-level control to external
// consumers. It is a thin wrapper over internal/log, which is otherwise
// unimportable outside this module.
package log

import ilog "github.com/teslamotors/vehicle-command/internal/log"

// Level mirrors the internal log level. Higher levels log more.
type Level = ilog.Level

const (
	LevelNone    = ilog.LevelNone    // Disables logging.
	LevelError   = ilog.LevelError   // Unexpected anomalies.
	LevelWarning = ilog.LevelWarning // Occasional expected anomalies.
	LevelInfo    = ilog.LevelInfo    // Major events.
	LevelDebug   = ilog.LevelDebug   // Detailed IO.
)

// SetLevel sets the global log verbosity for the library.
func SetLevel(level Level) {
	ilog.SetLevel(level)
}
