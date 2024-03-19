package repo

import (
	"encoding/json"
	"os"

	"parser/types"
)

type Repo struct{}

func (r *Repo) ReadTrxs(path string) ([]types.Transaction, error) {
	return readTrxs(path)
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

func (r *Repo) SaveSubs(path string, data map[string]struct{}) error {
	return write(path, data)
}

func (r *Repo) LoadSubs(path string) (map[string]struct{}, error) {
	bytes, err := os.ReadFile(path)
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

func readTrxs(path string) ([]types.Transaction, error) {

	bytes, err := os.ReadFile(path)
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
