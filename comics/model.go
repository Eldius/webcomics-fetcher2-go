package comics

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
	ID           int    `json:"id" storm:"id,increment"`
	URL          string `json:"url" storm:"unique"`
	Name         string `json:"name"`
	WebcomicName string `json:"webcomic_name"`
}
