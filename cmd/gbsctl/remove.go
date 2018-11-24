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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a book from your bookshelf",
	RunE: func(cmd *cobra.Command, args []string) error {
		return remove(context.Background(), strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func remove(ctx context.Context, title string) error {
	_, err := client.Remove(ctx, &gbookshelf.Title{Title: title})
	if err != nil {
		return fmt.Errorf("could not remove a book: %v", err)
	}
	fmt.Println("Remove book successfully")
	return nil
}
