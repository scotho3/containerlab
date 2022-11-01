// Copyright 2020 Nokia
// Licensed under the BSD 3-Clause License.
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import "github.com/spf13/cobra"

var redeployCmd = &cobra.Command{
	Use:     "redeploy",
	Short:   "redeploy a lab",
	Long:    "redeploy a lab based defined by means of the topology definition file\nreference: https://containerlab.dev/cmd/redeploy/",
	Aliases: []string{"red"},
	PreRunE: sudoCheck,
	RunE:    redeployFn,
}

// init registers the deploy command to the root command.

func init() {
	rootCmd.AddCommand(redeployCmd)
}

// redeployFn destroys and then deploys a lab
// it is a wrapper around destroy and deploy commands

func redeployFn(cmd *cobra.Command, args []string) error {
	var err error

	destroyCmd.RunE(cmd, args)
	deployCmd.RunE(cmd, args)

	return err
}
