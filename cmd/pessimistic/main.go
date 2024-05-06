package main

import (
	"backend/internal/config"
	"backend/internal/model"
	postgres "backend/internal/util"
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	fmt.Println("----------------------")
	appConfig := config.NewAppConfig()
	db := postgres.NewPostgres(postgres.PostgresConfig{
		DbConnectionStr: appConfig.DbConnectionStr,
	})
	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to database!")
		return
	}
	fmt.Println("Connected to database!")

	// list all inventory
	ctx := context.Background()
	inventories, err := model.Inventories().All(ctx, db)
	if err != nil {
		fmt.Println("Failed to list inventories!")
		return
	}

	fmt.Println("list all inventory")
	for _, inventory := range inventories {
		fmt.Println("inventory-->", inventory.ID, inventory.Name, inventory.Amount)
	}
	fmt.Print("\n")
	targetID := "mock-id-1"
	delay := 0
	if len(os.Args) > 1 {
		delayStr := os.Args[1]
		delayNum, err := strconv.Atoi(delayStr)
		if err != nil {
			fmt.Println("Invalid delay value")
			return
		}
		delay = delayNum
	}
	performUpdate(ctx, db, targetID, -1, delay)

	fmt.Println("----------------------")
}

// delay is in seconds
func performUpdate(ctx context.Context, db *sql.DB, targetID string, incrVal int, delay int) error {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Failed to begin transaction!")
		return err
	}
	inventory, err := model.Inventories(model.InventoryWhere.ID.EQ(targetID), qm.For("update")).One(ctx, tx)
	if err != nil {
		fmt.Println("Failed to find inventory!")
		return err
	}
	fmt.Println("Get inventory-->", inventory.ID, inventory.Name, inventory.Amount)

	// simulate delay
	time.Sleep(time.Duration(delay) * time.Second)

	inventory.Amount.Int = inventory.Amount.Int + incrVal
	_, err = inventory.Update(ctx, tx, boil.Whitelist("amount"))
	if err != nil {
		fmt.Println("Failed to update inventory!")
		return err
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("Failed to commit transaction!")
		tx.Rollback()
		return err
	}

	// print updated inventory
	fmt.Println("Updated inventory-->", inventory.ID, inventory.Name, inventory.Amount)
	fmt.Print("\n")
	fmt.Println("Update inventory successfully!")
	return nil
}
