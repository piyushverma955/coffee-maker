package service

import (
	"testing"
)

func TestAddInventory1(t *testing.T) {
	res := addInventory("sugar", 2)
	item, ok := res["hotWaters"]
	if ok && item == 2 {
		t.Error("add inventory test failed", ok)
	}
}

func TestAddInventory2(t *testing.T) {
	res := addInventory("hotWater", 2)
	item, ok := res["hotWater"]
	if !(ok && item == 2) {
		t.Error("add inventory test failed")
	}
}

func TestAddInventory3(t *testing.T) {
	res := addInventory("hotWater", 2)
	item, ok := res["hotWater"]
	if !(ok && item == 4) {
		t.Error("add inventory test failed")
	}
}

func TestUpdateInventory1(t *testing.T) {
	addInventory("ginger_syrup", 40)
	addInventory("hot_water", 400)
	addInventory("sugar_syrup", 60)
	addInventory("tea_leaves_syrup", 40)
	bverage := Beverage{Name: "black_tea", IngredientQuantityMap: map[string]int{"ginger_syrup": 30, "hot_water": 300, "sugar_syrup": 50, "tea_leaves_syrup": 30}}
	res := UpdateInventory(bverage)

	if !res {
		t.Error("update inventory test failed")
	}
}

func TestUpdateInventory2(t *testing.T) {
	bverage := Beverage{Name: "black_tea", IngredientQuantityMap: map[string]int{"ginger_syrup": 30, "hot_water": 300, "sugar_syrup": 50, "tea_leaves_syrup": 30}}
	res := UpdateInventory(bverage)

	if res {
		t.Error("update inventory test failed")
	}
}
