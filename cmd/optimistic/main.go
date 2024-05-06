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
)

func main() {
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
	targetID := "mock-id-1"
	ctx := context.Background()
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

}

func performUpdate(ctx context.Context, db *sql.DB, targetID string, incrVal int, delay int) error {
	inventoryItem, err := model.FindInventory(ctx, db, targetID)
	if err != nil {
		fmt.Println("Failed to find inventory item")
		return err
	}
	currentVersion := inventoryItem.Version
	// check if the version is same
	if inventoryItem.Version != currentVersion {
		fmt.Println("Version mismatch")
		return err
	}

	// simulate delay
	time.Sleep(time.Duration(delay) * time.Second)

	// update the amount and increment the version
	inventoryItem.Amount.Int = inventoryItem.Amount.Int + incrVal
	inventoryItem.Version++

	// Update the inventory item with the new amount and version
	updateCount, err := model.Inventories(
		model.InventoryWhere.ID.EQ(targetID),
		model.InventoryWhere.Version.EQ(currentVersion),
	).UpdateAll(ctx, db, model.M{
		"amount":  inventoryItem.Amount,
		"version": inventoryItem.Version,
	})
	fmt.Println("updateCount", updateCount)
	if updateCount == 0 {
		fmt.Println("Failed to update inventory item")
		newErr := fmt.Errorf("Failed to update inventory item")
		return newErr

	}
	if err != nil {
		fmt.Println("Failed to update inventory item")
		return err
	}
	fmt.Println("Update inventory successfully!")
	return nil
}
