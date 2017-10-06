package instrumentation

import (
	"github.com/newrelic/go-agent"
	"github.com/porthos-rpc/porthos-go"
)

type TxRequest struct {
	newrelic.Transaction
	porthos.Request
}
