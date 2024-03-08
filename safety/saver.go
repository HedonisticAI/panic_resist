package safety

import (
	"encoding/json"
	"io"
	"log"
	"panic_resist/create"

	"github.com/gin-gonic/gin"
)

func SaveData(c *gin.Context) {
	if err := recover(); err != nil {
		log.Printf("Marking unsent data: %v", err)
		bytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Print(err.Error())
			return
		}
		var Data create.Request
		json.Unmarshal(bytes, &Data)
		db := create.NewDB()
		db.Exec("INSERT INTO data(data_string) VALUES ($1)", Data.Data)
		defer db.Close()
	}
}
