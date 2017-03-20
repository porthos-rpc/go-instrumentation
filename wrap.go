package instrumentation

import (
	"github.com/porthos-rpc/porthos-go"
)

func WrapHandler(app Application, method string, handler porthos.MethodHandler) (string, porthos.MethodHandler) {
	return method, func(req porthos.Request, res porthos.Response) {
		tx := app.StartTransaction(method)
		defer tx.End()

		handler(&TxRequest{tx, req}, &TxResponse{tx, res})
	}
}
