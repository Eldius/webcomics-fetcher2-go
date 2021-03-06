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
func (r *WebcomicRepository) SaveComic(c *comics.Webcomic) error {
	return r.db.Save(c)
}

/*
SaveComicStrip saves a ComicStrip
*/
func (r *WebcomicRepository) SaveComicStrip(s *comics.ComicStrip) error {
	return r.db.Save(s)
}

/*
ListComicStrip list ComicStrips by name
*/
func (r *WebcomicRepository) ListComicStrip(name string) ([]*comics.ComicStrip, error) {
	var result []*comics.ComicStrip
	err := r.db.Select(q.Eq("WebcomicName", name)).OrderBy("Order").Find(&result)
	return result, err
}

/*
ListAllComicStrip list ComicStrips by name
*/
func (r *WebcomicRepository) ListAllComicStrip() ([]*comics.ComicStrip, error) {
	var result []*comics.ComicStrip
	err := r.db.All(&result)
	return result, err
}

/*
NewRepository creates a new repository
*/
func NewRepository() *WebcomicRepository {
	return &WebcomicRepository{
		db: NewCustomDB("webcomics"),
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
NewCustomDB returns a DB pointing to a custom db file
*/
func NewCustomDB(dbFile string) *storm.DB {
	db, err := storm.Open(fmt.Sprintf("%s.db", dbFile))
	if err != nil {
		panic(err)
	}
	return db
}
