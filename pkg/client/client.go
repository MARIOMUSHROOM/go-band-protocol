package client

import (
	"band_protocol_go/pkg/config"
	"band_protocol_go/pkg/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/exp/slog"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		HTTPClient: &http.Client{},
		BaseURL:    cfg.APIBaseURL,
	}
}

func (c *Client) GetStatus(endpoint string) (*entity.TxStatus, error) {
	path := c.BaseURL + endpoint
	slog.Info("Sending request", "method", http.MethodPost, "path", path)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response entity.TxStatus
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) PostTransaction(endpoint string, data interface{}) (*entity.TxHash, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	path := c.BaseURL + endpoint
	slog.Info("Sending request", "method", http.MethodPost, "path", path, "body", string(jsonData))

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response entity.TxHash
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
