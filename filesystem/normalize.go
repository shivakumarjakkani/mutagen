package filesystem

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func tildeExpand(path string) (string, error) {
	// Only process relevant paths.
	if path == "" || path[0] != '~' {
		return path, nil
	}

	// If the second character isn't a path separator, then someone is probably
	// trying to do a ~username expansion, but we can't support that without CGO
	// to support user.Lookup.
	if len(path) > 1 && !os.IsPathSeparator(path[1]) {
		return "", errors.New("unable to perform alternate user lookup")
	}

	// Compute the path.
	return filepath.Join(homeDirectory, path[1:]), nil
}

func Normalize(path string) (string, error) {
	// Expand any leading tilde.
	path, err := tildeExpand(path)
	if err != nil {
		return "", errors.Wrap(err, "unable to perform tilde expansion")
	}

	// Convert to an absolute path. This will also invoke filepath.Clean.
	path, err = filepath.Abs(path)
	if err != nil {
		return "", errors.Wrap(err, "unable to compute absolute path")
	}

	// Success.
	return path, nil
}
