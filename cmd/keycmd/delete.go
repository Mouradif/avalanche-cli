// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package keycmd

import (
	"errors"
	"os"

	"github.com/ava-labs/avalanche-cli/pkg/ux"
	"github.com/spf13/cobra"
)

var forceDelete bool

// avalanche key delete
func newDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "delete",
		Short:        "Delete a signing key",
		Long:         "The key delete command deletes an existing signing key.",
		RunE:         deleteKey,
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
	}
	cmd.Flags().BoolVarP(
		&forceDelete,
		forceFlag,
		"f",
		false,
		"allow overwrites of existing keys with the same name",
	)
	return cmd
}

func deleteKey(cmd *cobra.Command, args []string) error {
	keyName := args[0]
	keyPath := app.GetKeyPath(keyName)

	// Check file exists
	_, err := os.Stat(keyPath)
	if err != nil {
		return errors.New("key does not exist")
	}

	if !forceDelete {
		confStr := "Are you sure you want to delete " + keyName + "?"
		conf, err := app.Prompt.CaptureNoYes(confStr)
		if err != nil {
			return err
		}

		if !conf {
			ux.Logger.PrintToUser("Delete cancelled")
			return nil
		}
	}

	// exists
	if err = os.Remove(keyPath); err != nil {
		return err
	}

	ux.Logger.PrintToUser("Key deleted")

	return nil
}