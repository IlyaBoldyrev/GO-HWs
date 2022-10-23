package file

import (
	"encoding/json"
	"os"

	"github.com/IlyaBoldyrev/GO-HWs/HW9/another"
)

func GetConfigFromFile(str string) (another.Config, error) {
	f, err := os.Open(str)
	if err != nil {
		return another.Config{}, err
	}
	defer f.Close()

	var s another.Config
	if err := json.NewDecoder(f).Decode(&s); err != nil {
		return another.Config{}, err
	}

	return s, nil
}
