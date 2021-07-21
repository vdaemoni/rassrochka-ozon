package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"strconv"
	"os"
)

var mapSize int
var elements map[string]string

func set(cmd []string, elements map[string]string, conn net.Conn) {
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
	fmt.Print("Enter the DB size: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	result, _ := strconv.Atoi(text)
	if result == 0 {
		fmt.Println("WARNING! Null-sized DB will be created")
	}
	return result
}

func main() {

	mapSize = mapSizing()
	
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	
	elements = make(map[string]string, mapSize)
	
	for {		
		fmt.Println("Listening to commands...")
		// Будем прослушивать все сообщения разделенные \n
		message, _ := bufio.NewReader(conn).ReadString('\n')
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
		default:
			conn.Write([]byte("parsing error!"))
		}
		conn.Write([]byte("\n"))
	}
}
