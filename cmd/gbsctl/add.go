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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new book on your bookshelf",
	Long:  `add a new book on your bookshelf. Give title and its number of pages.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		page, err := cmd.Flags().GetInt32("page")
		if err != nil {
			return err
		}
		return add(context.Background(), strings.Join(args, " "), page)
	},
}

func init() {
	addCmd.Flags().Int32P("page", "p", 0, "The number of pages of the book that you add to bookshelf.")

	rootCmd.AddCommand(addCmd)
}

func add(ctx context.Context, title string, page int32) error {
	_, err := client.Add(ctx, &gbookshelf.Book{Title: title, Page: page, Done: false, Current: 0})
	if err != nil {
		return fmt.Errorf("could not send a book to the backend: %v", err)
	}

	fmt.Println("book added successfully")
	return nil
}
