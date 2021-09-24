package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/Eldius/webcomics-fetcher2-go/config"
	"github.com/Eldius/webcomics-fetcher2-go/plugins"
)

var (
	t = template.Must(template.ParseGlob("templates/*.html"))
	//t = template.Must(template.New("templates").ParseGlob("*.html"))
)

func Start(port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	engine := plugins.NewPluginEngine()

	fs := http.FileServer(http.Dir(config.GetStripsFolder()))
	mux.Handle("/strips/img/", http.StripPrefix("/strips/img", fs))

	mux.HandleFunc("/strips/", HandleStrips(engine))
	mux.HandleFunc("/", IndexHandler(engine))

	return mux
}

func HandleStrips(e *plugins.PluginEngine) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("handling:", r.RequestURI)
		stripName := strings.TrimLeft(r.RequestURI, "/strips/")
		if list, err := e.ListStrips(stripName); err != nil {
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"error": err.Error(),
				"url":   r.RequestURI,
			})
		} else {
			result := &StripPageVO{}
			for _, s := range list {
				result.Strips = append(result.Strips, StripVO{
					Name:         s.Name,
					WebcomicName: stripName,
					Order:        s.Order,
					FileName:     s.FileName,
				})
			}
			t.ExecuteTemplate(rw, "stripsList.html", result)
		}
	}
}

func IndexHandler(engine *plugins.PluginEngine) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, _ *http.Request) {
		log.Println("Processing index request...")
		response := &IndexVO{}
		for _, p := range engine.ListRegisteredPlugins() {
			response.Webcomics = append(response.Webcomics, WebcomicVO{
				Name:        p.Name,
				Description: p.Description,
			})
		}
		err := t.ExecuteTemplate(rw, "index.html", &response)
		if err != nil {
			log.Printf("Failed to parse template index: %s", err.Error())
		}
	}
}
