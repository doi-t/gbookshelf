// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/doi-t/gbookshelf/pkg/apis/gbookshelf"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates status of a book in your bookshelf",
	Long: `Updates status of a book in your bookshelf. 
	
Currently this command tries to find a book of given title in bookshelf and update status.
Note that 'status' becomes 'incomplete' by default.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		page, err := cmd.Flags().GetInt32("page")
		if err != nil {
			return err
		}
		current, err := cmd.Flags().GetInt32("current")
		if err != nil {
			return err
		}
		status, err := cmd.Flags().GetString("status")
		if err != nil {
			return err
		}
		var s bool
		switch status {
		case "incomplete":
			s = false
		case "done":
			s = true
		default:
			return fmt.Errorf("'%s' is invalid status. 'status' must be either 'incomplete' or 'done'.", status)
		}

		return update(context.Background(), strings.Join(args, " "), page, s, current)
	},
}

func init() {
	updateCmd.Flags().Int32P("page", "p", -1, "Update the number of pages of the book that you specified.")
	updateCmd.Flags().StringP("status", "s", "incomplete", "Give 'done' if you finally finish to read the book!")
	updateCmd.Flags().Int32P("current", "c", -1, "Update the current page position of the book you specified..")

	rootCmd.AddCommand(updateCmd)
}

func update(ctx context.Context, title string, page int32, status bool, current int32) error {
	new := &gbookshelf.Book{Title: title, Page: page, Done: status}
	old, err := client.Update(ctx, &gbookshelf.Book{Title: title, Page: page, Done: status, Current: current})
	if err != nil {
		return fmt.Errorf("could not send a book status to the backend: %v", err)
	}

	fmt.Printf("book updated successfully: %v-> %v\n", old, new)
	return nil
}
