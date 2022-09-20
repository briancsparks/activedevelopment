package ad

import "log"

func MustNot(err error) {
  assert(err != nil)

  if err != nil {
    log.Panic(err)
  }
}



