package completion

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestNewCompletionCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "kptn"}
	completionCmd := NewCompletionCommand(rootCmd)

	tests := []struct {
		shell       string
		expectedErr bool
		expectedOut string
	}{
		{"bash", false, "# bash completion"},
		{"zsh", false, "#compdef kptn"},
		{"fish", false, "# fish completion"},
		{"powershell", false, "# powershell completion"},
		{"unsupported", true, "Unsupported shell: unsupported"},
	}

	for _, tt := range tests {
		t.Run(tt.shell, func(t *testing.T) {
			buf := new(bytes.Buffer)
			completionCmd.SetOut(buf)
			completionCmd.SetArgs([]string{tt.shell})

			err := completionCmd.Execute()
			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				output := buf.String()
				require.Contains(t, output, tt.expectedOut)
			}
		})
	}
}

func TestZshCompletion(t *testing.T) {

	rootCmd := &cobra.Command{Use: "kptn"}
	completionCmd := NewCompletionCommand(rootCmd)

	// Create a buffer to capture the output
	buf := new(bytes.Buffer)
	completionCmd.SetOut(buf)
	completionCmd.SetArgs([]string{"zsh"})

	// Execute the command
	err := completionCmd.Execute()
	require.NoError(t, err)

	// Check that the output contains the expected Zsh completion script
	output := buf.String()
	require.Contains(t, output, "#compdef kptn")
	require.Contains(t, output, "completion")
}
