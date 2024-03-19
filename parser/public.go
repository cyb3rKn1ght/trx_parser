package parser

import (
	"log"
	"parser/types"
	"regexp"
	"strconv"
)

var reg = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

func (p *Parser) GetCurrentBlock() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	parsedBlockNum, err := strconv.ParseInt(p.lastBlock, 0, 64)
	if err != nil {
		log.Println(err)
		return -1
	}

	return int(parsedBlockNum)
}

func (p *Parser) Subscribe(address string) bool {

	if !reg.MatchString(address) {
		return false
	}

	p.mu.Lock()
	p.subscriptions[address] = struct{}{}
	p.mu.Unlock()

	err := p.saveSubs()
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (p *Parser) GetTransactions(address string) []types.Transaction {
	return p.readTrxs(address)
}
