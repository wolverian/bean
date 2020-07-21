package file

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/wolverian/bean/brew"
)

const FormulaeJSONFile = "homebrew-formulae.json"
const FormulaeAPI = "https://formulae.brew.sh/api/formula.json"

var cacheDir = must(os.UserCacheDir())

func ReadFormulae() ([]brew.Formula, error) {
	jsonBuf, err := ioutil.ReadFile(FormulaeJSONFile)
	if err != nil {
		resp, err := http.Get(FormulaeAPI)
		if err != nil {
			return nil, fmt.Errorf("could not download formulae json from %s: %w", FormulaeAPI, err)
		}
		defer resp.Body.Close()

		f, err := os.Create(path.Join(cacheDir, FormulaeJSONFile))
		if err != nil {
			return nil, fmt.Errorf("could not save cache: %w", err)
		}
		defer f.Close()
		_, err = io.Copy(f, resp.Body)
		if err != nil {
			return nil, fmt.Errorf("could not save cache: %w", err)
		}
		// Retry once
		jsonBuf, err = ioutil.ReadFile(FormulaeJSONFile)
		if err != nil {
			return nil, fmt.Errorf("could not read input: %w", err)
		}
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

func must(dir string, err error) string {
	if err != nil {
		panic(err)
	}
	return dir
}
