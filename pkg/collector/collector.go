package collector

import (
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

type Tag struct {
	version string
	commits []object.Commit
}

type Collector interface {
	Collect() ([]Tag, error)
}
