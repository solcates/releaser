package utils

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"sort"
)

func sortVerions(in []string) (out []*version.Version, err error) {
	out = make([]*version.Version, len(in))
	for i, r := range in {
		v, err := version.NewVersion(r)
		if err != nil {
			fmt.Errorf("Error parsing version: %s", err)
		}

		out[i] = v
	}

	sort.Sort(version.Collection(out))
	return
}
