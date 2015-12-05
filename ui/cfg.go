package ui

import (
	"fmt"
	"os"

	"errors"
)

// Cfg Maps to the lazarus config ini file.
type Cfg struct {
	TmpLocation string `ini:"tmp_location"`
	LogLocation string `ini:"log_location"`
	MazSize     string `ini:"max_size"`
	LogFile     string `ini:"log_file"`
}

// AllOk Makes sure that the mandatory values in the config file exist and are ok.
func (c *Cfg) AllOk() error {
	if err := c.isTmpLocationOk(); err != nil {
		return err
	}

	if err := c.isLogLocationOk(); err != nil {
		return err
	}

	return nil
}

func (c *Cfg) isLogLocationOk() error {
	return isLocationOk(c.LogLocation, "log_location")
}

func (c *Cfg) isTmpLocationOk() error {
	return isLocationOk(c.TmpLocation, "tmp_location")
}

func isLocationOk(loc string, name string) error {
	if loc == "" {
		return errors.New(fmt.Sprintf("Missing directive in the ini file: %s", name))
	}

	if _, err := os.Stat(loc); os.IsNotExist(err) {
		return err
	}

	return nil
}