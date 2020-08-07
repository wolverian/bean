package file

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/wolverian/bean/brew"
)

const CaskAPI = "https://formulae.brew.sh/api/cask.json"

var CaskJSONFile = path.Join(cacheDir, "cask.json")

func ReadCasks() ([]brew.Cask, error) {
	if err := refreshCache(CaskJSONFile, CaskAPI); err != nil {
		return nil, err
	}

	jsonBuf, err := readCache(CaskJSONFile, CaskAPI)
	if err != nil {
		return nil, err
	}

	var casks []brew.Cask
	if err = json.Unmarshal(jsonBuf, &casks); err != nil {
		return nil, fmt.Errorf("cannot parse cask: %w", err)
	}

	return casks, nil
}
