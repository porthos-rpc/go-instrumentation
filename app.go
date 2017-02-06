package instrumentation

import (
	"time"

	"github.com/newrelic/go-agent"
)

type Application interface {
	StartTransaction(name string) Transaction
	Shutdown(timeout time.Duration)
}

type application struct {
	newrelic.Application
}

func (a *application) StartTransaction(name string) Transaction {
	tx := a.Application.StartTransaction(name, nil, nil)

	return &transaction{
		tx,
	}
}

func NewApplication(name, key string) (Application, error) {
	config := newrelic.NewConfig(name, key)
	app, err := newrelic.NewApplication(config)

	if err != nil {
		return nil, err
	}

	return &application{
		app,
	}, nil
}
