package utils

import (
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-version"
	"testing"
)

func Test_sortVerions(t *testing.T) {
	type args struct {
		in []*version.Version
	}
	a, _ := version.NewVersion("0.7.1")
	b, _ := version.NewVersion("1.1.0")
	c, _ := version.NewVersion("1.4.0-beta")
	d, _ := version.NewVersion("1.4.0")
	e, _ := version.NewVersion("2.0.0")
	tests := []struct {
		name    string
		args    args
		wantOut []*version.Version
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				in: []*version.Version{b, c, d, e, a,},
			},
			wantOut: []*version.Version{a, b, c, d, e,},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := SortVerions(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("sortVerions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotOut, tt.wantOut) {
				t.Errorf("sortVerions() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
