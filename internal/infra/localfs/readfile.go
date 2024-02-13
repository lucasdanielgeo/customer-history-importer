package localfs

import (
	"bufio"
	"os"
)

func NewFileScanner(path string) (*bufio.Scanner, *os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)

	return scanner, file, nil
}
