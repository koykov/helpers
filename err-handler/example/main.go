package main

import (
	"errors"
	"fmt"
	"github.com/koykov/helpers/err-handler"
)

func main() {
	fmt.Println("Before try-catch.")
	err_handler.TryCatch{
		Try: func() {
			if err := errors.New("artificial panic"); err != nil {
				panic(err)
			}
		},
		Catch: func(e err_handler.Exception) {
			fmt.Printf("Oops, we caught a panic: %s\n", e)
		},
	}.Do()
	fmt.Println("Hey, I've survived the panic!")
}
