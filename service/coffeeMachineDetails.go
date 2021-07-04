package service

import (
	"sync"
)

type CoffeeMachine interface {
	Start()
}

type CoffeeMachineDetails struct {
	Machine Machine `json:"machine"`
}

func (c CoffeeMachineDetails) Start() {
	outlet := c.Machine.Outlets.Count
	ingredients := c.Machine.IngredientQuantityMap

	inputChannel := make(chan Beverage)
	closeChannels := make([]chan int, outlet)
	beverages := c.Machine.Beverages
	var wg sync.WaitGroup

	for k, v := range ingredients {
		addInventory(k, v) //adding incredients to inventory
	}

	for i := 0; i < outlet; i++ { // starting outets number of threads
		closeChannels[i] = startRoutine(inputChannel, i, &wg)
	}

	for k, v := range beverages {
		wg.Add(1)
		inputChannel <- Beverage{Name: k, IngredientQuantityMap: v} //it doest not close the channel until all beverages are recieved
	}

	wg.Wait()

	for i := 0; i < outlet; i++ {
		closeChannels[i] <- 1 //closing all goroutines/threads
	}

}

func startRoutine(inputChannel chan Beverage, goValue int, wg *sync.WaitGroup) chan int {
	closeChannel := make(chan int)
	go func(inputChannel chan Beverage, goValue int, wg *sync.WaitGroup) {
		for {
			select {
			case <-closeChannel: // close signal for closing the goroutines/threads
				return
			case beverage := <-inputChannel: //receiving beverage via channel for processing
				UpdateInventory(beverage)
				wg.Done()
			}
		}

	}(inputChannel, goValue, wg)

	return closeChannel
}
