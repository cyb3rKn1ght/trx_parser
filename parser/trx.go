package parser

import (
	"encoding/json"
	"fmt"
	"parser/types"
)

// CheckTrxs fetches transactions included to a block
// and gathers transactions with needed contracts addresses
func (p *Parser) checkTrxs() (map[string][]types.Transaction, error) {

	// commented for demo purposes
	// if len(p.addresses) == 0 {
	// 	return map[string][]types.Transaction{}, nil
	// }

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

	// Uncomment to start saving address transactions info
	// firstAddr := respStruct.Result.Transactions[0].To
	// p.Subscribe(firstAddr)
	// fmt.Printf("subscribed %v\n", firstAddr)

	filtered := make(map[string][]types.Transaction, len(p.addresses))

	for _, trx := range respStruct.Result.Transactions {

		p.mu.RLock()
		if _, subscribed := p.addresses[trx.To]; subscribed {

			if _, added := filtered[trx.To]; !added {
				filtered[trx.To] = make([]types.Transaction, 0, 10)
			}

			filtered[trx.To] = append(filtered[trx.To], trx)
		}
		p.mu.RUnlock()
	}

	return filtered, nil
}
