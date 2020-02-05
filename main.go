package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	file "github.com/evzpav/filecli/file"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "filecli",
	Short: "Files read and write",
}

func main() {

	fileApp, err := file.New()
	if err != nil {
		log.Fatal(err)
	}

	readCmd := buildReadCommand(fileApp)
	writeCmd := buildWriteCommand(fileApp)

	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(writeCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func buildReadCommand(app *file.File) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "read",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filepath := args[0]
			content, err := app.ReadFile(filepath)
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}
			fmt.Printf("Content: %s\n", content)
		},
	}

	return cmd
}

func buildWriteCommand(app *file.File) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "write",
		Args: cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filepath := args[0]
			reader := bufio.NewReader(os.Stdin)
			
			fmt.Println("Type here the file content:")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}
			content, err := app.WriteFile(filepath, text)
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}
			fmt.Printf("Content: %s\n", content)
		},
	}

	return cmd

}
