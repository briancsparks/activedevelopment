package cmd

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/briancsparks/activedevelopment/ad"
  afmt "github.com/briancsparks/activedevelopment/ad/fmt"
  "github.com/briancsparks/activedevelopment/ad/verbose/v"
  "github.com/briancsparks/activedevelopment/ad/verbose/vv"
  "github.com/briancsparks/activedevelopment/ad/verbose/vvv"
  "github.com/briancsparks/activedevelopment/ad/verbose/vvvx"
  "github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run and demo all features",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("all called")
    runAll()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runAll() {
  ad := ad.NewActiveDevelopment("all")
  //fmt.Printf("ad: %+v\n", ad)
  afmt.PrintVal(ad, "ad: %+v\n")

  showEach()

  fmt.Println("\nSetting AD mode")
  ad.SetActiveD(true)
  showEach()

  fmt.Println("\nSetting AD mode -- 2")
  ad.SetActiveD(true)
  ad.SetVerbosity(2)
  showEach()

  fmt.Println("\nSetting AD mode -- 3")
  ad.SetActiveD(true)
  ad.SetVerbosity(3)
  ad.SetDoX(false)
  showEach()

  fmt.Println("\nSetting AD mode -- 3x")
  ad.SetActiveD(true)
  ad.SetVerbosity(3)
  ad.SetDoX(true)
  showEach()
}

func showEach() {
  fmt.Println("Showing each:")
  a := 42
  v.PrintVal(a, "v    a: %+v\n")
  vv.PrintVal(a, "vv   a: %+v\n")
  vvv.PrintVal(a, "vvv  a: %+v\n")
  vvvx.PrintVal(a, "vvvx a: %+v\n")

  api()

  //fmt.Println("Done showing each:")
}

func api() {
  b := true
  if ad.LogApis() {
    vvvx.PrintVal(b, "api  b: %+v\n")
  }

}
