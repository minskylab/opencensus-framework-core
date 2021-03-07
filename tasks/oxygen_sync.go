package tasks

import (
	"context"
	"opencensus/core/client"
	"opencensus/core/extractors/oxygen"

	"github.com/pkg/errors"
)

func SyncOxygenRecords() error {
	entClient, err := client.NewClient()
	if err != nil {
		return errors.WithStack(err)
	}

	ctx := context.Background()

	records := oxygen.Extract(1)

	for recordBatch := range records {
		if err = oxygen.Processor(ctx, entClient, recordBatch); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
