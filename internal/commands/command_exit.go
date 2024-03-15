package commands

import "os"

func CommandExit(c *Config, param string) error {
	os.Exit(0)
	return nil
}
