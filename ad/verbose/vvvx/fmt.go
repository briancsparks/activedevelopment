package vvvx

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "github.com/briancsparks/activedevelopment/ad"
  "log"
)

func Printf(format string, a ...any) {
	if !shouldVvvx() {
		return
	}

  log.Printf(format, a...)
}

func Print(a ...any) {
	if !shouldVvvx() {
		return
	}

  log.Print(a...)
}

func Println(a ...any) {
	if !shouldVvvx() {
		return
	}

  log.Println(a...)
}

func PrintVal(a any, format string) {
  if !shouldVvvx() {
    return
  }

  log.Printf(format, a)
}

func shouldVvvx() bool {
	return (ad.TheConfig.Verbosity() >= 3) && ad.DoX()
}

