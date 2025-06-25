package store

import (
	"encoding/csv"
	"fmt"
	"os"
)

func read(filename string) ([][]string, error) {
	file, err := os.Open(fmt.Sprintf("data/%s.csv", filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
