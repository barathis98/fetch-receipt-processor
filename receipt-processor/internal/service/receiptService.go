package service

import (
	"fmt"
	"math"
	"receipt-processor/internal/model"
	"strings"
	"time"

	"github.com/google/uuid"
)

var receipt = make(map[string]model.Receipt)

func saveReceipt(receipt model.Receipt) (string, error) {
	// receipt[receipt.ID] = receipt
	receipt.ID = uuid.New().String()

	if _, err := time.Parse("2006-01-02", receipt.PurchaseDate); err != nil {
		return "", fmt.Errorf("invalid date format: %s", receipt.PurchaseDate)
	}
	if _, err := time.Parse("15:04", receipt.PurchaseTime); err != nil {
		return "", fmt.Errorf("invalid time format: %s", receipt.PurchaseTime)
	}
	receipt.Points = calculatePoints()
	return receipt.ID, nil
}

func calculatePoints(receipt model.Receipt) int {
	points := 0
	for _, ch := range receipt.Retailer {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			points++
		}
	}

	if receipt.Total == float64(int(receipt.Total)) {
		points += 50
	}

	if math.Mod(receipt.Total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 != 0 {
		points += 6
	}

	for _, item := range receipt.Items {
		description := strings.TrimSpace(item.ShortDesc)
		if len(description)%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.Hour() == 14 && purchaseTime.Minute() >= 0 && purchaseTime.Minute() < 60 {
		points += 10
	}

	return points
}
