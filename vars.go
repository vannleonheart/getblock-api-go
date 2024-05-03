package getblock

type Client struct {
	Config *Config
	rpcUrl string
}

type Config struct {
	RpcUrl string `json:"rpc_url"`
}

type Request struct {
	Id      string   `json:"id"`
	JsonRPC string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

type Response struct {
	Id      string `json:"id"`
	JsonRPC string `json:"jsonrpc"`
}

type TransactionResponse struct {
	Response
	Result Transaction `json:"result"`
}

type Transaction struct {
	BlockHash            string   `json:"blockHash"`
	BlockNumber          string   `json:"blockNumber"`
	From                 string   `json:"from"`
	To                   string   `json:"to"`
	Gas                  string   `json:"gas"`
	GasPrice             string   `json:"gasPrice"`
	MaxFeePerGas         string   `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
	Hash                 string   `json:"hash"`
	Input                string   `json:"input"`
	Nonce                string   `json:"nonce"`
	TransactionIndex     string   `json:"transactionIndex"`
	Value                string   `json:"value"`
	Type                 string   `json:"type"`
	ChainId              string   `json:"chainId"`
	V                    string   `json:"v"`
	R                    string   `json:"r"`
	S                    string   `json:"s"`
	YParity              string   `json:"yParity"`
	AccessList           []string `json:"accessList"`
}
