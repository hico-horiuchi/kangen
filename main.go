package main

import (
	"./kangen"
	"fmt"
	"github.com/VividCortex/godaemon"
	"github.com/spf13/cobra"
)

var version string

func main() {
	var (
		expire string
		port   int
	)

	rootCmd := &cobra.Command{
		Use:   "kangen",
		Short: "URL shortening tool by golang",
		Long:  "URL shortening tool by golang\nhttps://github.com/hico-horiuchi/kangen",
	}

	addCmd := &cobra.Command{
		Use:   "add [shorten] [url]",
		Short: "Add shorten URL",
		Long:  "Add shorten URL",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 2:
				fmt.Print(kangen.Add(args[0], args[1], expire))
			}
		},
	}
	addCmd.Flags().StringVarP(&expire, "expire", "e", "", "Set timeout of shorten like 15m, 1h, 1d")
	rootCmd.AddCommand(addCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "remove [shorten]",
		Short: "Remove shorten URL",
		Long:  "Remove shorten URL",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 1:
				fmt.Print(kangen.Remove(args[0]))
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "Show all pairs of shorten and URL",
		Long:  "Show all pairs of shorten and URL",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(kangen.List())
		},
	})

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Start kangen server (http daemon)",
		Long:  "Start hangen server (http daemon)",
		Run: func(cmd *cobra.Command, args []string) {
			godaemon.MakeDaemon(&godaemon.DaemonAttr{})
			kangen.Server(port)
		},
	}
	serverCmd.Flags().IntVarP(&port, "port", "p", 5252, "Runs kangen on specified port")
	rootCmd.AddCommand(serverCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print kangen version",
		Long:  "Print kangen version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("kangen version", version)
		},
	})

	rootCmd.Execute()
}
