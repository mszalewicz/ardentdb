package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"os"
)

func Save(path string, data []byte) error {
	updatedFile := fmt.Sprintf("%s.tmp.%d", path, rand.Int64N(math.MaxInt64))

	fp, err := os.OpenFile(updatedFile, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)

	if err != nil {
		return err
	}

	defer func() {
		fp.Close()

		// remove updatedFile if any error encountered during whole process
		if err != nil {
			os.Remove(updatedFile)
		}
	}()

	if _, err = fp.Write(data); err != nil {
		return err
	}

	if err = fp.Sync(); err != nil {
		return err
	}

	err = os.Rename(path, updatedFile)
	return err
}
