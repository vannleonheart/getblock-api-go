### GetBlock API

#### Installation
```go
go get -u github.com/vannleonheart/getblock-api-go
```

#### Config
```go
getblockConfig := getblock.Config {
    RpcUrl: "{your_chain_rpc_url}",
}
```

#### Create Client
```go
getblockClient := getblock.New(&getblockConfig)
```

#### Set RPC URL Manually
```go
rpcUrl := "{your_chain_rpc_url}"

getblockClient = getblockClient.WithRpcUrl(rpcUrl)
```
#### Get Transaction Detail By Hash
```go
hash := "{your_transaction_hash"

result, err := getblockClient.GetTransactionByHash(hash)

if err != nil {
    // handle error
}

fmt.Println(result)
```