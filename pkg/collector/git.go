package collector

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
	"path/filepath"
)

type GitCollector struct {
	reponame string
	path     string
}

//NewGitCollector returns a GitCollector for a given path (assumed-cloned)
func NewGitCollector(reponame, path string) (gc *GitCollector, err error) {
	path, err = filepath.Abs(path)
	gc = &GitCollector{
		reponame: reponame,
		path:     path,
	}
	return
}

func (gc *GitCollector) Collect() (tags []*Tag, err error) {
	var repo *git.Repository
	if repo, err = git.PlainOpen(gc.path); err != nil {
		return
	}

	var tagsRef storer.ReferenceIter
	if tagsRef, err = repo.Tags(); err != nil {
		return
	}

	// build channels
	var commitRef object.CommitIter
	err = tagsRef.ForEach(func(reference *plumbing.Reference) error {
		commitRef, err = repo.Log(&git.LogOptions{
			From:  reference.Hash(),
			Order: 0,
		})

		t := &Tag{
			version: reference.Name().Short(),
			commits: []*object.Commit{},
		}
		if err = commitRef.ForEach(func(commit *object.Commit) error {
			t.commits = append(t.commits, commit)
			return nil
		}); err != nil {
			return err
		}
		tags = append(tags, t)
		return nil
	})

	return
}
