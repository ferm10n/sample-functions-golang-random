package lib

import (
	"fmt"
	"math/rand"
	"time"

	"rsc.io/quote"
)

const (
	// upperBound is the upper bound of the random int range
	upperBound = 1001
)

func Init() {
	rand.Seed(time.Now().Unix())
}

// Request is the function's request struct
type Request struct {
}

// Response is the function's response struct
type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	fmt.Println(quote.Go())
	return &Response{
		StatusCode: 200,
		Body:       fmt.Sprintf("Hello, your random number is %d", rand.Intn(upperBound)),
	}, nil
}
