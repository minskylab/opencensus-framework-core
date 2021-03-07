package tests

import (
	"opencensus/core/extractors/bed"
	"testing"
)

func TestExtractBedData(t *testing.T) {
	records := bed.Extract(2)

	totalRecords := 0

	for batch := range records {
		t.Logf("batched records: %d", len(batch))
		t.Logf("head: %+v", batch[0])

		totalRecords += len(batch)
	}

	t.Logf("totalRecords: %d", totalRecords)

	if totalRecords != 200 {
		t.Fail()
	}

	t.FailNow()
}
