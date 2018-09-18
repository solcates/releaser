package collector

import (
	"github.com/davecgh/go-spew/spew"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

var testCollector *GitCollector
var testAuth transport.AuthMethod

func setupTestCase(t *testing.T) func(t *testing.T) {
	var err error
	testCollector = &GitCollector{}
	hd, err := homedir.Dir()
	sshPem := filepath.Join(hd, ".ssh", "id_rsa")
	if testAuth, err = ssh.NewPublicKeysFromFile("git", sshPem, ""); err != nil {
		t.Fatal(err)
	}
	t.Log("setup test case")
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("setup sub test")
	return func(t *testing.T) {
		t.Log("teardown sub test")
	}
}
func TestGitCollector_Collect(t *testing.T) {
	teardown := setupTestCase(t)
	defer teardown(t)
	tests := []struct {
		name     string
		gc       *GitCollector
		wantTags []*Tag
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "ok",
			gc: &GitCollector{
				reponame: "releaser",
				path:     "/Users/scates/go/src/github.com/solcates/releaser",
			},
			wantTags: []*Tag{

			},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTags, err := tt.gc.Collect()
			if (err != nil) != tt.wantErr {
				t.Errorf("GitCollector.Collect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTags, tt.wantTags) {
				t.Errorf("GitCollector.Collect() = %v, want %v", gotTags, tt.wantTags)
				spew.Dump(gotTags)
			}
		})
	}
}

func TestNewGitCollector(t *testing.T) {
	type args struct {
		reponame string
		path     string
	}
	localref, _ := filepath.Abs("../..")
	tests := []struct {
		name    string
		args    args
		wantGc  *GitCollector
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				reponame: "releaser",
				path:     "../..",
			},
			wantGc: &GitCollector{
				reponame: "releaser",
				path:     localref,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGc, err := NewGitCollector(tt.args.reponame, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGitCollector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGc, tt.wantGc) {
				t.Errorf("NewGitCollector() = %v, want %v", gotGc, tt.wantGc)
			}
		})
	}
}
