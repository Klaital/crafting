package main

import (
	"flag"
	"github.com/klaital/crafting/pkg/recipes"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	libraryPath := flag.String("library", "library.json", "File containing the material library")
	flag.Parse()

	library, err := recipes.LoadLibraryFromFile(*libraryPath)

	if err != nil {
		log.WithField("LibraryPath", *libraryPath).Errorf("Failed to load library from disk: %v", err)
		os.Exit(1)
	}

	log.WithField("LibraryLength", len(library.Items)).Infof("Loaded library")
}