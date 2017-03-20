package instrumentation

import (
	"github.com/porthos-rpc/porthos-go"
)

type TxRequest struct {
	Transaction
	porthos.Request
}
