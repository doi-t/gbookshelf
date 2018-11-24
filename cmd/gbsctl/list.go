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

	"github.com/doi-t/gbookshelf/pkg/apis/gbookshelf"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of books piled on your bookshelf",
	Long: `Lists all of books piled on your bookshelf including the following information:
- A book title
- How many pages a book has
- Done flag (It indicates whether you finish to read the corresponding book or not)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filsterDone, err := cmd.Flags().GetBool("incomplete_only")
		if err != nil {
			return err
		}
		return list(context.Background(), filsterDone)
	},
}

func init() {
	listCmd.Flags().BoolP("incomplete_only", "i", false, "Show only books that are still being stacked on your book pile.")

	rootCmd.AddCommand(listCmd)
}

func list(ctx context.Context, filterDone bool) error {
	l, err := client.List(ctx, &gbookshelf.Void{})
	if err != nil {
		return fmt.Errorf("cloud not fetch books: %v", err)
	}
	for _, b := range l.Books {
		if b.Done {
			if filterDone {
				continue
			}
			fmt.Printf("👍")
		} else {
			fmt.Printf("😱")
		}
		fmt.Printf(" %s (P%d)\n", b.Title, b.Page)
	}
	return nil
}
