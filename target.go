package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type target struct {
	name string
	path string
}

func (t *target) write() (int, error) {
	fmt.Printf("writing %v\n", t.name)

	data, err := Asset("data/" + t.name)
	if err != nil {
		fmt.Printf("failed to read %v, %v\f", t.name, err)
		return 0, errors.Wrap(err, "failed to read asset")
	}

	f, err := os.Create(t.path + t.name)
	if err != nil {
		fmt.Printf("failed to create %v, %v\f", t.name, err)
		return 0, errors.Wrap(err, "failed to create file")
	}
	defer func() {
		if e := f.Close(); e != nil {
			fmt.Printf("failed to close %v, %v\f", t.name, e)
		}
	}()

	n, err := f.Write(data)
	if err != nil {
		fmt.Printf("failed to write %v, %v\f", t.name, err)
		return 0, errors.Wrap(err, "failed to write file")
	}

	fmt.Print("...finished\n")
	return n, nil
}
