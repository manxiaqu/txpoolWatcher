package main

import (
	"context"
	"github.com/ethereum/go-ethereum/rpc"
)

const defaultRpcUrl = "http://127.0.0.1:8545"

type Client struct {
	client *rpc.Client
}

func DialDefault() (*Client, error) {
	return Dial(defaultRpcUrl)
}

func Dial(url string) (*Client, error) {
	c, err := rpc.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return NewClient(c), nil
}

func NewClient(c *rpc.Client) *Client {
	return &Client{c}
}

func (c *Client) TxPoolContent() (map[string]map[string]map[string]interface{}, error) {
	var result map[string]map[string]map[string]interface{}
	err := c.client.Call(&result, "txpool_content", nil)
	return result, err
}

func (c *Client) TxPoolInspect() (map[string]map[string]map[string]string, error) {
	var result map[string]map[string]map[string]string
	err := c.client.Call(&result, "txpool_inspect", "")
	return result, err
}

func (c *Client) TxPoolStatus() (map[string]string, error) {
	var result map[string]string
	err := c.client.Call(&result, "txpool_status", nil)
	return result, err
}
