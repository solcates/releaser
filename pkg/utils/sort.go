package utils

import (
	"github.com/hashicorp/go-version"
	"sort"
)

func SortVerions(in []*version.Version) (out []*version.Version, err error) {
	out = make([]*version.Version, len(in))
	for i, r := range in {

		out[i] = r
	}

	sort.Sort(version.Collection(out))
	return
}
