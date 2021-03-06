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
		totalRecords += len(records)
	}

	if totalRecords != 200 {
		t.Fail()
	}
}
