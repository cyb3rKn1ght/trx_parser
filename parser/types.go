package parser

import (
	"sync"
	"time"

	"parser/types"
)

type Parser struct {
	mu             sync.RWMutex
	repo           repoReadWriter
	rpcURL         string // TODO make it a list of RPCs to handle request fails
	lastBlock      string
	subsPath       string
	updateInterval time.Duration
	subscriptions  map[string]struct{}
}

type repoReadWriter interface {
	ReadTrxs(path string) ([]types.Transaction, error)
	WriteTrxs(trxs map[string][]types.Transaction) error
	SaveSubs(path string, data map[string]struct{}) error
	LoadSubs(path string) (map[string]struct{}, error)
}

type result struct {
	Transactions []types.Transaction
}

type respBlockNum struct {
	Version string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      uint   `json:"id"`
}

type trxResp struct {
	Version string `json:"jsonrpc"`
	Result  result `json:"result"`
	ID      uint   `json:"id"`
}

type reqRPC struct {
	Version string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      uint          `json:"id"`
}
