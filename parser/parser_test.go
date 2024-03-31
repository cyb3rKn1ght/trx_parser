package parser_test

import (
	"parser/parser"
	"parser/repo"
	"testing"
)

func TestGetCurrentBlock(t *testing.T) {
	p := parser.New(repo.New(), "https://cloudflare-eth.com")
	blockNum := p.GetCurrentBlock()

	if blockNum <= 0 {
		t.Errorf("blockNum: %v", blockNum)
	}
}
