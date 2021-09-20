package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func GetBinaryPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(ex)
	fmt.Println(exePath)
	return exePath
}

func GetPluginsFolder() string {
	//return filepath.Join(GetBinaryPath(), "..", "plugins")
	return viper.GetString("webcomics.plugins.folder")
}
