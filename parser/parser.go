package parser

import (
	"fmt"
	"log"
	"sync"
	"time"

	"parser/types"
)

func New(r repoReadWriter, rpc string) Parser {
	return Parser{
		mu:             sync.RWMutex{},
		repo:           r,
		rpcURL:         rpc,
		addresses:      map[string]struct{}{},
		updateInterval: 2 * time.Second,
	}
}

// Start launches an infinite loop that checks for block number updates
// and calls transactions fetching and filtering,
// saves transactions of addresses we are subscribed to
func (p *Parser) Start() {

	ticker := time.NewTicker(p.updateInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			num, err := p.getBlockNum()
			if err != nil {
				log.Println(err)
				continue
			}

			if p.lastBlock != num {
				p.lastBlock = num

				// For demo purposes
				fmt.Printf("current block number %v\n", p.GetCurrentBlock())

				trxs, err := p.checkTrxs()
				if err != nil {
					log.Println(err)
					continue
				}

				err = p.writeTrxs(trxs)
				if err != nil {
					log.Println(err)
					continue
				}

			}

		}
	}
}

func (p *Parser) readTrxs(address string) []types.Transaction {
	trxs, err := p.repo.Read(address)
	if err != nil {
		log.Println(err)
		return []types.Transaction{}
	}

	return trxs
}

func (p *Parser) writeTrxs(trxs map[string][]types.Transaction) error {
	return p.repo.Write(trxs)
}
