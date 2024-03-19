package parser

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func makeRequest(url string, body reqRPC) ([]byte, error) {
	reqBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
