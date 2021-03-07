package main

import (
	"fmt"
	"opencensus/core/client"
)

func main() {
	cl, err := client.NewClient()
	fmt.Println(err)
	cl.Place.Create().SetKind("new place")
}
