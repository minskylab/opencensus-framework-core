package main

import (
	"fmt"
	"opencensus/core/client"
)

func main() {
	client, err := client.NewClient()
	if err != nil {
		panic(err)
	}

	fmt.Println(client)
	// cl.Place.Create().SetKind("new place")
}
