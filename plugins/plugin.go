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
	"github.com/Eldius/webcomics-fetcher2-go/repository"
	"github.com/asdine/storm/v3"
)

/*
PluginEngine handles plugins
*/
type PluginEngine struct {
	pluginFolder string
	db           *storm.DB
}

/*
NewPluginEngine creates a new PluginEngine
*/
func NewPluginEngine() *PluginEngine {
	return &PluginEngine{
		pluginFolder: config.GetPluginsFolder(),
		db:           repository.NewCustomDB("plugins"),
	}
}

/*
RefreshPluginList refreshes plugin database
*/
func (e *PluginEngine) RefreshPluginList() map[string]PluginInfo {
	plugins := e.ListPlugins()
	result := make(map[string]PluginInfo)
	for _, p := range plugins {
		result[p.Name] = p
		e.db.Save(p)
	}
	return result
}

/*
ListPlugins lists plugins in plugin folder
*/
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

func execCmd(cmd *exec.Cmd, wg *sync.WaitGroup, c chan []byte) {
	defer wg.Done()
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	_ = cmd.Run()

	c <- buf.Bytes()
}

/*
GetAbsolutePath returns binary absolute path (to be used from plugins to return self location)
*/
func GetAbsolutePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return ex
}
