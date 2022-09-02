package v

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
	"fmt"
	"github.com/briancsparks/activedevelopment/ad"
)

func Printf(format string, a ...any) (n int, err error) {
	if !shouldV() {
		return 0, nil
	}

	return fmt.Printf(format, a...)
}

func Print(a ...any) (n int, err error) {
	if !shouldV() {
		return 0, nil
	}

	return fmt.Print(a...)
}

func Println(a ...any) (n int, err error) {
	if !shouldV() {
		return 0, nil
	}

	return fmt.Println(a...)
}

func shouldV() bool {
	return ad.ConfigVerbosity() >= 1
}
