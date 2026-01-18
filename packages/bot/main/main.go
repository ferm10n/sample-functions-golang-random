package main

import (
	"bot/main/lib"
)

func init() {
	lib.Init()
}

func Main(in lib.Request) (*lib.Response, error) {
	return lib.Main(in)
}
