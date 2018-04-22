package utils

import (
	"runtime"
	"os"
	"path/filepath"
)

// UserHomeDir returns the path to user's home directory
func UserHomeDir() string {
	if runtime.GOOS == "windows" { // Windows
		return os.Getenv("USERPROFILE")
	}

	// *nix
	return os.Getenv("HOME")
}

// GetPathToFilename returns the current executed path
func GetPathToFilename() (string) {
	_, filename, _, _ := runtime.Caller(1)
	return  filename
}

// GetPathFromFilename returns the full path from filename `string`
func GetPathFromFilename(filename string) string {
	return filepath.Join(filepath.Dir(GetPathToFilename()), filename)
}
