/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
  "github.com/briancsparks/activedevelopment/ad"
  "github.com/hypebeast/go-osc/osc"
  "log"

  "github.com/spf13/cobra"
)

// oscCmd represents the osc command
var oscCmd = &cobra.Command{
	Use:   "osc",
	Short: "Use OSC to manage dynamic variables",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("osc called")
    runOsc()
	},
}

func init() {
	rootCmd.AddCommand(oscCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oscCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oscCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runOsc() {
  var ok bool
  var var1 int32 = 42
  var var1name string

  ad := ad.NewActiveDevelopmentWithFeedback("adosc")
  var1name, err := ad.App.RegisterVariable("var1", func(msg *osc.Message) {
    var1, ok = msg.Arguments[0].(int32)
    fmt.Println("var1:", var1, ok)
    message := osc.NewMessage(var1name)
    message.Append(int32(var1))
    _= ad.App.Feedback(message)
  })
  if err != nil {
    log.Fatal(err)
  }

  _, err = ad.App.RegisterDefaultHandler(func(msg *osc.Message) {
    fmt.Println(msg)
  })
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Listening...")
  err = ad.App.StartBlocking("127.0.0.1", 54321)
  if err != nil {
    log.Fatal(err)
  }

}
