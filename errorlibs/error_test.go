package errorlibs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrImpl_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		ei       *errImpl
		expected string
	}{
		{
			name: "Success",
			ei: &errImpl{
				err: errors.New("err"),
			},
			expected: "err",
		},
		{
			name:     "Success(*errImpl is nil)",
			ei:       nil,
			expected: "",
		},
		{
			name:     "Success(err is nil)",
			ei:       &errImpl{},
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, tt.ei.Error())
		})
	}
}

func TestErrImpl_Cause(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		ei       *errImpl
		expected ErrorCause
	}{
		{
			name: "Success",
			ei: &errImpl{
				err:   errors.New("err"),
				cause: CAUSE_INTERNAL,
			},
			expected: CAUSE_INTERNAL,
		},
		{
			name:     "Success(*errImpl is nil)",
			ei:       nil,
			expected: CAUSE_UNDEFINED,
		},
		{
			name:     "Success(err is nil)",
			ei:       &errImpl{},
			expected: CAUSE_UNDEFINED,
		},
		{
			name: "Error(cause is not defined)",
			ei: &errImpl{
				err: errors.New("err"),
			},
			expected: CAUSE_UNDEFINED,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, tt.ei.Cause())
		})
	}
}

func TestErrImpl_LogLevel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		ei       *errImpl
		expected LogLevel
	}{
		{
			name: "Success",
			ei: &errImpl{
				err:      errors.New("err"),
				logLevel: LOG_LEVEL_INFO,
			},
			expected: LOG_LEVEL_INFO,
		},
		{
			name:     "Success(*errImpl is nil)",
			ei:       nil,
			expected: LOG_LEVEL_UNDEFINED,
		},
		{
			name:     "Success(err is nil)",
			ei:       &errImpl{},
			expected: LOG_LEVEL_UNDEFINED,
		},
		{
			name: "Error(logLevel is not defined)",
			ei: &errImpl{
				err: errors.New("err"),
			},
			expected: LOG_LEVEL_UNDEFINED,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, tt.ei.LogLevel())
		})
	}
}

func TestNewErr(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		err      error
		cause    ErrorCause
		logLevel LogLevel
		want     Err
	}{
		{
			name:     "normal case",
			err:      errors.New("error"),
			cause:    CAUSE_INTERNAL,
			logLevel: LOG_LEVEL_ERROR,
			want: &errImpl{
				err:      errors.New("error"),
				cause:    CAUSE_INTERNAL,
				logLevel: LOG_LEVEL_ERROR,
			},
		},
		{
			name:     "err is nil",
			err:      nil,
			cause:    CAUSE_INTERNAL,
			logLevel: LOG_LEVEL_ERROR,
			want:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, NewErr(tt.err, tt.cause, tt.logLevel))
		})
	}
}

func TestNewErrFromMsg(t *testing.T) {
	t.Parallel()
	assert.Equal(
		t,
		&errImpl{
			err:      errors.New("error"),
			cause:    CAUSE_INTERNAL,
			logLevel: LOG_LEVEL_ERROR,
		},
		NewErrFromMsg("error", CAUSE_INTERNAL, LOG_LEVEL_ERROR),
	)
}
