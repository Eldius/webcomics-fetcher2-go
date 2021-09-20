package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/Eldius/webcomics-fetcher2-go/config"
)

type PluginEngine struct {
	pluginFolder string
}

func NewPluginEngine() *PluginEngine {
	return &PluginEngine{
		pluginFolder: config.GetPluginsFolder(),
	}
}

func (e *PluginEngine) ListPlugins() []PluginInfo {
	var result []PluginInfo
	files, err := ioutil.ReadDir(e.pluginFolder)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan []byte)
	go fetchPluginsInfo(files, c)
	for b := range c {
		var p PluginInfo
		fmt.Println(string(b))
		err := json.Unmarshal(b, &p)
		if err != nil {
			fmt.Printf("Failed to parse plugin info: %s\n", err.Error())
		}
		result = append(result, p)
	}
	return result
}

func fetchPluginsInfo(files []fs.FileInfo, c chan []byte) {
	var wg sync.WaitGroup
	for _, f := range files {
		if !f.IsDir() {

			wg.Add(1)
			cmd := exec.Command(filepath.Join(config.GetPluginsFolder(), f.Name()), "info")
			go execCmd(cmd, &wg, c)
		}
	}
	wg.Wait()
	close(c)
}

func (e *PluginEngine) RefreshPluginList() map[string]PluginInfo {
	return make(map[string]PluginInfo)
}

func execCmd(cmd *exec.Cmd, wg *sync.WaitGroup, c chan []byte) {
	defer wg.Done()
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	_ = cmd.Run()

	c <- buf.Bytes()
}

func GetAbsolutePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return ex
}
