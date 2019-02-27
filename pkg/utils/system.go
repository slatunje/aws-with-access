// Copyright © 2018 Sylvester La-Tunje. All rights reserved.

package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const (
	ModeSecretConfig = 0755
)

// HomeDir returns the path to user's home directory
func HomeDir() string {
	if runtime.GOOS == "windows" { // Windows
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME") // *nix
}

// LoginPath returns the login path
func LoginPath() string {
	if runtime.GOOS == "windows" {
		path, err := exec.LookPath("/usr/bin/login")
		if err != nil {
			link := "https://docs.microsoft.com/en-us/windows/wsl/install-win10#install-the-windows-subsystem-for-linux"
			panic(errors.New(fmt.Sprintf("%s. visit %s to fix this problem ", err, link)))
		}
		return path // Windows
	}
	return "/usr/bin/login" // *nix
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
	if err := os.MkdirAll(dir, os.FileMode(ModeSecretConfig)); err != nil {
		log.Fatalln(err)
	}
	return
}

// GetPathToFilename returns the current executed path
func GetPathToFilename() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
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
