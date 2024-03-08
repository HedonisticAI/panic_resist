package main

import (
	"log"
	"os"
	httphandler "panic_resist/http_handler"
	"panic_resist/safety"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var Q httphandler.Queue
	Q.Name = os.Getenv("QUEQUE_NAME")
	q, err := safety.Recovery(Q.Name)
	if err != nil {
		log.Print(err.Error())
		return
	}
	Q.TheQ = q
	r := gin.Default()
	r.POST("/", Q.PostLine)
	r.Run()
}
