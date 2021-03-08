package tests

import (
	"opencensus/core/extractors/oxygen"
	"testing"
)

func TestExtractOxygenData(t *testing.T) {
	lapses := 5
	records := oxygen.Extract(lapses)

	totalRecords := 0

	for batch := range records {
		t.Logf("batched records: %d", len(batch))
		t.Logf("head: %+v", batch[0])

		totalRecords += len(batch)
	}
	t.Logf("totalRecords: %d", totalRecords)

	if totalRecords != lapses*100 {
		t.Fail()
	}

	t.FailNow()
}
