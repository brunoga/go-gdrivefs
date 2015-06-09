package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/brunoga/go-gdrivefs/filesystem"
	"github.com/brunoga/go-gdrivefs/gdrive"
	"github.com/brunoga/go-gdrivefs/settings"
	"github.com/hanwen/go-fuse/fuse"
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

	h, err := gdrive.NewHandler(a)
	if err != nil {
		log.Fatal(err)
	}

	gDriveNodeOpts := nodefs.NewOptions()

	gDriveMountOpts := &fuse.MountOptions{
		Options: []string{
			"max_read=131072",
		},
		MaxWrite: 131072,
		Name:     "GDrive",
	}

	gDriveNode := filesystem.NewRootNode(h)

	conn := nodefs.NewFileSystemConnector(gDriveNode, gDriveNodeOpts)

	state, err := fuse.NewServer(conn.RawFS(), flag.Arg(0), gDriveMountOpts)
	if err != nil {
		log.Fatalf("Failed to mount GDriveFS : %s", err)
	}

	state.Serve()
}
