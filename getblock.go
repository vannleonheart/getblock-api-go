package getblock

import (
	"errors"
	"github.com/vannleonheart/goutil"
)

func New(config *Config) *Client {
	return &Client{Config: config}
}

func (c *Client) WithRpcUrl(url string) *Client {
	c.rpcUrl = url

	return c
}

func (c *Client) SendRequest(method string, params []string, result interface{}) error {
	rpcUrl := c.getRpcUrl()
	if len(rpcUrl) <= 0 {
		return errors.New("rpc url is empty")
	}

	requestData := Request{
		Id:      Id,
		JsonRPC: JsonRPCVersion,
		Method:  method,
		Params:  params,
	}

	requestHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	_, err := goutil.SendHttpPost(rpcUrl, &requestData, &requestHeaders, result, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetTransactionByHash(hash string) (*Transaction, error) {
	var tr TransactionResponse

	if err := c.SendRequest(MethodEthGetTransactionByHash, []string{hash}, &tr); err != nil {
		return nil, err
	}

	return &tr.Result, nil
}

func (c *Client) getRpcUrl() string {
	if len(c.rpcUrl) > 0 {
		return c.rpcUrl
	}

	if c.Config != nil && len(c.Config.RpcUrl) > 0 {
		return c.Config.RpcUrl
	}

	return ""
}
