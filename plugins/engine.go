package plugins

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/Eldius/webcomics-fetcher2-go/comics"
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
	_ = db.Init(&PluginEngine{})
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
	_ = e.db.Select(q.Not(q.In("Name", keys))).Delete(new(PluginInfo))
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
FetchStrips fetches
*/
func (e *PluginEngine) FetchStrips(name string) {
	var p PluginInfo
	if err := e.db.Select(q.Eq("Name", name)).First(&p); err != nil {
		fmt.Printf("Could not find plugin '%s': %s", name, err.Error())
		return
	}

	strips := p.FetchStrips()
	for _, s := range strips {
		fmt.Printf("- saving %s\n", s.Name)
		_ = e.comicRepo.SaveComicStrip(s)

		fmt.Printf("- Downloading %s\n", s.Name)
		_, _ = s.Download()
	}
}

/*
ListStrips list all strip info from database
*/
func (e *PluginEngine) ListStrips(name string) ([]*comics.ComicStrip, error) {
	return e.comicRepo.ListComicStrip(name)
}

/*
LoadPlugins lists plugins in plugin folder
*/
func (e *PluginEngine) LoadPlugins() []*PluginInfo {
	var result []*PluginInfo
	fmt.Printf("Loading plugins from %s\n", e.pluginFolder)
	files, err := ioutil.ReadDir(e.pluginFolder)
	if err != nil {
		log.Fatalf("Failed to list plugin files in folder '%s': %s", e.pluginFolder, err.Error())
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
			go execCmdAsyncOutputToFile(&wg, c, filepath.Join(config.GetPluginsFolder(), f.Name()), "info")
		}
	}
	wg.Wait()
	close(c)
}

func execCmdOutputToFile(p string, scmd string, params ...string) ([]byte, error) {

	o, clear := CreateOutputFile()
	defer clear()

	args := append([]string{scmd, "-o", o}, params...)
	cmd := exec.Command(p, args...)
	cmd.Stdout = os.Stdout
	fmt.Println("Executing command:")
	fmt.Println(cmd.String())
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to get plugin info: %s\n", err.Error())
		return make([]byte, 0), err
	}

	out, err := os.Open(o)
	if err != nil {
		fmt.Printf("Failed to open output file '%s': %s\n", o, err.Error())
		return make([]byte, 0), err
	}
	return ioutil.ReadAll(out)
}

func execCmdAsyncOutputToFile(wg *sync.WaitGroup, c chan []byte, p string, scmd string, params ...string) {
	defer wg.Done()

	b, err := execCmdOutputToFile(p, scmd, params...)
	if err != nil {
		fmt.Printf("Failed to get plugin info from %s: %s\n", p, err.Error())
		return
	}

	c <- b
}
