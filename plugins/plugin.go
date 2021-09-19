package plugins

import (
	"fmt"
	"os"
	"path/filepath"
)

type PluginEngine struct {
	pluginFolder string
}

type Plugin struct {
	Name        string
	Path        string
	Description string
}

func NewPluginEngine() *PluginEngine {
	return &PluginEngine{
		pluginFolder: filepath.Join(GetBinaryPath(), "plugins"),
	}
}

func (e *PluginEngine) ListPlugins() []string {
	return make([]string, 0)
}

func (e *PluginEngine) RefreshPluginList() map[string]Plugin {
	return make(map[string]Plugin)
}

func GetBinaryPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(ex)
	fmt.Println(exePath)
	return exePath
}
