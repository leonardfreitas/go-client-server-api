package repositories

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/leonardfreitas/go-client-server-api/server/models"
)

func CreateDollar(ctx context.Context, dollarReal *models.DollarReal) error {
	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = RunMigrations(db)

	if err != nil {
		fmt.Println(err)
		return err
	}

	stmt, err := db.Prepare(`
		insert into dollars(code, codein, name, high, low, varbid, pct_change, bid, ask, timestamp, create_date) 
		values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		dollarReal.USDBRL.Code, dollarReal.USDBRL.Codein, dollarReal.USDBRL.Name, dollarReal.USDBRL.High,
		dollarReal.USDBRL.Low, dollarReal.USDBRL.VarBid, dollarReal.USDBRL.PctChange, dollarReal.USDBRL.Bid,
		dollarReal.USDBRL.Ask, dollarReal.USDBRL.Timestamp, dollarReal.USDBRL.CreateDate,
	)

	if err != nil {
		return err
	}

	return nil
}
