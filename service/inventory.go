package service

import (
	"fmt"
	"sync"
)

var inventory = make(map[string]int)
var mu sync.Mutex

func addInventory(ingredient string, quantity int) map[string]int {
	count, ok := inventory[ingredient]

	if ok {
		inventory[ingredient] = quantity + count //if item is already present add it to new items
	} else {
		inventory[ingredient] = quantity //else create new item with given count
	}
	return inventory
}

func UpdateInventory(beverage Beverage) bool {
	isPossible := true
	ingredientMap := beverage.IngredientQuantityMap
	mu.Lock() // for syncronization

	for k, v := range ingredientMap {
		ivgCount, ok := inventory[k]

		if !ok { //if item is not present
			fmt.Println(beverage.Name + " can't prepared because " + k + " is not available")
			isPossible = false
			break
		} else if ivgCount < v { // if item is not sufficient
			fmt.Println(beverage.Name + " can't prepared because " + k + " is not suffficient")
			isPossible = false
			break
		}
	}

	if isPossible { //if all items are present in sufficient quantity make the beverage
		for k1, v1 := range ingredientMap {
			inventory[k1] = inventory[k1] - v1
		}
		fmt.Println(beverage.Name + " is prepared.")
	}

	mu.Unlock()
	return isPossible
}
