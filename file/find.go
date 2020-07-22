package file

import (
	"fmt"

	"github.com/wolverian/bean/brew"
)

func FindOne(name string, comp func(name, search string) bool) (brew.Formula, error) {
	fs, err := ReadFormulae()
	if err != nil {
		return brew.Formula{}, err
	}
	for _, f := range fs {
		if comp(f.Name, name) {
			return f, nil
		}
	}
	return brew.Formula{}, fmt.Errorf("no such formula")
}

func FindAll(name string, comp func(name, search string) bool) ([]brew.Formula, error) {
	fs, err := ReadFormulae()
	if err != nil {
		return nil, err
	}
	var found []brew.Formula
	for _, f := range fs {
		if comp(f.Name, name) {
			found = append(found, f)
		}
	}
	if len(found) == 0 {
		return nil, fmt.Errorf("no such formula")
	}
	return found, nil
}
