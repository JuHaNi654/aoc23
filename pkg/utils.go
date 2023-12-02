package pkg

import (
	"bytes"
	"os"
)

func readFile(f string) ([]byte, error) {
	data, err := os.ReadFile(f)
	data = bytes.TrimSuffix(data, []byte("\n"))
	return data, err
}
