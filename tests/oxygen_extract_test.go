package tests

import (
	"opencensus/core/extractors/oxygen"
	"testing"
)

func TestExtractOxygenData(t *testing.T) {
	records := oxygen.Extract(2)

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
