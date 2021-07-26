package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelper(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		err      string
		wantErr  bool
	}{
		{
			name:     "Success",
			fileName: "../config.yml",
			wantErr:  false,
		},
		{
			name:     "Invalid config file",
			fileName: "../coig.yml",
			wantErr:  true,
			err:      "open ../coig.yml: no such file or directory",
		},
		{
			name:     "Wrong config file",
			fileName: "../go.mod",
			wantErr:  true,
			err:      "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `module ...` into helper.Config",
		},
	}
	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			f, err := LoadConfig(tc.fileName)
			if tc.wantErr {
				assert.EqualError(t, err, tc.err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, f)
			}
		})
	}
}
