package lsp

type Request struct {
    RPC string `json:"jsonrpc"`
    ID int `json:"id"`
    Method string `json:"method"`

    
}
