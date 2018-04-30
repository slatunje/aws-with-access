package utils

import (
	"runtime"
	"os"
	"path/filepath"
)

const (
	ModeSecretConfig = 0755
)

// HomeDir returns the path to user's home directory
func HomeDir() string {
	if runtime.GOOS == "windows" { // Windows
		return os.Getenv("USERPROFILE")
	}

	// *nix
	return os.Getenv("HOME")
}

// ProjectDir returns the path to project directory
func ProjectDir(dir string) string {
	return GetPathFromFilename(dir)
}

// ProjectConfigDir returns the path to project config directory
func ProjectConfigDir(dir string) string {
	return GetPathFromFilename(dir)
}
// ProjectConfigDir returns the path to project config directory
func ProjectDefaultConfigDir(app string) (dir string) {
	dir = filepath.Join(HomeDir(), app)
	os.MkdirAll(dir, os.FileMode(ModeSecretConfig))
	return
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

// Exists check if the file exist
func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err != nil, err
}
