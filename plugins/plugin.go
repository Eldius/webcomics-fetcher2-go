package plugins

import (
	"encoding/json"
	"log"

	"github.com/Eldius/webcomics-fetcher2-go/comics"
)

/*
PluginInfo store plugin data
*/
type PluginInfo struct {
	Name        string `storm:"id" json:"name"`
	Path        string `storm:"unique" json:"path"`
	Description string `json:"description"`
}

/*
FetchStrips calls plugin fetch command
*/
func (p *PluginInfo) FetchStrips() []*comics.ComicStrip {
	result := make([]*comics.ComicStrip, 0)

	b, err := execCmdOutputToFile(p.Path, "fetch")
	if err != nil {
		log.Fatalf("Failed to unmarshal plugin fetch response for '%s': %s\nresponse:\n%s\n---\n", p.Name, err.Error(), string(b))
	}
	//fmt.Println(string(b))
	if err := json.Unmarshal(b, &result); err != nil {
		log.Fatalf("Failed to parse plugin response: %s", err.Error())
	}

	return result
}

/*
FetchWebcomics calls plugin fetch command
*/
func (p *PluginInfo) FetchWebcomics() *comics.Webcomic {
	return nil
}
