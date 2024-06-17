package main

import (
	"context"
	"log"

	"projeto-api/internal/bff"
	_ "projeto-api/internal/bff"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), bff.Server); err != nil {
		log.Fatal(err)
	}
}
