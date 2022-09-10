/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
  "github.com/hypebeast/go-osc/osc"
  "log"
  "os"
  "strconv"

  "github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
    fmt.Printf("args: %+v\n", args)
    runOscClient(args)
	},
}

func init() {
	oscCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runOscClient(args []string) {

  fmt.Printf("os.Args: %+v\n", os.Args)
  numArgs := len(args)

  if numArgs != 1 {
    printOscClientUsage()
    os.Exit(1)
  }

  port, err := strconv.ParseInt(args[0], 10, 32)
  if err != nil {
    printOscClientUsage()
    log.Fatal(err)
  }

  ip := "localhost"
  client := osc.NewClient(ip, int(port))

  message := osc.NewMessage("/adosc/var1")
  message.Append(int32(99))

  fmt.Printf("sending message: %+v\n", message)

  err = client.Send(message)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Got err: %+v\n", err)

}


func printOscClientUsage() {
  fmt.Printf("Usage: %s PORT\n", os.Args[0])
}

