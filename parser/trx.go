package parser

import (
	"encoding/json"
	"fmt"
	"parser/types"
)

// CheckTrxs fetches transactions included to a block
// and gathers transactions with needed contracts addresses
func (p *Parser) checkTrxs() (map[string][]types.Transaction, error) {

	if len(p.subscriptions) == 0 {
		return map[string][]types.Transaction{}, nil
	}

	reqBody := reqRPC{"2.0", "eth_getBlockByNumber", []interface{}{p.lastBlock, true}, 1} // request id may be iterated if needed

	respBytes, err := makeRequest(p.rpcURL, reqBody)
	if err != nil {
		return nil, err
	}

	var respStruct trxResp

	err = json.Unmarshal(respBytes, &respStruct)
	if err != nil {
		return nil, err
	}

	fmt.Printf("transactions count: %v\n", len(respStruct.Result.Transactions))

	filtered := make(map[string][]types.Transaction, len(p.subscriptions))

	for _, trx := range respStruct.Result.Transactions {

		p.mu.RLock()
		if _, subscribed := p.subscriptions[trx.To]; subscribed {

			if _, added := filtered[trx.To]; !added {
				filtered[trx.To] = make([]types.Transaction, 0, 10)
			}

			filtered[trx.To] = append(filtered[trx.To], trx)
		}
		p.mu.RUnlock()
	}

	return filtered, nil
}
