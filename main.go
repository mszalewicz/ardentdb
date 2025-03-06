package main

import (
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"os"
	"path/filepath"
)

// Save new file through new file creation and swapping with original after updates process is finished
func Save(path string, data []byte) error {
	updatedFile := fmt.Sprintf("%s.tmp.%d", path, rand.Int64N(math.MaxInt64))

	fp, err := os.OpenFile(updatedFile, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)

	if err != nil {
		return err
	}

	defer func() {
		closeErr := fp.Close()

		if closeErr != nil {
			log.Fatal(closeErr)
		}

		// remove temporary file if any error encountered during whole process
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

	err = os.Rename(updatedFile, path)
	return err
}

// Sync directory after file creation/rename to ensure presence is persisted in the file system structure
func SyncDirectory(directoryPath string) error {
	directory, err := os.Open(directoryPath)

	if err != nil {
		return err
	}

	err = directory.Sync()

	if err != nil {
		return err
	}

	err = directory.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	return err
}

func main() {
	homeDirectory, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	dbDirectory := filepath.Join(homeDirectory, "Documents")
	filePath := filepath.Join(dbDirectory, "argentPage")

	data, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = Save(filePath, append(data, []byte("\ntest data")...))

	if err != nil {
		log.Fatal(err.Error())
	}

	err = SyncDirectory(dbDirectory)

	if err != nil {
		log.Fatal(err.Error())
	}
}

