package gobin

import (
	"errors"
	"io"
)

type Config struct {
	GoBin  string
	Cwd    string
	Stderr io.Writer
}

type GoBin struct {
	goBin  string
	cwd    string
	stderr io.Writer
}

func New(c Config) (*GoBin, error) {
	if c.Cwd == "" {
		return nil, errors.New("missing BuildConfig field: Cwd")
	}
	if c.Stderr == nil {
		return nil, errors.New("missing BuildConfig field: Stderr")
	}

	// default GoBin to "go"
	if c.GoBin == "" {
		c.GoBin = "go"
	}

	return &GoBin{
		cwd:    c.Cwd,
		goBin:  c.GoBin,
		stderr: c.Stderr,
	}, nil
}
