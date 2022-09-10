package afmt

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "fmt"
  "log"
)

func Printf(format string, a ...any) {
  //return 0, nil
  fmt.Printf(format, a...)
}

func Print(a ...any) {
  fmt.Print(a...)
}

func Println(a ...any) {
  fmt.Println(a...)
}

func PrintVal(a any, format string) {
  log.Printf(format, a)
}

