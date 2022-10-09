package file_system

import (
	"os"
)

func Read0(pathFile string) ([]byte, error) {
	f, err := os.ReadFile(pathFile)
	if err != nil {
		return nil, err
	}
	return f, nil
}
