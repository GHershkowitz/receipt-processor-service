package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/shopspring/decimal"
)

const (
	retailPointValue = 1
	purchaseDatePointValue = 6
	purchaseTimePointValue = 10
	itemCountPointValue = 5
	itemDescriptionPointValue = .2
	totalRoundDollarPointValue = 50
	totalMultipleOfTwentyFivePointValue = 25
)

func CreateReceipt(receiptID string, receipt Receipt) {
	dbReceipts[receiptID] = &receipt
}

func CalculateReceiptScore(receipt *Receipt) int {
	points := 0
	if receipt == nil {
		return 0
	}

	points += calculateRetailerPoints(receipt.Retailer)
	points += calculatePurchaseDatePoints(receipt.PurchaseDate)
	points += calculatePurchaseTimePoints(receipt.PurchaseTime)
	points += calculateItemsPoints(receipt.Items)
	points += calculateTotalPoints(receipt.Total)

	return points
}

func calculateRetailerPoints(retailer string) int {
	points := 0
	for _, char := range retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points += retailPointValue
		}
	}

	fmt.Printf("retailer points %d\n", points)
	return points
}

// calculatePurchaseDatePoints calculates points awarded based on the purchaseDate on the receipts.
func calculatePurchaseDatePoints(purchaseDate string) int {
	points := 0
	date, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		fmt.Println("error parsing purchaseDate", err)
		return 0
	}

	// Check if date is odd
	if date.Day()%2 != 0 {
		points += purchaseDatePointValue
	}

	fmt.Printf("purchaseDate points %d\n", points)
	return points
}

// calculatePurchaseTimePoints calculates points awarded based on the purchaseTime on the receipts.
func calculatePurchaseTimePoints(purchaseTime string) int {
	points := 0
	pTime, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		fmt.Println("error parsing purchaseTime", err)
		return 0
	}
	// 2pm
	startTime := time.Date(pTime.Year(), pTime.Month(), pTime.Day(), 14, 0, 0, 0, pTime.Location())
	// 4pm
	endTime := time.Date(pTime.Year(), pTime.Month(), pTime.Day(), 16, 0, 0, 0, pTime.Location())

	if pTime.After(startTime) && pTime.Before(endTime) {
		points += purchaseTimePointValue
	}

	fmt.Printf("purchaseTime points %d\n", points)
	return points
}

// calculateTotalPoints calculate how many points are awarded for the total price of the items.
// 50 points if the total is a round dollar amount with no cents.
// 25 points if the total is a multiple of 0.25.
func calculateTotalPoints(total string) int {
	points := 0
	totalArr := strings.SplitAfter(total, ".")

	if len(totalArr) < 2 {
		return 0
	}

	valueAfterDecimal := totalArr[1]
	tempTotal, _ := strconv.Atoi(valueAfterDecimal)

	// if the tempTotal is 0 and the first value in the string is equal to "0" (rune value 48)
	if tempTotal == 0 && valueAfterDecimal[1] == 48 {
		points += totalRoundDollarPointValue
	}

	if tempTotal%25 == 0 {
		points += totalMultipleOfTwentyFivePointValue
	}

	fmt.Printf("total points %d\n", points)
	return points
}

// calculateItemsPoints calculates points awarded based on the items on the receipt.
func calculateItemsPoints(items []*Item) int {
	points := 0
	if items == nil {
		return 0
	}

	for _, item := range items {
		if item == nil {
			continue
		}
		points += itemDescriptionLengthPoints(item)
	}
	points += itemCountPoints(items)

	fmt.Printf("item points %d\n", points)
	return points
}

// itemCountPoints calculates how many points are awarded based on the number of items on the receipt.
// 5 points are awarded for every two items on the receipt.
func itemCountPoints(items []*Item) int {
	return int(math.Floor(float64(len(items)/2))) * itemCountPointValue
}

// itemCountPoints calculates how many points are awarded based on the description length of an item.
// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
// The result is the number of points earned.
func itemDescriptionLengthPoints(item *Item) int {
	points := 0
	if item == nil {
		return 0
	}

	price, err := decimal.NewFromString(item.Price)
	if err != nil {
		fmt.Printf("unable to convert %s to decimal, err: %v\n", item.Price, err)
		return 0
	}

	trimmedDesc := strings.TrimSpace(item.Description)
	if len(trimmedDesc)%3 == 0 {
		priceMultiplier := decimal.NewFromFloat(itemDescriptionPointValue)
		points += int(price.Mul(priceMultiplier).Ceil().IntPart())
	}

	return points
}
