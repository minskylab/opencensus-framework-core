package tests

import (
	"opencensus/core/extractors/oxygen"
	"testing"
)

func TestExtractOxygenData(t *testing.T) {
	records := oxygen.Extract(2)

	for batch := range records {
		t.Log(batch)
	}

	// if len(records) != 200 {
	// 	t.FailNow()
	// }
}
