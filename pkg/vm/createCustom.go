// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package vm

import (
	"fmt"
	"os"

	"github.com/ava-labs/avalanche-cli/pkg/application"
	"github.com/ava-labs/avalanche-cli/pkg/models"
	"github.com/ava-labs/avalanche-cli/pkg/ux"
)

func CreateCustomSubnetConfig(app *application.Avalanche, subnetName string, genesisPath, vmPath string) ([]byte, *models.Sidecar, error) {
	ux.Logger.PrintToUser("creating custom VM subnet %s", subnetName)

	genesisBytes, err := loadCustomGenesis(app, genesisPath)
	if err != nil {
		return nil, &models.Sidecar{}, err
	}

	if vmPath == "" {
		vmPath, err = app.Prompt.CaptureExistingFilepath("Enter path to vm binary")
		if err != nil {
			return nil, &models.Sidecar{}, err
		}
	}

	rpcVersion, err := GetVMBinaryProtocolVersion(vmPath)
	if err != nil {
		return nil, &models.Sidecar{}, fmt.Errorf("unable to get RPC version: %w", err)
	}

	sc := &models.Sidecar{
		Name:       subnetName,
		VM:         models.CustomVM,
		VMVersion:  "",
		RPCVersion: rpcVersion,
		Subnet:     subnetName,
		TokenName:  "",
	}

	return genesisBytes, sc, app.CopyVMBinary(vmPath, subnetName)
}

func loadCustomGenesis(app *application.Avalanche, genesisPath string) ([]byte, error) {
	var err error
	if genesisPath == "" {
		genesisPath, err = app.Prompt.CaptureExistingFilepath("Enter path to custom genesis")
		if err != nil {
			return nil, err
		}
	}

	genesisBytes, err := os.ReadFile(genesisPath)
	return genesisBytes, err
}
