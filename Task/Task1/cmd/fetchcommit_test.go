package cmd

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

var rootCmd = &cobra.Command{}

func TestExecute(t *testing.T) {

	tests := []struct {
		name    string
		setup   *cobra.Command
		args    []string
		err     string
		wanterr bool
	}{
		{
			name:    "Success",
			setup:   &cobra.Command{},
			args:    []string{"spf13", "cobra"},
			wanterr: false,
		},
		{
			name: "Failure without owner args",
			setup: &cobra.Command{
				RunE: func(cmd *cobra.Command, args []string) error {
					return errors.New("repo arguement not passed")
				},
			},
			args:    []string{"owner", ""},
			err:     "repo arguement not passed",
			wanterr: true,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			cmd := tc.setup
			cmd.SetArgs(tc.args)

			if tc.wanterr {
				assert.EqualError(t, Execute(cmd), tc.err)
			} else {
				assert.Nil(t, Execute(cmd))
			}

		})
	}

}
