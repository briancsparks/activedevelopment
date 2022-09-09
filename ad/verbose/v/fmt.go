package v

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "github.com/briancsparks/activedevelopment/ad"
  "log"
)

func Printf(format string, a ...any) {
	if !shouldV() {
		return
	}

  log.Printf(format, a...)
}

func Print(a ...any) {
	if !shouldV() {
		return
	}

  log.Print(a...)
}

func Println(a ...any) {
	if !shouldV() {
		return
	}

  log.Println(a...)
}

func shouldV() bool {
	return ad.ConfigVerbosity() >= 1
}
