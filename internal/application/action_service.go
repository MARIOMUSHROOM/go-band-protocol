package application

import (
	"band_protocol_go/internal/ports"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/exp/slog"
)

type ActionService struct {
}

func NewActionService() ports.ActionService {
	return &ActionService{}
}

func (s *ActionService) CheckRevenge(text string) string {
	var buffer bytes.Buffer
	var shots []string
	lenText := len(text)
	if lenText == 0 || text[0] == 'R' || text[lenText-1] == 'S' {
		return "Bad boy"
	}
	slog.Info("Info", "len", lenText)
	for i, c := range text {
		if i > 0 {
			if text[i-1] == 'R' && c == 'S' {
				newText := buffer.String()
				shots = append(shots, newText)
				buffer.Reset()
			}
		}
		if c == 'S' {
			buffer.WriteString("S")
		}
		if c == 'R' {
			buffer.WriteString("R")
		}
		if (i + 1) == lenText {
			newText := buffer.String()
			shots = append(shots, newText)
			buffer.Reset()
		}
		slog.Info("Info", "i", i, "len", lenText, "buffer", buffer.String())
	}
	slog.Info("Info", "group", shots)

	shotsLength := len(shots)
	for i, short := range shots {
		slog.Info("Info", "short", short)
		var shotCount, revengeCount int
		for _, sr := range short {
			slog.Info("short", "sr", short)
			if sr == 'S' {
				shotCount++
			}
			if sr == 'R' {
				revengeCount++
			}
		}
		if (i + 1) == shotsLength {
			if shotCount > revengeCount {
				return "Bad boy"
			}
		}
	}
	return "Good boy"
}

func (s *ActionService) MaxChickensProtected(n, k int, positions []int) int {
	left := 0
	maxProtected := 0
	for right := 0; right < n; right++ {
		for positions[right]-positions[left] >= k {
			left++
		}
		count := right - left + 1
		if count > maxProtected {
			maxProtected = count
		}
	}
	return maxProtected
}

func PostData(apiURL string, authToken string, requestData interface{}) (string, error) {
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request data to JSON: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if authToken != "" {
		req.Header.Set("Authorization", "Bearer "+authToken)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
