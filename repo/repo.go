package repo

import (
	"encoding/json"
	"os"

	"parser/types"
)

type Repo struct{}

func (r *Repo) Read(address string) ([]types.Transaction, error) {
	return read(address)
}

func (r *Repo) Write(data map[string][]types.Transaction) error {

	for addr, newTrxs := range data {
		trxs, err := read(addr)
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

func read(address string) ([]types.Transaction, error) {

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

func write(address string, trxs []types.Transaction) error {

	bytes, err := json.Marshal(trxs)
	if err != nil {
		return err
	}

	return os.WriteFile(address, bytes, 0700)
}
