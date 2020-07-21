package file

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/wolverian/bean/brew"
)

const FormulaeJSONFile = "homebrew-formulae.json"
const FormulaeAPI = "https://formulae.brew.sh/api/formula.json"

var cacheDir = must(os.UserCacheDir())
var cacheFile = path.Join(cacheDir, FormulaeJSONFile)

func ReadFormulae() ([]brew.Formula, error) {
	if err := refreshCache(); err != nil {
		return nil, err
	}

	jsonBuf, err := readCache()
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

func readCache() ([]byte, error) {
	jsonBuf, err := ioutil.ReadFile(cacheFile)
	if err != nil {
		if err := download(); err != nil {
			return nil, err
		}

		// Retry once
		jsonBuf, err = ioutil.ReadFile(FormulaeJSONFile)
		if err != nil {
			return nil, fmt.Errorf("could not read input: %w", err)
		}
	}

	return jsonBuf, nil
}

func refreshCache() error {
	if stat, err := os.Stat(FormulaeJSONFile); err == nil {
		// Ignore err
		if stat.ModTime().Before(time.Now().Add(-time.Hour * 24 * 7)) {
			if err := download(); err != nil {
				return err
			}
		}
	}
	return nil
}

func download() error {
	resp, err := http.Get(FormulaeAPI)
	if err != nil {
		return fmt.Errorf("could not download formulae json from %s: %w", FormulaeAPI, err)
	}
	defer resp.Body.Close()

	f, err := os.Create(cacheFile)
	if err != nil {
		return fmt.Errorf("could not save cache: %w", err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("could not save cache: %w", err)
	}

	return nil
}

func must(dir string, err error) string {
	if err != nil {
		panic(err)
	}
	println(dir)
	return dir
}
