package comics

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/Eldius/webcomics-fetcher2-go/config"
)

/*
DownloadStrip downloads strip image
*/
func DownloadStrip(src string, dest string) error {
	c := http.Client{}

	res, err := c.Get(src)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	outFile := filepath.Join(config.GetStripsFolder(), dest)
	if err = createFolder(outFile); err != nil {
		return err
	}
	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, res.Body); err != nil {
		return err
	}

	return nil
}

/*
SanitizeFilename replaces special characters from name
*/
func SanitizeFilename(filename string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9-]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(filename, "_")
}

func createFolder(f string) error {
	folder := filepath.Dir(f)
	if _, err := os.Stat(folder); err != nil {
		return os.MkdirAll(folder, os.ModePerm)
	}
	return nil
}
