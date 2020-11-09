package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/d1y/note/conf"
	"github.com/d1y/note/utils/env"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/sqlite"
)

var dev = "./cache/dev.db"
var prod = "./note.db"

func getDatabaseFile() string {
	switch env.AppMode {
	case env.DebugCode:
		return dev
	case env.ReleaseCode:
		return prod
	default:
		return dev
	}
}

var sqlx = sqlite.ConnectionURL{
	Database: getDatabaseFile(),
}

// NoteDB `note` 会话
var NoteDB db.Collection

// Session `sql` 会话
var Session db.Session

func init() {
	var session, err = sqlite.Open(sqlx)
	if err != nil {
		var database = sqlx.Database
		fmt.Println("database", database)
		panic(err)
	}
	Session = session
	NoteDB = session.Collection(conf.DatabaseName)
	_, err = NoteDB.Exists()
	if errors.Is(err, db.ErrCollectionDoesNotExist) {
		// todo
		log.Printf("Collection does not exist: %v", err)
	}
	// err = NoteDB.Truncate()
	// if err != nil {
	// 	log.Fatalf("Truncate(): %q\n", err)
	// }
	// defer session.Close()
}
