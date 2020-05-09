/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdShow() *cobra.Command {
	type Options struct {
		optint int
		optstr string
	}

	var (
		o = &Options{}
	)

	cmd := &cobra.Command{
		Use:   "show",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("show called: optint: %d, optstr: %s", o.optint, o.optstr)
		},
	}
	cmd.Flags().IntVarP(&o.optint, "int", "i", 0, "int option")
	cmd.Flags().StringVarP(&o.optstr, "str", "s", "default", "string option")

	return cmd
}

func init() {
}
