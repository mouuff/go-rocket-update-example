package main

import (
	"log"
	"runtime"

	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
)

func main() {
	u := &updater.Updater{
		Provider: &provider.Github{
			RepositoryURL: "github.com/mouuff/go-rocket-update-example",
			ZipName:       "binaries_" + runtime.GOOS + ".zip",
		},
		ExecutableName: "go-rocket-update-example",
		Version:        "v1.2.0", // You can change this value to trigger an update
	}
	log.Println(u.Version)
	err := u.Update()
	if err != nil {
		log.Fatal(err)
	}
}
