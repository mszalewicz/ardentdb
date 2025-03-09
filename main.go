package main

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/mszalewicz/ardentdb/database"
)

func main() {
	currentDirectory, err := os.Getwd()

	if err != nil {
		log.Fatal(err.Error())
	}

	logPath := filepath.Join(currentDirectory, "log")

	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		slog.Error("Could not create log file.", "error", err)
	}

	loggerArgs := &slog.HandlerOptions{AddSource: true}
	logger := slog.New(slog.NewJSONHandler(logFile, loggerArgs))
	slog.SetDefault(logger)

	homeDirectory, err := os.UserHomeDir()

	if err != nil {
		slog.Error(err.Error())
	}

	dbDirectory := filepath.Join(homeDirectory, "Documents")
	filePath := filepath.Join(dbDirectory, "argentPage")

	data, err := os.ReadFile(filePath)

	if err != nil {
		slog.Error(err.Error())
	}

	err = database.Save(filePath, append(data, []byte("\ntest data")...))

	if err != nil {
		slog.Error(err.Error())
	}

	err = database.SyncDirectory(dbDirectory)

	if err != nil {
		slog.Error(err.Error())
	}
}
