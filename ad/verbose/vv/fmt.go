package vv

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "github.com/briancsparks/activedevelopment/ad"
  "log"
)

func Printf(format string, a ...any) {
	if !shouldVv() {
		return
	}

	log.Printf(format, a...)
}

func Print(a ...any) {
	if !shouldVv() {
		return
	}

	log.Print(a...)
}

func Println(a ...any) {
	if !shouldVv() {
		return
	}

	 log.Println(a...)
}

func PrintVal(a any, format string) {
  if !shouldVv() {
    return
  }

  log.Printf(format, a)
}

func shouldVv() bool {
	return ad.TheConfig.Verbosity() >= 2
}
