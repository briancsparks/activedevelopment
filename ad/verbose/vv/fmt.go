package vv

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
	"fmt"
	"github.com/briancsparks/activedevelopment/ad"
)

func Printf(format string, a ...any) {
	if !shouldVv() {
		return
	}

	fmt.Printf(format, a...)
}

func Print(a ...any) (n int, err error) {
	if !shouldVv() {
		return 0, nil
	}

	return fmt.Print(a...)
}

func Println(a ...any) (n int, err error) {
	if !shouldVv() {
		return 0, nil
	}

	return fmt.Println(a...)
}

func shouldVv() bool {
	return ad.ConfigVerbosity() >= 2
}
