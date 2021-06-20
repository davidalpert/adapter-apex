// Package apex provides a Logur adapter for TEMPLATE.
package apex

import (
	"context"

	"github.com/apex/log"
)

// Logger is a Logur adapter for apex/log.
type Logger struct {
	entry *log.Entry
}

// New returns a new Logur logger.
func New() *Logger {
	return NewFromFields(log.Fields{})
}

// NewFromFields returns a new Logur logger from a set of apex log.Fields.
func NewFromFields(fields log.Fields) *Logger {
	return NewFromEntry(log.WithFields(fields))
}

// NewFromEntry returns a new Logur logger from an apex log.Entry.
// If entry is nil, a default instance is created.
func NewFromEntry(entry *log.Entry) *Logger {
	if entry == nil {
		entry = log.WithFields(log.Fields{})
	}

	return &Logger{
		entry: entry,
	}
}

// entryWithFields merges fields into the entry and returns a log.Entry with the given fields.
func (l *Logger) entryWithFields(fields []map[string]interface{}) *log.Entry {
	if len(fields) == 0 {
		return l.entry
	}

	entry := l.entry

	for k, v := range fields[0] {
		entry = entry.WithField(k, v)
	}

	return entry
}

// Trace implements the Logur Logger interface.
func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {
	// apex log library doesn't include Trace level
	// (it uses the concept of Trace as a verb, using go's defer to bookend a transaction)
	// so we'll write these at debug
	l.entryWithFields(fields).Debug(msg)
}

// Debug implements the Logur Logger interface.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	l.entryWithFields(fields).Debug(msg)
}

// Info implements the Logur Logger interface.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	l.entryWithFields(fields).Info(msg)
}

// Warn implements the Logur Logger interface.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	l.entryWithFields(fields).Warn(msg)
}

// Error implements the Logur Logger interface.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	l.entryWithFields(fields).Error(msg)
}

func (l *Logger) TraceContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Trace(msg, fields...)
}

func (l *Logger) DebugContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Debug(msg, fields...)
}

func (l *Logger) InfoContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Info(msg, fields...)
}

func (l *Logger) WarnContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Warn(msg, fields...)
}

func (l *Logger) ErrorContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Error(msg, fields...)
}

// LevelEnabled implements the Logur LevelEnabled interface.
func (l *Logger) LevelEnabled(level log.Level) bool {
	if globalLogger, ok := log.Log.(*log.Logger); ok {
		return globalLogger.Level >= level
	}

	return true // if we can't determine the log level, fail open
}
