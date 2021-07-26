package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	var rescueStdout, r, w *os.File
	tests := []struct {
		name       string
		setup      func()
		errMsg     string
		successMsg string
		wantErr    bool
	}{
		{
			name: "Success",
			setup: func() {
				rescueStdout = os.Stdout
				r, w, _ = os.Pipe()
				os.Stdout = w
				main()
			},
			successMsg: "publishing message Hello Everyone serial number 1",
			wantErr:    true,
		},
		{
			name: "Fail wrong config file",
			setup: func() {
				rescueStdout = os.Stdout
				r, w, _ = os.Pipe()
				os.Stdout = w
				fileName = "../condsdfig.yml"
				main()
			},
			errMsg:  "Unable to read configs open ../condsdfig.yml: no such file or directory",
			wantErr: true,
		},
		{
			name: "Wrong emitter host port",
			setup: func() {
				rescueStdout = os.Stdout
				r, w, _ = os.Pipe()
				os.Stdout = w
				fileName = "../config.yml"
				emitterHostPort = "tcp://locaost:9000"
				main()
			},
			errMsg:  "wrong emitter address",
			wantErr: true,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			tc.setup()
			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = rescueStdout

			if tc.wantErr {
				assert.Contains(t, string(out), tc.errMsg)
			} else {
				assert.Contains(t, string(out), tc.successMsg)
			}
		})
	}

}
