package main

import (
	"errors"
	"fmt"
)

func func1() error {
	return errors.New("func1 error")
}

func func2() error {
	return fmt.Errorf("func2 error: %w", func1())
}

func main() {
	err := func2()
	if err != nil {
		fmt.Println(err)
	}
}
