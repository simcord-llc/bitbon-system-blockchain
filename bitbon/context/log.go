// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
//
// The Bitbon System Blockchain Node library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Bitbon System Blockchain Node library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Bitbon System Blockchain Node library. If not, see <http://www.gnu.org/licenses/>.

package context

import (
	"context"

	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

type loggerKeyType int

const loggerKey loggerKeyType = iota

// NewLoggerContext creates context from parent context passed as a first argument
// and sets logger instance in it as value.
// If logger is nil default logger root will be set.
func NewLoggerContext(ctx context.Context, logger log.Logger) context.Context {
	if logger == nil {
		logger = log.Root()
	}
	return context.WithValue(ctx, loggerKey, logger)
}

// NewLoggerFieldsContext do the same as NewLoggerContext does,
// but it creates
func NewLoggerFieldsContext(ctx context.Context, fields ...interface{}) context.Context {
	logger := log.New(fields...)
	return NewLoggerContext(ctx, logger)
}

// LoggerFromContext receives context as an argument and tries to get logger from it.
// If logger is not found in context method returns default logger root
func LoggerFromContext(ctx context.Context) log.Logger {
	if ctx == nil {
		return log.Root()
	}
	if ctxLogger, ok := ctx.Value(loggerKey).(log.Logger); ok {
		return ctxLogger
	}

	return log.Root()
}
