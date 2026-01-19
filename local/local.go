package main

import (
	"fmt"
	"time"

	"discord-social-media-bot/main/lib"
)

func main() {
	for {
		// mock a request
		req := lib.Request{}

		// Get a greeting message and print it.
		res, err := lib.Main(req)
		if err != nil {
			// panic
			panic(err)
		}

		// serialize res and print
		fmt.Printf("Response: %+v\n", res)

		time.Sleep(time.Hour)
	}
}
