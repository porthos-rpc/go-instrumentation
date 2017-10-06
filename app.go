package instrumentation

import (
	"github.com/newrelic/go-agent"
)

func NewApplication(name, key string) (newrelic.Application, error) {
	config := newrelic.NewConfig(name, key)
	return newrelic.NewApplication(config)
}
