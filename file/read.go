package file

import (
	"bean/brew"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const FormulaeJSONFile = "homebrew-formulae.json"

func ReadFormulae() ([]brew.Formula, error) {
	jsonBuf, err := ioutil.ReadFile(FormulaeJSONFile)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	// HACK
	for i, b := range jsonBuf {
		if b == '\x1b' {
			jsonBuf[i] = 'E'
		}
	}

	var fs []brew.Formula

	if err = json.Unmarshal(jsonBuf, &fs); err != nil {
		return nil, fmt.Errorf("could not parse formulae: %w", err)
	}

	return fs, nil
}
