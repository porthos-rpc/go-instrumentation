package instrumentation

import (
	"github.com/newrelic/go-agent"
	"github.com/porthos-rpc/porthos-go"
)

func WrapHandler(app newrelic.Application, method string, handler porthos.MethodHandler) (string, porthos.MethodHandler) {
	return method, func(req porthos.Request, res porthos.Response) {
		tx := app.StartTransaction(method, nil, nil)
		defer tx.End()

		handler(&TxRequest{tx, req}, &TxResponse{tx, res})
	}
}
