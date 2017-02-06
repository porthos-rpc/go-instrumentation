package instrumentation

import (
	porthos "github.com/porthos-rpc/porthos-go/server"
)

type TxRequest struct {
	Transaction
	porthos.Request
}
