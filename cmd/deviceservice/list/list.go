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

package list

import (
	"encoding/json"
	"fmt"
	"io"
	"text/tabwriter"
	"time"

	models "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	client "github.com/edgexfoundry-holding/edgex-cli/pkg"
	"github.com/edgexfoundry-holding/edgex-cli/pkg/utils"
)

type deviceServiceList struct {
	rd []models.DeviceService
}

// NewCommand returns the list device services command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists existing devices services",
		Long:  `Return the list fo current device services.`,
		Run: func(cmd *cobra.Command, args []string) {

			data, err := client.GetAllItems("deviceservice", "48081")

			if err != nil {
				fmt.Println(err)
				return
			}

			deviceServiceList1 := deviceServiceList{}

			errjson := json.Unmarshal(data, &deviceServiceList1.rd)
			if errjson != nil {
				fmt.Println(errjson)
			}

			pw := viper.Get("writer").(io.WriteCloser)
			w := new(tabwriter.Writer)
			w.Init(pw, 0, 8, 1, '\t', 0)
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", "Service ID", "Service Name", "Created", "Operating State")
			for _, device := range deviceServiceList1.rd {
				tCreated := time.Unix(device.Created/1000, 0)
				fmt.Fprintf(w, "%s\t%s\t%v\t%v\t\n",
					device.Id,
					device.Name,
					utils.HumanDuration(time.Since(tCreated)),
					device.OperatingState,
				)
			}
			w.Flush()
		},
	}
	return cmd
}
