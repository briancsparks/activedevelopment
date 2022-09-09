package vvv

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "github.com/briancsparks/activedevelopment/ad"
  "log"
)

func Printf(format string, a ...any) {
	if !shouldVvv() {
    return
  }

  log.Printf(format, a...)
}

func Print(a ...any) {
	if !shouldVvv() {
		return
	}

  log.Print(a...)
}

func Println(a ...any)  {
	if !shouldVvv() {
		return
	}

  log.Println(a...)
}

func shouldVvv() bool {
	return ad.ConfigVerbosity() >= 3
}
