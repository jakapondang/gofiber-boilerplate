package configuration

import (
	"github.com/joho/godotenv"
	"goamartha/exception"
	"os"
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
	err := godotenv.Load(filenames...)
	exception.PanicLogging(err)
	return &envImpl{}
}
