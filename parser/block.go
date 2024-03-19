package parser

import (
	"encoding/json"
	"log"
)

func (p *Parser) getBlockNum() (string, error) {
	reqBody := reqRPC{"2.0", "eth_blockNumber", []interface{}{}, 1} // request id may be iterated if needed

	respBytes, err := makeRequest(p.rpcURL, reqBody)
	if err != nil {
		return "", err
	}

	var blockNum respBlockNum

	err = json.Unmarshal(respBytes, &blockNum)
	if err != nil {
		log.Println(err)
	}

	return blockNum.Result, nil
}
