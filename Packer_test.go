package main

import (
	"fmt"
	"testing"
)

func TestPackItems_251Items(t *testing.T) {
	availablePacks := []int{5000, 2000, 1000, 500, 250}
	quantity := 251

	res := packItems(availablePacks, quantity)
	//NOT TRUE
	if len(res) != 1 {
		t.Error("Expected 1 pack type", res)
	}

	validatePack(t, res, 500, 1)
}

func TestPackItems_501Items(t *testing.T) {
	availablePacks := []int{5000, 2000, 1000, 500, 250}
	quantity := 501

	res := packItems(availablePacks, quantity)
	//NOT TRUE
	if len(res) != 2 {
		t.Error("Expected 2 pack types", res)
	}

	validatePack(t, res, 250, 1)
	validatePack(t, res, 500, 1)
}

func TestPackItems_1Item(t *testing.T) {
	availablePacks := []int{5000, 2000, 1000, 500, 250}
	quantity := 1

	res := packItems(availablePacks, quantity)
	//NOT TRUE
	if len(res) != 1 {
		t.Error("Expected 1 pack type", res)
	}

	validatePack(t, res, 250, 1)
}

func TestPackItems_12001Items(t *testing.T) {
	availablePacks := []int{5000, 2000, 1000, 500, 250}
	quantity := 12001

	res := packItems(availablePacks, quantity)
	//NOT TRUE
	if len(res) != 3 {
		t.Error("Expected 3 pack type", res)
	}

	validatePack(t, res, 250, 1)
	validatePack(t, res, 2000, 1)
	validatePack(t, res, 5000, 2)
}

func validatePack(t *testing.T, res map[int]PackedEntry, packSize int, quantity int) {
	pack := res[packSize]
	if pack.quantity != quantity {
		errMsg := fmt.Sprintf("Expected %d packs of %d, got %d", quantity, packSize, pack.quantity)
		t.Error(errMsg)
	}
	if pack.pack_volume != packSize {
		errMsg := fmt.Sprintf("Pack of %d does not have %d, it has %d", packSize, packSize, pack.pack_volume)
		t.Error(errMsg)
	}
}
