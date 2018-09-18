package collector

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

type Tag struct {
	version string
	commits []*object.Commit
	hash    plumbing.Hash
}

type Collector interface {
	Collect() ([]*Tag, error)
}
