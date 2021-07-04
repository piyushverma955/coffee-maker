package main

import (
	"coffee-maker/service"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//taking file name from env variable
	file_name := os.Getenv("FILE_NAME")
	data, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Printf("failed reading data from file: %s\n", err)
		os.Exit(0)
	}

	service.Process(data) //starting the process
}
