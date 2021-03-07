package main

import (
	"fmt"
	"opencensus/core/tasks"
)

func main() {
	if err := tasks.SyncBedRecords(); err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}
}
