package repositories

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/leonardfreitas/go-client-server-api/server/models"
)

const url = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

func HttpGetDollarToReal(ctx context.Context) (*models.DollarReal, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var response_to_model models.DollarReal
	json.Unmarshal([]byte(result), &response_to_model)

	return &response_to_model, nil
}
