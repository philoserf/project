package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

type target struct {
	name string
	path string
}

func (t *target) Write() (int, error) {
	log.Printf("writing %s\n", t.name)

	data, err := Asset("data/" + t.name)
	if err != nil {
		log.Printf("failed to read %s, %s\f", t.name, err)
		return 0, errors.Wrap(err, "failed to read asset")
	}

	f, err := os.Create(t.path + t.name)
	if err != nil {
		log.Printf("failed to create %s, %s\f", t.name, err)
		return 0, errors.Wrap(err, "failed to create file")
	}
	defer func() {
		if e := f.Close(); e != nil {
			log.Printf("failed to close %s, %s\f", t.name, e)
		}
	}()

	n, err := f.Write(data)
	if err != nil {
		log.Printf("failed to write %s, %s\f", t.name, err)
		return 0, errors.Wrap(err, "failed to write file")
	}

	log.Printf("finished writing %s\n", t.name)
	return n, nil
}
