package errorlibs

import (
	"errors"
)

type ErrorCause int

const (
	CAUSE_UNDEFINED ErrorCause = iota
	CAUSE_INVALID_ARGUMENT
	CAUSE_NOT_FOUND
	CAUSE_INTERNAL
)

type LogLevel int

const (
	LOG_LEVEL_UNDEFINED LogLevel = iota
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
)

type Err interface {
	Error() string
	Cause() ErrorCause
	LogLevel() LogLevel
}

type errImpl struct {
	err      error
	cause    ErrorCause
	logLevel LogLevel
}

func (ei *errImpl) Error() string {
	if ei == nil || ei.err == nil {
		return ""
	}
	return ei.err.Error()
}

func (ei *errImpl) Cause() ErrorCause {
	if ei == nil || ei.err == nil {
		return CAUSE_UNDEFINED
	}
	return ei.cause
}

func (ei *errImpl) LogLevel() LogLevel {
	if ei == nil || ei.err == nil {
		return LOG_LEVEL_UNDEFINED
	}
	return ei.logLevel
}

func NewErr(err error, cause ErrorCause, logLevel LogLevel) Err {
	if err == nil {
		return nil
	}
	return &errImpl{
		err:      err,
		cause:    cause,
		logLevel: logLevel,
	}
}

func NewErrFromMsg(msg string, cause ErrorCause, logLevel LogLevel) Err {
	return NewErr(errors.New(msg), cause, logLevel)
}
