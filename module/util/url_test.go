package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveQueryParameter(t *testing.T) {
	type args struct {
		rawurl string
		keys   []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{rawurl: "http://127.0.0.1:7312/?a=1&b=2&c=3", keys: []string{"a", "b"}},
			want: "http://127.0.0.1:7312/?c=3",
		},
		{
			name: "case2",
			args: args{rawurl: "https://114514.com/?a=1&b=2&c=3", keys: []string{"a", "b", "c"}},
			want: "https://114514.com/",
		},
		{
			name: "mock token",
			args: args{rawurl: "https://114514.com/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.ey&sort=top",
				keys: []string{"token"}},
			want: "https://114514.com/?sort=top",
		},
		{
			name: "invalid query",
			args: args{rawurl: "https://114514.com/?a=1&b=2&c=3", keys: []string{"d"}},
			want: "https://114514.com/?a=1&b=2&c=3",
		},
		{
			name: "empty query",
			args: args{rawurl: "https://114514.com/?token=eyJhb", keys: []string{}},
			want: "https://114514.com/?token=eyJhb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RemoveQueryParameter(tt.args.rawurl, tt.args.keys...),
				"RemoveQueryParameter(%v, %v)", tt.args.rawurl, tt.args.keys)
		})
	}
}
