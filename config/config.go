package configuration

import (
	"github.com/joho/godotenv"
	"goamartha/exception"
	"os"
	"path/filepath"
	"runtime"
)

type Env interface {
	Get(key string) string
}

type envImpl struct {
}

func (env *envImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Env {
	// Load config
	err := godotenv.Load(filenames...) // path default
	if err == nil {
		return &envImpl{}
	}
	// Get the directory of the currently executing file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		exception.PanicLogging("Error getting current file information")
	}

	// Find the root directory (assuming this file is located in a subdirectory)
	// Adjust the number of ".." according to the depth of your file in the directory structure
	rootDir := filepath.Join(filepath.Dir(filename), ".", "..") // Adjust as needed

	// Print the root directory
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		exception.PanicLogging("Error getting absolute path: %v" + err.Error())
	}

	envFilePath := absRootDir + "/.env"
	err = godotenv.Load(envFilePath)
	exception.PanicLogging(err)
	return &envImpl{}
}
