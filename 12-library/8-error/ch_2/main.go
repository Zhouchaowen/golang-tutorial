package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func foo() error {
	return errors.Wrap(sql.ErrNoRows, "foo failed")
}

func bar() error {
	return errors.WithMessage(foo(), "bar failed")
}

func main() {
	err := bar()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		// unknown error
	}
}
