// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authconfigs

import (
	"os"

	"internal/apiclient"
	"internal/client/authconfigs"

	"github.com/spf13/cobra"
)

// PatchCmd to create a new connection
var PatchCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing authconfig",
	Long:  "Update an existing authconfig in a region",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		cmdProject := cmd.Flag("proj")
		cmdRegion := cmd.Flag("reg")

		if err = apiclient.SetRegion(cmdRegion.Value.String()); err != nil {
			return err
		}
		return apiclient.SetProjectID(cmdProject.Value.String())
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		name := cmd.Flag("name").Value.String()

		if _, err := os.Stat(authConfigFile); os.IsNotExist(err) {
			return err
		}

		content, err := os.ReadFile(authConfigFile)
		if err != nil {
			return err
		}

		_, err = authconfigs.Patch(name, content, updateMask)
		return err
	},
}

var updateMask []string

func init() {
	var name string

	PatchCmd.Flags().StringVarP(&name, "name", "n",
		"", "AuthConfig name")
	PatchCmd.Flags().StringVarP(&authConfigFile, "file", "f",
		"", "AuthConfig details JSON file path")
	PatchCmd.Flags().StringArrayVarP(&updateMask, "update-mask", "",
		nil, "Update mask: A list of comma separates values to update")

	_ = PatchCmd.MarkFlagRequired("updateMask")
}
