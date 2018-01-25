package config

import (
	"gopkg.in/urfave/cli.v1"
)

var (
	// LoadConfigFileFlag used to specify loading app config params from a config file.
	LoadConfigFileFlag = cli.StringFlag{
		Name:        "config, c",
		Usage:       "Load configuration from `FILE`",
		Value:       ConfigValues.ConfigFilePath,
		Destination: &ConfigValues.ConfigFilePath,
	}

	//DataFolderPathFlag specifies app persisted data root directory.
	DataFolderPathFlag = cli.StringFlag{
		Name:        "dataFolder, df",
		Usage:       "Set root data folder`",
		Value:       ConfigValues.DataFilePath,
		Destination: &ConfigValues.DataFilePath,
	}
)
