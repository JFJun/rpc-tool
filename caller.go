package rpc_tool

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/rpc-tool/http"
	"github.com/JFJun/rpc-tool/websocket"
	"io"
	"strings"
)

type Caller interface {
	Call(api uint8, method string, args []interface{}, reply interface{}) error
	SetCallback(api uint8, method string, callback func(raw json.RawMessage)) error
	Connect() error
}

type CallCloser interface {
	Caller
	io.Closer
}

type RpcTool struct {
	CC CallCloser
	id uint8
}

func NewRPcTool(url string) *RpcTool {
	var cc CallCloser
	var err error
	var id uint8
	if strings.HasPrefix(url, "http") || strings.HasPrefix(url, "https") {
		cc, err = http.NewHttpTransport(url)
		id = 0
	} else {
		cc, err = websocket.NewTransport(url)
		id = 1
	}
	if err != nil {
		panic(fmt.Sprintf("Init rpc tool error,Err=[%v]", err))
	}
	rt := new(RpcTool)
	rt.CC = cc
	rt.id = id
	return rt
}
func (rt *RpcTool) Close() error {
	return rt.CC.Close()
}

/*
调用rpc接口
*/
func (rt *RpcTool) Call(method string, args []interface{}, reply interface{}) error {
	err := rt.CC.Connect()
	if err != nil {
		return err
	}
	return rt.CC.Call(rt.id, method, args, reply)
}
