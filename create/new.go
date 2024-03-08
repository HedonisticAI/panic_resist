package create

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

func NewDB() *sql.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PWD")
	Conn := fmt.Sprintf("host=%s user=%s port=%s dbname=%s password=%s sslmode=disable", host, user, port, dbname, password)
	re, err := sql.Open("postgres", Conn)
	if err != nil {
		log.Print(err.Error())
		return nil
	}
	return re
}

func NewQueQue() *amqp.Connection {
	conn := fmt.Sprintf("amqp://guest:guest@localhost:%s/", os.Getenv("QUEUE_PORT_IN"))
	res, err := amqp.Dial(conn)
	if err != nil {
		return nil
	}
	return res
}
