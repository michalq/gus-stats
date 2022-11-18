package main

import (
	"fmt"

	gus "github.com/michalq/gus-stats/pkg/client_gus"
)

func main() {
	gusConfig := gus.NewConfiguration()
	gusClient := gus.NewAPIClient(gusConfig)

	fmt.Println("Hello world!")
}
