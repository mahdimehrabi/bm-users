package env

import (
	"github.com/joho/godotenv"
	"os"
)

type Env struct {
	Port       string
	RedisAddr  string
	DatabaseCS string
	Secret     string
}

func NewEnv() *Env {
	return &Env{}
}

func (e *Env) LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	e.Port = os.Getenv("Port")
	e.RedisAddr = os.Getenv("RedisAddr")
	e.DatabaseCS = os.Getenv("DatabaseCS")
	e.Secret = os.Getenv("Secret")
	return nil
}
