package safety

import (
	"database/sql"
	"log"
)

func clear(db *sql.DB) {
	db.Exec("DELETE FROM data")
	log.Printf("Database cleared")
}
