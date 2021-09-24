package comics

import (
	"fmt"
)

/*
Webcomic is the webcomic
*/
type Webcomic struct {
	ID   int    `json:"id" storm:"id,increment"`
	Name string `json:"name" storm:"unique"`
	//Strips []ComicStrip `json:"strips"`
}

/*
ComicStrip is a single image
*/
type ComicStrip struct {
	ID           string `json:"id" storm:"id"`
	URL          string `json:"url" storm:"unique"`
	Name         string `json:"name"`
	WebcomicName string `json:"webcomic_name"`
	Order        int    `json:"order"`
	FileName     string `json:"file_name"`
}

/*
RelativeFilename returns the strip relative file path
*/
func (s *ComicStrip) RelativeFilename() string {
	return s.FileName
}

/*
Download downloads strip image
*/
func (s *ComicStrip) Download() (string, error) {
	relPath := s.RelativeFilename()
	fmt.Println(relPath)
	err := DownloadStrip(s.URL, relPath)
	return relPath, err
}
