package file

import (
	"fmt"

	"github.com/wolverian/bean/item"
)

type CompFunc func(name, search string) bool

func FindAll(substr string, comp CompFunc) ([]item.Interface, error) {
	var found []item.Interface

	fs, err := ReadFormulae()
	if err != nil {
		return nil, err
	}
	for _, f := range fs {
		if comp(f.CanonicalName(), substr) {
			found = append(found, f)
		}
	}

	cs, err := ReadCasks()
	if err != nil {
		return nil, err
	}
	for _, c := range cs {
		if comp(c.CanonicalName(), substr) {
			found = append(found, c)
		}
	}

	if len(found) == 0 {
		return nil, fmt.Errorf("no such formula")
	}

	return found, nil
}
