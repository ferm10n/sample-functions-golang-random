package main

import (
	"fmt"
	"log"
	"time"

	"discord-social-media-bot/main/lib"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	for {
		// mock a request
		req := lib.Request{}

		// Get a greeting message and print it.
		res, err := lib.Main(req)
		if err != nil {
			log.Fatal(err)
		}

		// serialize res and print
		fmt.Printf("Response: %+v\n", res)

		time.Sleep(time.Hour)
	}
}
