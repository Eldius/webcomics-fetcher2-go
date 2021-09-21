package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/Eldius/webcomics-fetcher2-go/config"
	"github.com/Eldius/webcomics-fetcher2-go/repository"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
)

/*
PluginEngine handles plugins
*/
type PluginEngine struct {
	pluginFolder string
	db           *storm.DB
	comicRepo    repository.WebcomicRepository
}

/*
NewPluginEngine creates a new PluginEngine
*/
func NewPluginEngine() *PluginEngine {
	db := repository.NewCustomDB("plugins")
	db.Init(&PluginEngine{})
	return &PluginEngine{
		pluginFolder: config.GetPluginsFolder(),
		db:           db,
		comicRepo:    *repository.NewRepository(),
	}
}

/*
RefreshPluginList refreshes plugin database
*/
func (e *PluginEngine) RefreshPluginList() map[string]*PluginInfo {
	plugins := e.LoadPlugins()
	result := make(map[string]*PluginInfo)
	keys := make([]string, 0)
	for _, p := range plugins {
		log.Printf("Registering '%s'", p.Name)
		result[p.Name] = p
		keys = append(keys, p.Name)
		if err := e.db.Save(p); err != nil {
			fmt.Printf("Failed to register plugin '%s': %s\n", p.Name, err.Error())
		}
	}
	e.db.Select(q.Not(q.In("Name", keys))).Delete(new(PluginInfo))
	return result
}

/*
ListRegisteredPlugins list all registered plugins
*/
func (e *PluginEngine) ListRegisteredPlugins() []PluginInfo {
	var result []PluginInfo
	if err := e.db.All(&result); err != nil {
		panic(err.Error())
	}
	return result
}

/*
Fetch fetches
*/
func (e *PluginEngine) Fetch(name string) {
	var p PluginInfo
	if err := e.db.Select(q.Eq("Name", name)).First(&p); err != nil {
		fmt.Printf("Could not find plugin '%s': %s", name, err.Error())
		return
	}

	strips := p.FetchStrips()
	for _, s := range strips {
		e.comicRepo.SaveComicStrip(s)
	}

}

func (e *PluginEngine) List(name string) {

}

/*
LoadPlugins lists plugins in plugin folder
*/
func (e *PluginEngine) LoadPlugins() []*PluginInfo {
	var result []*PluginInfo
	files, err := ioutil.ReadDir(e.pluginFolder)
	if err != nil {
		log.Fatalf("Failed to list plugins in folder '%s': %s", e.pluginFolder, err.Error())
	}

	c := make(chan []byte)
	go fetchPluginsInfo(files, c)
	for b := range c {
		var p *PluginInfo
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
			go execCmdAsync(cmd, &wg, c)
		}
	}
	wg.Wait()
	close(c)
}

func execCmd(cmd *exec.Cmd) ([]byte, error) {
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to get plugin info: %s\n", err.Error())
		return buf.Bytes(), err
	}

	return buf.Bytes(), nil
}

func execCmdAsync(cmd *exec.Cmd, wg *sync.WaitGroup, c chan []byte) {
	defer wg.Done()
	b, err := execCmd(cmd)
	if err != nil {
		fmt.Printf("Failed to get plugin info: %s\n", err.Error())
		return
	}

	c <- b
}
