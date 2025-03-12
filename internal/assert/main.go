package assert

import "log"

func Assert(condition bool, info string) {
	if !condition {
		log.Fatal(info)
	}
}