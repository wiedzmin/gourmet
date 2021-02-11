package main

import (
	"fmt"
	"os"

	"github.com/wiedzmin/gourmet/impl"
)

func main() {
	app := impl.CreateCLI()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
