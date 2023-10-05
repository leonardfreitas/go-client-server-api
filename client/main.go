package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/leonardfreitas/go-client-server-api/client/models"
)

func main() {
	const url = "http://localhost:8080/cotacao"

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var jsonResponse models.DollarReal
	json.Unmarshal(body, &jsonResponse)

	f, err := os.Create("cotacao.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	bid := jsonResponse.USDBRL.Bid

	_, err = f.Write([]byte(fmt.Sprintf("Dólar: %s", bid)))

	if err != nil {
		panic(err)
	}

}
