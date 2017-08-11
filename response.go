package instrumentation

import (
	"github.com/porthos-rpc/porthos-go"
)

type TxResponse struct {
	Transaction
	porthos.Response
}

func (r *TxResponse) JSON(statusCode int32, body interface{}) {
	defer r.StartSegment("Response time").End()
	r.Response.JSON(statusCode, body)
}

func (r *TxResponse) Raw(statusCode int32, contentType string, body []byte) {
	defer r.StartSegment("Response time").End()
	r.Response.Raw(statusCode, contentType, body)
}

func (r *TxResponse) Empty(statusCode int32) {
	defer r.StartSegment("Response time").End()
	r.Response.Empty(statusCode)
}
