package tasks

import (
	"context"
	"opencensus/core/client"
	"opencensus/core/extractors/bed"

	"github.com/pkg/errors"
)

func SyncBedRecords() error {
	entClient, err := client.NewClient()
	if err != nil {
		return errors.WithStack(err)
	}

	ctx := context.Background()

	records := bed.Extract(1000)

	for recordBatch := range records {
		if err = bed.Processor(ctx, entClient, recordBatch); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
