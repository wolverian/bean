package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var cacheDir = must(os.UserCacheDir())

func readCache(file, url string) ([]byte, error) {
	jsonBuf, err := ioutil.ReadFile(file)
	if err != nil {
		if err := download(file, url); err != nil {
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

func refreshCache(file, url string) error {
	if stat, err := os.Stat(file); err == nil {
		// Ignore err
		if stat.ModTime().Before(time.Now().Add(-time.Hour * 24)) {
			if err := download(file, url); err != nil {
				return err
			}
		}
	}
	return nil
}

func download(file, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not download formulae json from %s: %w", url, err)
	}
	defer mustClose(resp.Body)

	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("could not save cache: %w", err)
	}
	defer mustClose(f)

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("could not save cache: %w", err)
	}

	return nil
}

func mustClose(resp io.Closer) func() {
	return func() {
		if err := resp.Close(); err != nil {
			panic(err)
		}
	}
}

func must(dir string, err error) string {
	if err != nil {
		panic(err)
	}
	return dir
}
