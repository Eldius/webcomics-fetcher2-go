package server

type IndexVO struct {
	Webcomics []WebcomicVO
}

type WebcomicVO struct {
	Name        string
	Description string
}

type StripVO struct {
	Name         string
	WebcomicName string
	Order        int
	FileName     string
}

type StripPageVO struct {
	Name   string
	Strips []StripVO
}
