package plugins

/*
PluginInfo store plugin data
*/
type PluginInfo struct {
	Name        string `storm:"id" json:"name"`
	Path        string `storm:"unique" json:"path"`
	Description string `json:"description"`
}
