package config

import (
	"fmt"
	"os"
	"path/filepath"
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
	return "/home/eldius/dev/workspaces/go/webcomics-fetcher2-plugins/webcomics-fetcher2-oots/../plugins/"
}
