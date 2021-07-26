package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var mapSize int
var elements map[string]string

func set(cmd []string, elements map[string]string, conn net.Conn) {
	if len(elements) >= mapSize {
		conn.Write([]byte("Error! DB limit was reached"))
		return
	}
	switch {
	case len(cmd) > 3:
		conn.Write([]byte("Too many arguments!\nUsage: set <key> <value>"))
	case len(cmd) < 3:
		conn.Write([]byte("Too few arguments!\nUsage: set <key> <value>"))
	default:
		elements[cmd[1]] = cmd[2]
	}
	return
}

func get(cmd []string, elements map[string]string, conn net.Conn) {
	switch {
	case len(cmd) > 2:
		conn.Write([]byte("Too many arguments!\nUsage: get <key>"))
	case len(cmd) < 2:
		conn.Write([]byte("Too few arguments!\nUsage: get <key>"))
	default:
		conn.Write([]byte(elements[cmd[1]]))
	}
	return
}

func del(cmd []string, elements map[string]string, conn net.Conn) {
	switch {
	case len(cmd) > 2:
		conn.Write([]byte("Too many arguments!\nUsage: del <key>"))
	case len(cmd) < 2:
		conn.Write([]byte("Too few arguments!\nUsage: del <key>"))
	default:
		delete(elements, cmd[1])
	}
	return
}

func mapSizing() int {
	var err error

	fmt.Print("Enter the DB size: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Unable to read from stdin: %s", err)
	}
	text = strings.TrimSpace(text)
	result, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("Error while converting string to an int: %s", err)
	}
	if result == 0 {
		fmt.Println("WARNING! Null-sized DB will be created")
	}
	return result
}

func main() {

	var err error
	mapSize = mapSizing()
	
	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Unable to launch the server: %s", err)
	}
	conn, err := ln.Accept()
	if err != nil {
		log.Fatalf("Unable to connect a client: %s", err)
	}

	elements = make(map[string]string, mapSize)
	
	for {		
		fmt.Println("Listening to commands...")
		// Будем прослушивать все сообщения разделенные \n
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalf("Unable to read from stdin: %s", err)
		}
		fmt.Printf("Message Received: %s", message)

		// Парсер для полученной команды
		words := strings.Fields(message)
		switch words[0] {
		case "set":
			set(words, elements, conn)
		case "get":
			get(words, elements, conn)
		case "del":
			del(words, elements, conn)
		case "":
			conn.Write([]byte("Usage: del <key>"))
		default:
			conn.Write([]byte("parsing error!"))
		}
		conn.Write([]byte("\n"))
	}
}
