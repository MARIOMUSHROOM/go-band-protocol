package service

import (
	"band_protocol_go/pkg/client"
	"band_protocol_go/pkg/entity"
	"fmt"
)

type APIService struct {
	Client *client.Client
}

// NewAPIService creates a new instance of APIService
func NewAPIService(client *client.Client) *APIService {
	return &APIService{Client: client}
}

// GetData fetches data from the API
func (s *APIService) GetData(endpoint string) (*entity.TxStatus, error) {
	var response *entity.TxStatus
	response, err := s.Client.GetStatus(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to get data from API: %w", err)
	}
	return response, nil
}

func (s *APIService) PostTransaction(endpoint string, data interface{}) (*entity.TxHash, error) {
	var response *entity.TxHash
	response, err := s.Client.PostTransaction(endpoint, data)
	if err != nil {
		return nil, fmt.Errorf("failed to get data from API: %w", err)
	}
	return response, nil
}
