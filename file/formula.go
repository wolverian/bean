package file

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/wolverian/bean/brew"
)

const FormulaeAPI = "https://formulae.brew.sh/api/formula.json"

var FormulaeJSONFile = path.Join(cacheDir, "homebrew-formulae.json")

func ReadFormulae() ([]brew.Formula, error) {
	if err := refreshCache(FormulaeJSONFile, FormulaeAPI); err != nil {
		return nil, err
	}

	jsonBuf, err := readCache(FormulaeJSONFile, FormulaeAPI)
	if err != nil {
		return nil, err
	}

	return parseFormulae(jsonBuf)
}

func parseFormulae(jsonBuf []byte) ([]brew.Formula, error) {
	// HACK
	for i, b := range jsonBuf {
		if b == '\x1b' {
			jsonBuf[i] = 'E'
		}
	}

	var fs []brew.Formula
	if err := json.Unmarshal(jsonBuf, &fs); err != nil {
		return nil, fmt.Errorf("could not parse formulae: %w", err)
	}

	return fs, nil
}
