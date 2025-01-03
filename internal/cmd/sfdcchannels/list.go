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

package sfdcchannels

import (
	"internal/apiclient"
	"internal/client/sfdc"
	"internal/clilog"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// ListCmd to get integration flow
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List sfdcchannels in Application Integration",
	Long:  "List sfdcchannels in Application Integration",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		project := cmd.Flag("proj").Value.String()
		region := cmd.Flag("reg").Value.String()

		if err = apiclient.SetRegion(region); err != nil {
			return err
		}
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			clilog.Debug.Printf("%s: %s\n", f.Name, f.Value)
		})
		return apiclient.SetProjectID(project)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		instance := cmd.Flag("instance").Value.String()
		_, err = sfdc.ListChannels(instance)
		return
	},
}

func init() {
	var instance string

	ListCmd.Flags().StringVarP(&instance, "instance", "i",
		"", "sfdc instance name")

	_ = ListCmd.MarkFlagRequired("instance")
}
