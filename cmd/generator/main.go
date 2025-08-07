/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/crossplane/upjet/pkg/pipeline"
	tfvaultschema "github.com/hashicorp/terraform-provider-vault/schema"
	tfvault "github.com/hashicorp/terraform-provider-vault/vault"

	"github.com/upbound/provider-vault/config"
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
	sdkProvider := tfvaultschema.NewProvider(tfvault.Provider()).SchemaProvider()
	pc, err := config.GetProvider(context.Background(), sdkProvider, true)
	if err != nil {
		panic(fmt.Sprintf("cannot get cluster provider configuration: %v", err))
	}
	pns, err := config.GetProviderNamespaced(context.Background(), sdkProvider, true)
	if err != nil {
		panic(fmt.Sprintf("cannot get namespaced provider configuration: %v", err))
	}
	pipeline.Run(pc, pns, absRootDir)
}
