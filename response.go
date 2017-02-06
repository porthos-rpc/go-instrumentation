package instrumentation

import (
	porthos "github.com/porthos-rpc/porthos-go/server"
)

type TxResponse struct {
	Transaction
	porthos.Response
}

func (r *TxResponse) JSON(statusCode int16, body interface{}) {
	defer r.StartSegment("Response time").End()
	r.Response.JSON(statusCode, body)
}

func (r *TxResponse) Raw(statusCode int16, contentType string, body []byte) {
	defer r.StartSegment("Response time").End()
	r.Response.Raw(statusCode, contentType, body)
}

func (r *TxResponse) Empty(statusCode int16) {
	defer r.StartSegment("Response time").End()
	r.Response.Empty(statusCode)
}
