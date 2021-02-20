package impl

import (
	"fmt"

	"github.com/adrg/xdg"
)

const (
	bookmarksTable = "bookmarks"
	bukuDefaultDB  = "bookmarks.db"
)

type Bookmark struct {
	ID       int    `db:"id"`
	URL      string `db:"URL"`
	Metadata string `db:"metadata"`
	Tags     string `db:"tags"`
	Desc     string `db:"desc"`
	Flags    int    `db:"flags"`
}

func getDefaultBukuDatabase() string {
	return fmt.Sprintf("%s/buku/%s", xdg.DataHome, bukuDefaultDB)
}

func importBukuDB(dbFile string) error {
	return nil
}

func exportBukuDB(dbFile string) error {
	return nil
}
