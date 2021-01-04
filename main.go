package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)


func getInterval() int{
	interval_s, err := getEnv("MESSAGE_INTERVAL")
	if err != nil {
		return 5
	}
	interval, err := strconv.Atoi(interval_s)
	if err != nil {
		fmt.Printf("%s must be an integer", "MESSAGE_INTERVAL")
		return 5
	}
	return interval
}

func getMockMessage() []byte{
	mock_message, err := getEnv("MOCK_MESSAGE")
	if err != nil {
		return []byte("Default Message from Go Websocket Mock")
	}
	return []byte(mock_message)
}

func reader(conn *websocket.Conn) {
	for {
		if err:= conn.WriteMessage(1, getMockMessage() ); err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Duration(getInterval()) * time.Second)
	}
}


func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client successfully connected")
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}

func getEnv(envKey string) (string, error) {
	val, ok := os.LookupEnv(envKey)
	if !ok {
		return "", errors.New(fmt.Sprintf("%s not set\n", envKey))
	} else {
		return val, nil
	}
}

func main() {
	setupRoutes()
	panic(http.ListenAndServe(":8085",nil))
}