package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"time"
)

// Logger wraps logrus.Logger
type Logger struct {
	*logrus.Logger
}

// New creates a new preconfigured logrus.Logger
func New() Logger {
	logger := logrus.New()
	logger.Formatter = &prefixed.TextFormatter{
		FullTimestamp:    true,
		ForceColors:      true,
		DisableSorting:   false,
		QuoteEmptyFields: true,
		SpacePadding:     32,
	}

	return Logger{
		Logger: logger,
	}
}

// RequestEnd logs details of request end
func (logger *Logger) RequestEnd(startAt time.Time, act string, status *uint32, errMsg *string) {
	logger.WithFields(logrus.Fields{
		"action":        act,
		"status_code":   status,
		"error_message": errMsg,
		"created_at":    startAt.Unix(),
		"response_time": fmt.Sprintf("%.4f", time.Since(startAt).Seconds()),
	}).Info("request")
}
