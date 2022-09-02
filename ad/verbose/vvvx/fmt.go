package vvvx

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
	"fmt"
	"github.com/briancsparks/activedevelopment/ad"
)

func Printf(format string, a ...any) (int, error) {
	if !shouldVvvx() {
		return 0, nil
	}

	return fmt.Printf(format, a...)
}

func Print(a ...any) (int, error) {
	if !shouldVvvx() {
		return 0, nil
	}

	return fmt.Print(a...)
}

func Println(a ...any) (int, error) {
	if !shouldVvvx() {
		return 0, nil
	}

	return fmt.Println(a...)
}

func shouldVvvx() bool {
	return (ad.ConfigVerbosity() >= 3) && ad.DoX()
}
