package repository

import (
	"fmt"

	"github.com/Eldius/webcomics-fetcher2-go/comics"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
)

/*
WebcomicRepository is responsive for manage comics persistence
*/
type WebcomicRepository struct {
	db *storm.DB
}

/*
SaveComic saves a new comic data
*/
func (r *WebcomicRepository) SaveComic(c *comics.Webcomic) {
	r.db.Save(c)
}

/*
SaveComicStrip saves a ComicStrip
*/
func (r *WebcomicRepository) SaveComicStrip(s *comics.ComicStrip) {
	r.db.Save(s)
}

/*
SaveComicStrip saves a ComicStrip
*/
func (r *WebcomicRepository) ListComicStrip(name string) []*comics.ComicStrip {
	r.db.Select(q.Eq("", name))
}

/*
NewRepository creates a new repository
*/
func NewRepository() *WebcomicRepository {
	return &WebcomicRepository{
		db: NewDB(),
	}
}

/*
NewRepositoryWithDB creates a new WebcomicRepository
*/
func NewRepositoryWithDB(db *storm.DB) *WebcomicRepository {
	return &WebcomicRepository{
		db: db,
	}
}

/*
NewDB returns the default DB for webcomics data
*/
func NewDB() *storm.DB {
	return NewCustomDB("webcomics.db")
}

/*
NewCustomDB returns a DB pointing to a custom db file
*/
func NewCustomDB(dbFile string) *storm.DB {
	db, err := storm.Open(fmt.Sprintf("%s.db", dbFile))
	if err != nil {
		panic(err.Error())
	}
	return db
}
