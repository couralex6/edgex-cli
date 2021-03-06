// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rm

import (
	"fmt"

	client "github.com/edgexfoundry-holding/edgex-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCommand returns the rm device service command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm [device service name|ID]",
		Short: "Removes device service by name or ID",
		Long:  `Removes a device service from the core-metadata DB.`,
		Run: func(cmd *cobra.Command, args []string) {

			// Checking for args
			if len(args) == 0 {
				fmt.Printf("Error: No device service ID/Name provided.\n")
				return
			}

			deviceID := args[0]
			respBody, err := client.DeleteItem(deviceID,
				"deviceservice/id/",
				"deviceservice/name/",
				"48081")

			if err != nil {
				fmt.Println(err)
				return
			}

			// Display Results
			if string(respBody) == "true" {
				fmt.Printf("Removed: %s\n", deviceID)
			} else {
				fmt.Printf("Remove Unsuccessful: %s\n", string(respBody))
			}
		},
	}
	return cmd
}
