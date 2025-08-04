package main

import (
	"context"
	"errors"

	"github.com/ybbus/jsonrpc/v3"
)

type Message struct {
	Account   string   `json:"account"`
	Recipient []string `json:"recipient"`
	Message   string   `json:"message"`
}

var rpcClient jsonrpc.RPCClient

func initRpcClient() {
	rpcClient = jsonrpc.NewClient(config.SignalCliEndpoint)
}

func sendMessage(recipient string, message string) error {
	params := Message{Account: config.Account, Recipient: []string{recipient}, Message: message}

	response, err := rpcClient.Call(context.Background(), "send", &params)

	if err != nil {
		return err
	}

	if response.Error != nil {
		return errors.New(response.Error.Message)
	}

	return nil
}
