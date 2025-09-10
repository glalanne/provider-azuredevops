/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alecthomas/kingpin/v2"
	"github.com/crossplane/upjet/v2/pkg/pipeline"
	"github.com/glalanne/provider-azuredevops/config"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}

	sdkProvider := azuredevops.Provider()

	pc, err := config.GetProvider(context.Background(), sdkProvider, true)
	kingpin.FatalIfError(err, "Cannot initialize the cluster-scoped provider configuration")

	pns, err := config.GetProviderNamespaced(context.Background(), sdkProvider, true)
	kingpin.FatalIfError(err, "Cannot initialize the namespaced provider configuration")

	pipeline.Run(pc, pns, absRootDir)
}
