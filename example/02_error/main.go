package main

import (
	"fmt"

	"log"
)

func main() {
	err := myFunc()
	if err != nil {
		log.Print("Error happened ", err)
		return
	}

	log.Print("Myfunc ran without error")
}

// add 10 to x, only accept positive value
func fnThatReturnError(x int) (int, error) {
	if x < 0 {
		return 0, fmt.Errorf("Can't handle negative value %d", x)
	}

	return x + 10, nil
}

func myFunc() error {
	i, err := fnThatReturnError(10)
	if err != nil {
		return err
	}

	fmt.Println("i is ", i)

	j, err := fnThatReturnError(-2)
	if err != nil {
		return err
	}

	fmt.Println(j)

	return nil
}
