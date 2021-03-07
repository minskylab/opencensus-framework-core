package positives

import (
	"os"
)

func extract(filename string) error {
	os.OpenFile(filename, os.O_RDONLY, 0644)

	// csv.NewReader()
	return nil
}
