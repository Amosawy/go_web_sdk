package tools

import (
	"os"
)

func GetRouteConfigPath() (string, string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", "", err
	}
	absPath := path + "/config/config-route.toml"
	return absPath, path, nil
}
