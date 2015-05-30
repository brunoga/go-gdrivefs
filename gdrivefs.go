package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/brunoga/go-gdrivefs/filesystem"
	"github.com/brunoga/go-gdrivefs/gdrive"
	"github.com/brunoga/go-gdrivefs/settings"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

var settingsPath = flag.String("settings_path", filepath.Join(os.Getenv("HOME"),
	".go-gdrivefs"), "Path to settings file.")

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("Usage : ", os.Args[0], " mountpoint")
	}

	s, err := settings.New(*settingsPath)
	if err != nil {
		log.Fatal(err)
	}

	a, err := gdrive.NewAuth(s)
	if err != nil {
		log.Fatal(err)
	}

	_, err = gdrive.NewHandler(a)
	if err != nil {
		log.Fatal(err)
	}

	gDriveOpts := nodefs.NewOptions()
	gDriveNode := filesystem.NewLoggingNode(true)

	state, _, err := nodefs.MountRoot(flag.Arg(0), gDriveNode, gDriveOpts)
	if err != nil {
		log.Fatalf("Failed to mount GDriveFS : %s", err)
	}

	state.Serve()
}
