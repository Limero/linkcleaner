package main

import (
	"fmt"
	"log"
	"os"

	"github.com/limero/linkcleaner"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./linkcleaner <url>")
		return
	}

	url, err := linkcleaner.CleanURLString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url.String())
}
