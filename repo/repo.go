package repo

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"parser/types"
)

type Repo struct {
	subsPath string
}

func New() *Repo {

	err := os.MkdirAll("storage", 0700)
	if err != nil {
		log.Println(err)
	}

	return &Repo{subsPath: filepath.Join("storage", "subs")}
}

func (r *Repo) ReadTrxs(address string) ([]types.Transaction, error) {
	return readTrxs(address)
}

func (r *Repo) WriteTrxs(trxs map[string][]types.Transaction) error {

	for addr, newTrxs := range trxs {
		trxs, err := readTrxs(addr)
		if err != nil {
			trxs = newTrxs
		} else {
			trxs = append(trxs, newTrxs...)
		}

		err = write(addr, trxs)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repo) SaveSubs(data map[string]struct{}) error {
	return write(r.subsPath, data)
}

func (r *Repo) LoadSubs() (map[string]struct{}, error) {
	bytes, err := os.ReadFile(r.subsPath)
	if err != nil {
		return nil, err
	}

	var resp map[string]struct{}

	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func readTrxs(address string) ([]types.Transaction, error) {

	bytes, err := os.ReadFile(address)
	if err != nil {
		return nil, err
	}

	var resp []types.Transaction

	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func write(path string, data any) error {

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(path, bytes, 0700)
}
