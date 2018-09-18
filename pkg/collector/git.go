package collector

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
	"path/filepath"
)

type GitCollector struct {
	reponame string
	path     string
}

func NewGetCollectorForSSH(reponame, path string) (gc *GitCollector, err error) {

	path, err = filepath.Abs(path)
	gc = &GitCollector{
		reponame: reponame,
		path:     path,
	}
	return
}

func (gc *GitCollector) Collect() (tags []Tag, err error) {
	var repo *git.Repository
	if repo, err = git.PlainOpen(gc.path); err != nil {
		return
	}

	var tagsRef storer.ReferenceIter
	if tagsRef, err = repo.Tags(); err != nil {
		return
	}
	err = tagsRef.ForEach(func(reference *plumbing.Reference) error {
		fmt.Println(reference.Name())
		return nil
	})
	return
}
