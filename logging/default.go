package logging

import (
	"io"
	"log/slog"
)

func SetDefault(component string, level slog.Level, additionalWriters ...io.Writer) {
	writers := make([]io.Writer, 0, 1+len(additionalWriters))
	writers = append(writers, GetConsoleWriter())
	writers = append(writers, additionalWriters...)
	slog.SetDefault(NewMultiLogger(level, writers...).With("component", component))
}

func init() {
	// Set default logger to "titanic" with Info level
	// This may be overridden by the application
	SetDefault("titanic", slog.LevelInfo)
}
