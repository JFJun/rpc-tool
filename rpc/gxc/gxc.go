package gxc

import (
	"encoding/json"
	rpc_tool "github.com/JFJun/rpc-tool"
	"github.com/pkg/errors"
	"reflect"

	"gxclient-go/types"
)

type GXCRpc struct {
	client *rpc_tool.RpcTool
}

func NewGXCRpc(url string) *GXCRpc {
	gr := new(GXCRpc)
	gr.client = rpc_tool.NewRPcTool(url)
	return gr
}

var EmptyParams = []interface{}{}

func (gr *GXCRpc) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	var resp DynamicGlobalProperties
	err := gr.client.Call("get_dynamic_global_properties", EmptyParams, &resp)
	return &resp, err
}

// Get block by block height
func (gr *GXCRpc) GetBlock(blockNum uint32) (*Block, error) {
	var resp Block
	err := gr.client.Call("get_block", []interface{}{blockNum}, &resp)
	return &resp, err
}

// GET ChainId of entry point
func (gr *GXCRpc) GetChainId() (string, error) {
	var resp string
	err := gr.client.Call("get_chain_id", EmptyParams, &resp)
	return resp, err
}

//get_account_by_name
func (gr *GXCRpc) GetAccount(account string) (*types.Account, error) {
	var resp *types.Account
	if err := gr.client.Call("get_account_by_name", []interface{}{account}, &resp); err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.Errorf("account %s not exist", account)
	}
	return resp, nil
}

func (gr *GXCRpc) GetTransactionByTxid(txid string) (*types.Transaction, error) {
	var resp *types.Transaction
	if err := gr.client.Call("get_transaction_rows", []interface{}{txid}, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// LookupAssetSymbols get assets corresponding to the provided symbol or IDs
func (gr *GXCRpc) GetAsset(symbol string) (*Asset, error) {
	var resp []*Asset
	if err := gr.client.Call("lookup_asset_symbols", []interface{}{[]string{symbol}}, &resp); err != nil {
		return nil, err
	}
	if resp[0] == nil {
		return nil, errors.Errorf("assets %s not exist", symbol)
	}
	return resp[0], nil
}

// GetRequiredFee fetchs fee for operations
func (gr *GXCRpc) GetRequiredFee(ops []types.Operation, assetID string) ([]types.AssetAmount, error) {
	var resp []types.AssetAmount

	opsJSON := []interface{}{}
	for _, o := range ops {
		_, err := json.Marshal(o)
		if err != nil {
			return []types.AssetAmount{}, err
		}

		opArr := []interface{}{o.Type(), o}

		opsJSON = append(opsJSON, opArr)
	}
	if err := gr.client.Call("get_required_fees", []interface{}{opsJSON, assetID}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (gr *GXCRpc) BroadcastTransactionSynchronous(tx *types.Transaction) (*types.BroadcastResponse, error) {
	response := types.BroadcastResponse{}
	var err error
	if typeof(gr.client.CC) == "*http.HttpTransport" {
		txs := make([]*types.Transaction, 1)
		txs[0] = tx
		err = gr.client.Call("call", []interface{}{2, "broadcast_transaction_synchronous", txs}, &response)
	} else {
		err = gr.client.Call("broadcast_transaction_synchronous", []interface{}{tx}, &response)
	}
	if err != nil {
		return nil, err
	}
	return &response, err
}
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
