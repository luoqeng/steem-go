package network_broadcast

import (
	"github.com/luoqeng/steem-go/transports"
	"github.com/luoqeng/steem-go/types"
)

const apiID = "condenser_api."

//API plug-in structure
type API struct {
	caller transports.Caller
}

//NewAPI plug-in initialization
func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call(apiID+method, params, resp)
}

//BroadcastTransaction api request broadcast_transaction
func (api *API) BroadcastTransaction(tx *types.Transaction) error {
	return api.call("broadcast_transaction", []interface{}{tx}, nil)
}

//BroadcastTransactionSynchronous api request broadcast_transaction_synchronous
func (api *API) BroadcastTransactionSynchronous(tx *types.Transaction) (*BroadcastResponse, error) {
	var resp BroadcastResponse
	err := api.call("broadcast_transaction_synchronous", []interface{}{tx}, &resp)
	return &resp, err
}
