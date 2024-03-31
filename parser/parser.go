package parser

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"parser/types"
)

func New(r repoReadWriter, rpc string) Parser {
	return Parser{
		mu:             sync.RWMutex{},
		repo:           r,
		rpcURL:         rpc,
		subsPath:       "subs",
		subscriptions:  map[string]struct{}{},
		updateInterval: 2 * time.Second,
	}
}

// Start launches an infinite loop that checks for block number updates
// and calls transactions fetching and filtering,
// saves transactions of addresses we are subscribed to
func (p *Parser) Start() {

	err := p.loadSubs()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}

	ticker := time.NewTicker(p.updateInterval)
	defer ticker.Stop()

	for {

		<-ticker.C
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

func (p *Parser) readTrxs(path string) []types.Transaction {
	trxs, err := p.repo.ReadTrxs(path)
	if err != nil {
		log.Println(err)
		return []types.Transaction{}
	}

	return trxs
}

func (p *Parser) writeTrxs(trxs map[string][]types.Transaction) error {
	return p.repo.WriteTrxs(trxs)
}

func (p *Parser) saveSubs() error {
	return p.repo.SaveSubs(p.subscriptions)
}

func (p *Parser) loadSubs() error {

	savedSubs, err := p.repo.LoadSubs()
	if err != nil {
		return err
	}

	p.subscriptions = savedSubs

	return nil
}
