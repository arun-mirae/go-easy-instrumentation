package main

import (
	"log"

	"github.com/arun-mirae/go-easy-instrumentation/cmd"
)

func main() {
	log.Default().SetFlags(0)
	cmd.Execute()
}
