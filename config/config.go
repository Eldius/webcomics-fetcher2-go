package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

/*
GetBinaryPath returns the path to the running binary
*/
func GetBinaryPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(ex)
	return exePath
}
/*
GetPluginsFolder returns the plugins folder
*/
func GetPluginsFolder() string {
	//return filepath.Join(GetBinaryPath(), "..", "plugins")
	return viper.GetString("webcomics.plugins.folder")
}
