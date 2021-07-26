package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		path1   string
		path2   string
		wantErr bool
		err     string
	}{
		{
			name:    "Success",
			url:     "https://api.github.com/repos",
			path1:   "spf13",
			path2:   "cobra",
			wantErr: false,
		},
		{
			name:    "Wrong url call",
			url:     "https://api.gitb.com/repos",
			path1:   "spf13",
			path2:   "cobra",
			wantErr: true,
			err:     "Get \"https://api.gitb.com/repos/spf13/cobra/commits\": dial tcp: lookup api.gitb.com: no such host",
		},
		{
			name:    "Inalid path 1",
			url:     "https://api.github.com/repos",
			path1:   "dsdsa",
			path2:   "cobra",
			wantErr: true,
			err:     "json: cannot unmarshal object into Go value of type []api.Response",
		},
		{
			name:    "Inalid path 2",
			url:     "https://api.github.com/repos",
			path1:   "spf13",
			path2:   "jhcdgcjgd",
			wantErr: true,
			err:     "json: cannot unmarshal object into Go value of type []api.Response",
		},
	}
	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			ti, err := Call(tc.url, tc.path1, tc.path2)
			if tc.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, ti)
				assert.EqualError(t, err, tc.err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, ti)
			}
		})
	}
}
