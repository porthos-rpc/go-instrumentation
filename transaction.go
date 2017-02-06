package instrumentation

import (
	"github.com/newrelic/go-agent"
)

type Transaction interface {
	StartSegment(name string) Segment
	End() error
	Ignore() error
	NoticeError(err error) error
	AddAttribute(key string, value interface{}) error
}

type transaction struct {
	newrelic.Transaction
}

func (t *transaction) StartSegment(name string) Segment {
	return newrelic.StartSegment(t, name)
}
