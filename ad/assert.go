package ad

import "fmt"

// TODO: For all below, allow Panic or Fatal or whatever

// -------------------------------------------------------------------------------------------------------------------

func Breakout(msg string, quiet bool) {

	// BBB: Breakpoint here when bad things happen
	if !quiet {
		fmt.Printf("  ---------- BREAKOUT!! %v !!\n", msg)
	}
}

// -------------------------------------------------------------------------------------------------------------------

func assert(test bool) {
	if !test {
		Breakout("", false)
	}
}

// -------------------------------------------------------------------------------------------------------------------

func asserter(test bool) bool {
	if !test {
		Breakout("", true)
	}
	return !test
}

// -------------------------------------------------------------------------------------------------------------------

func assertMsg(test bool, msg string) {
	if !test {
		Breakout(msg, false)
	}
}

// -------------------------------------------------------------------------------------------------------------------
