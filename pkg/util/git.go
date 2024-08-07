package util

import (
	"os"
)

func CloneRepository(url, dest string) error {
	options := git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	}
	_, err := git.PlainClone(dest, false, &options)
	return err
}
