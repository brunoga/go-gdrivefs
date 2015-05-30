package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/brunoga/go-gdrivefs/gdrive"
	"github.com/brunoga/go-gdrivefs/settings"
)

var settingsPath = flag.String("settings_path", filepath.Join(os.Getenv("HOME"),
	".go-gdrivefs"), "Path to settings file.")

func main() {
	flag.Parse()

	s, err := settings.New(*settingsPath)
	if err != nil {
		panic(err)
	}

	_, err = gdrive.NewAuth(s)
	if err != nil {
		panic(err)
	}
}
