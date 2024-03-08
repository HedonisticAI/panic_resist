package httphandler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"panic_resist/create"
	"panic_resist/safety"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type Queue struct {
	TheQ *amqp.Connection
	Name string
}

func (A *Queue) PostLine(c *gin.Context) {
	defer safety.SaveData(c)
	//panic("planned panic")

	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Request.Response.StatusCode = http.StatusInternalServerError
	}
	var name create.Request
	json.Unmarshal(bytes, &name)
	ch, err := A.TheQ.Channel()
	if err != nil {
		log.Print("QUEQUE inactive")
	}
	ch.Publish("", A.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(name.Data)})

}
