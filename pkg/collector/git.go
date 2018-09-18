package collector

import (
	"github.com/hashicorp/go-version"
	"github.com/pkg/errors"
	"github.com/solcates/releaser/pkg/utils"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
	"log"
	"path/filepath"
)

type GitCollector struct {
	reponame string
	path     string
	hash     string
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

func (gc *GitCollector) Collect() (tags map[string]*Tag, err error) {
	var repo *git.Repository
	if repo, err = git.PlainOpen(gc.path); err != nil {
		return
	}

	// Process Tags for their deltas between
	var tagsRef storer.ReferenceIter
	if tagsRef, err = repo.Tags(); err != nil {
		return
	}
	tags = make(map[string]*Tag)
	versions := []*version.Version{}
	err = tagsRef.ForEach(func(reference *plumbing.Reference) error {
		var vers *version.Version
		if vers, err = version.NewVersion(reference.Name().Short()); err != nil {
			return err
		}
		versions = append(versions, vers)
		t := &Tag{
			version: vers.String(),
			commits: []*object.Commit{},
		}
		tags[vers.String()] = t
		return nil
	})

	// Loop through the tags and find their commits, and not their friend
	if versions, err = utils.SortVerions(versions); err != nil {
		return
	}
	var previousTagHash string
	for _, version := range versions {
		tag := tags[version.String()]
		var commitRef object.CommitIter
		commitRef, err = repo.Log(&git.LogOptions{
			From:  tag.hash,
			Order: 0,
		})
		if err = commitRef.ForEach(func(commit *object.Commit) error {
			if commit.Hash.String() == previousTagHash {
				//log.Println("TESTSSTSTS")
				return errors.New("Test")
			}
			log.Println(tag.version, commit.Hash.String(), commit.Message)
			tag.commits = append(tag.commits, commit)
			return nil
		}); err != nil {
			return nil, err
		}
		previousTagHash = tag.hash.String()

	}

	return
}
