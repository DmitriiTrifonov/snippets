package application

import (
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Container is a DI container
type Container struct {
	logger *zap.Logger
}

// NewContainer inits a new Container
func NewContainer() *Container {
	return &Container{}
}

// GetLogger returns or inits zap.Logger
func (c *Container) GetLogger() (*zap.Logger, error) {
	if c.logger != nil {
		return c.logger, nil
	}
	zap.NewAtomicLevel()
	// This sets a TS to be more human-readable
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := config.Build()
	if err != nil {
		return nil, errors.Wrap(err, "cannot init logger")
	}
	c.logger = logger
	return c.logger, nil
}
