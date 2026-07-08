package options

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_IsOffline(t *testing.T) {
	require.False(t, Options{}.IsOffline())
	require.True(t, Options{FlagKeysFile: "flag-keys.txt"}.IsOffline())
}

func Test_ValidateRequired_offline(t *testing.T) {
	t.Run("only dir is required", func(t *testing.T) {
		opts := Options{FlagKeysFile: "flag-keys.txt", Dir: "."}
		require.NoError(t, opts.ValidateRequired())
	})

	t.Run("missing dir fails", func(t *testing.T) {
		opts := Options{FlagKeysFile: "flag-keys.txt"}
		require.EqualError(t, opts.ValidateRequired(), "missing required option(s): [dir]")
	})
}

func Test_Validate_offline(t *testing.T) {
	dir := t.TempDir()

	t.Run("missing flag keys file fails", func(t *testing.T) {
		opts := Options{FlagKeysFile: filepath.Join(dir, "missing.txt"), Dir: dir}
		require.ErrorContains(t, opts.Validate(), `invalid value for "flagKeysFile"`)
	})

	t.Run("existing flag keys file passes without repo or access token options", func(t *testing.T) {
		path := filepath.Join(dir, "flag-keys.txt")
		require.NoError(t, os.WriteFile(path, []byte("flag-one\n"), 0600))

		opts := Options{FlagKeysFile: path, Dir: dir}
		require.NoError(t, opts.Validate())
	})
}
