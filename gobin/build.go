package gobin

import (
	"os/exec"
	"syscall"
)

type BuildConfig struct {
	Source string
	Output string
}

func (b *GoBin) Build(c BuildConfig) (int, error) {
	args := []string{"build"}
	if c.Output != "" {
		args = append(args, "-o", c.Output)
	}
	if c.Source != "" {
		args = append(args, c.Source)
	}

	cmd := exec.Command(b.goBin, args...)
	cmd.Dir = b.cwd
	cmd.Stderr = b.stderr

	if err := cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus(), nil
			}
		}
		// If it's not an execerr or we can't get the status, return err
		return 0, err
	}

	return 0, nil
}
