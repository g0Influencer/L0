package Nats

import (
	"L0/cache"
	"L0/database"
	"L0/models"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func Subscribe() {
	or := models.Order{}
	database.Connect()
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	sc, err := stan.Connect(os.Getenv("CLUSTER_ID"), os.Getenv("SUB_ID"),
		stan.NatsURL(stan.DefaultNatsURL),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Printf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		panic(err)
	}

	sub, err := sc.Subscribe(os.Getenv("CHANNEL"), func(m *stan.Msg) {
		if err := m.Ack(); err != nil {
			log.Println(err)
		}
		if err := json.Unmarshal(m.Data, &or); err != nil {
			log.Printf("[err] Nats-pub JSON: %s\n", err.Error())
		}
		err = or.Validate()
		if err != nil {
			log.Printf("Nats.validate: %s\n", err)
			log.Printf("msg: %s\n", string(m.Data))
			return
		}
		cache.Set(or)
		database.Insert(or)

	}, stan.SetManualAckMode())
	if err != nil {
		panic(err)
	}

	sub.Unsubscribe()
	sc.Close()

}
