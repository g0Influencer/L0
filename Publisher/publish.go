package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

func main(){
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	sc, err := stan.Connect(os.Getenv("CLUSTER_ID"),os.Getenv("PROD_ID"),
		stan.NatsURL(stan.DefaultNatsURL),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Printf("Connection lost, reason: %v", reason)
		}))
	defer sc.Close()


	if err != nil {
		log.Printf("[err] Nats-pub: %s\n", err.Error())
		return
	}

	for i:=1; i <= 5; i++ {
		file, err := os.Open("./test_data/model"  + strconv.Itoa(i)  +".json")
		if err != nil {
			log.Printf("[err] Nats-pub open file: %s\n", err.Error())
			return
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("[err] Nats-pub ioutil: %s\n", err.Error())
			return
		}

		err = sc.Publish(os.Getenv("CHANNEL"), data)

		if err != nil {
			log.Printf("[err] Nats-pub: %s\n", err.Error())
		}
		fmt.Println("Publish was successful!")
		time.Sleep(5 * time.Second)
	}
}

