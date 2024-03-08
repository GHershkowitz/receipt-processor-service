package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculateRetailerPoints(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		expected int
	}{
		{"full-dollar", "1.00", 75},
		{"half-dollar", "100.50", 25},
		{"no-points", "10.57", 0},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			total := calculateTotalPoints(test.input)
			assert.Equal(t, test.expected, total)
		})
	}

}

func Test_calculatePurchaseDatePoints(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		expected int
	}{
		{"odd-date", "2021-03-07", 6},
		{"even-date", "2023-06-10", 0},
		{"empty-date", "", 0},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			total := calculatePurchaseDatePoints(test.input)
			assert.Equal(t, test.expected, total)
		})
	}

}

func Test_calculatePurchaseTimePoints(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		expected int
	}{
		{"time-is-between", "14:25", 10},
		{"time-is-not-between", "16:01", 0},
		{"time-is-empty", "", 0},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			total := calculatePurchaseTimePoints(test.input)
			assert.Equal(t, test.expected, total)
		})
	}

}
func Test_itemCountPoints(t *testing.T) {
	tests := []struct {
		testName string
		input    []*Item
		expected int
	}{
		{"ten-items", testItems(10), 25},
		{"nine-items", testItems(9), 20},
		{"four-items", testItems(4), 10},
		{"no-items", nil, 0},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			total := itemCountPoints(test.input)
			assert.Equal(t, test.expected, total)
		})
	}

}

func Test_itemDescriptionLengthPoints(t *testing.T) {
	tests := []struct {
		testName string
		input    *Item
		expected int
	}{
		{"multiple-of-3-test-1", &Item{Description: "the best des", Price: "2.25"}, 1},
		{"multiple-of-3-test-2", &Item{Description: "the best des", Price: "22.92"}, 5},
		{"not-multiple-of-3", &Item{Description: "the worst des", Price: "3.00"}, 0},
		{"nil-item", nil, 0},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			total := itemDescriptionLengthPoints(test.input)
			assert.Equal(t, test.expected, total)
		})
	}

}

func Test_calculateTotalPoints(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		expected int
	}{
		{"full-dollar", "1.00", 75},
		{"half-dollar", "100.50", 25},
		{"no-points", "10.57", 0},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			total := calculateTotalPoints(test.input)
			assert.Equal(t, test.expected, total)
		})
	}

}

func testItems(count int) []*Item {
	output := make([]*Item, count)
	for i := 0; i < count; i++ {
		item := &Item{}
		output[i] = item
	}

	return output
}
