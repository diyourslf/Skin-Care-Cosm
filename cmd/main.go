package main

import (
	"log"
	"skincarecosm"
)

func main() {
	srv := new(skincarecosm.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatal("error", err.Error())
	}
}
