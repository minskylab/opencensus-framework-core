package bed

import (
	"fmt"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	regDate := "20071009"

	layout1 := "20060102"
	t1, err := time.Parse(layout1, regDate)
	if err != nil {
		t.Fail()
	}
	// fmt.Printf("utc: %v\n", t1.Format("2017-01-22"))
	fmt.Printf("%v\n", t1)
}
