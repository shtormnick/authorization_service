package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
	once   = &sync.Once{}
)

// Get ...
func Get() *logrus.Logger {
	once.Do(func() {
		logger = logrus.New()
	})
	return logger
}
