package main

import (
	"notion2atlas/presentation"

	"github.com/joho/godotenv"
)

func main() {
	var err error = nil
	godotenv.Load()
	err = presentation.HandleUpdateData()
	if err != nil {
		panic(err)
	}
}
