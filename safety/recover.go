package safety

import (
	"log"
	"panic_resist/create"

	"github.com/streadway/amqp"
)

func Recovery(name string) (*amqp.Connection, error) {
	var data string
	db := create.NewDB()
	rows, err := db.Query("SELECT data_string FROM data")
	if err != nil {
		log.Printf("Error querying data")
		return nil, err
	}
	conn := create.NewQueQue()
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	_, err = ch.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return conn, nil
	}
	for rows.Next() {
		err := rows.Scan(&data)
		if err != nil {
			return nil, err
		}
		ch.Publish("", name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})

	}
	defer db.Close()

	log.Printf("Data recovery complete")
	clear(db)
	return conn, nil
}
