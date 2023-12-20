package pokta

import (
	_ "embed"
	"io"
	"os"

	"github.com/piaverous/pokta/pokta/config"
)

type App struct {
	Config *config.Config

	Out io.Writer
	Err io.Writer
}

func New() (*App, error) {
	app := &App{
		Config: &config.Config{},
		Out:    os.Stdout,
		Err:    os.Stderr,
	}
	return app, nil
}
