package service

import (
	"encoding/json"
	"fmt"
)

func Process(data []byte) {
	var coffeeMachineDetails CoffeeMachineDetails
	var coffeeMachine CoffeeMachine
	err := json.Unmarshal(data, &coffeeMachineDetails) //converting byte into json

	coffeeMachine = coffeeMachineDetails
	if err != nil {
		fmt.Println(err)
	}

	coffeeMachine.Start() //starting the process
}
