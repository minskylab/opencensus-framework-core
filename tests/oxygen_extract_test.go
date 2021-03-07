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
		totalRecords += len(records)
	}

	if totalRecords != lapses*100 {
		t.Fail()
	}
}
