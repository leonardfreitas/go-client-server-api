package services

import (
	"context"
	"time"

	"github.com/leonardfreitas/go-client-server-api/server/models"
	"github.com/leonardfreitas/go-client-server-api/server/repositories"
)

func GetDollarService() (*models.DollarReal, error) {
	httpCtx := context.Background()
	httpCtx, httpCancel := context.WithTimeout(httpCtx, 200*time.Millisecond)
	defer httpCancel()
	result, err := repositories.HttpGetDollarToReal(httpCtx)

	if err != nil {
		return nil, err
	}

	dbCtx := context.Background()
	dbCtx, dbCancel := context.WithTimeout(dbCtx, 10*time.Millisecond)
	defer dbCancel()
	err = repositories.CreateDollar(dbCtx, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
