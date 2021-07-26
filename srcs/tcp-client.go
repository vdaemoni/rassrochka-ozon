package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	var err error

	// Подключаемся к сокету
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalf("Unable to connect to a localhost: %s", err)
	}
	for {
		// Чтение входных данных от stdin
		fmt.Print("Text to send: ")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Unable to read from stdin: %s", err)
		}

		text = strings.TrimSpace(text)
		if text == "exit" {
			fmt.Println("bye!")
			conn.Close()
			os.Exit(0)
		}
		
		if len(text) > 0 {
			// Отправляем в socket
			fmt.Fprintf(conn, text+"\n")
			// Прослушиваем ответ
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Fatalf("Error with receiving the message from server: %s", err)
			}
			fmt.Print("Message from server: " + message)
		} else {
			fmt.Print("\n")
		}
	}
}
