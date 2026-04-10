package flags

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func writeFlagKeysFile(t *testing.T, contents string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "flag-keys.txt")
	require.NoError(t, os.WriteFile(path, []byte(contents), 0600))
	return path
}

func Test_getFlagKeysFromFile(t *testing.T) {
	tests := []struct {
		name     string
		contents string
		want     map[string][]string
	}{
		{
			name:     "reads one key per line",
			contents: "flag-one\nflag-two\n",
			want:     map[string][]string{offlineProjectKey: {"flag-one", "flag-two"}},
		},
		{
			name:     "skips comments, blank lines, and surrounding whitespace",
			contents: "# header comment\n\n  flag-one  \n\t\n# another\nflag-two",
			want:     map[string][]string{offlineProjectKey: {"flag-one", "flag-two"}},
		},
		{
			name:     "filters keys shorter than the minimum length",
			contents: "ab\nflag-one\n",
			want:     map[string][]string{offlineProjectKey: {"flag-one"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFlagKeysFromFile(writeFlagKeysFile(t, tt.contents))
			require.Equal(t, tt.want, got)
		})
	}
}
