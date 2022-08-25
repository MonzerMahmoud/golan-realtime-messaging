package main

import (
	"fmt"
	"log"
	"net"
)

const (
	HOST_NAME = "localhost"
	PORT_NUM = "5054"
	CONN_TYPE = "tcp"
)

type User struct {
	id int
	conn net.Conn
}

var (
	numberOfConnectedUsers = 0
	users []User
)

func checkErr(err error) {
	if err != nil {
		log.Fatal((err))
	}
}

func main() {

	psock,err := net.Listen(CONN_TYPE,HOST_NAME+":"+PORT_NUM)

	checkErr(err)

	fmt.Println("Server is running on PORT :: "+string(PORT_NUM))

	defer psock.Close()

	for {
		sock, err := psock.Accept()

		checkErr(err)

		go handleUser(sock)

	}

}

func handleUser(sock net.Conn) {

	buf := make([]byte, 1024)

	numberOfConnectedUsers += 1
	
	users = append(users, User{id:numberOfConnectedUsers, conn: sock})

	fmt.Println("New user connected")

	fmt.Println("Number of connected users is :: " + fmt.Sprint(numberOfConnectedUsers))

	_,err := sock.Read(buf)

	checkErr(err)

	for {
		if availUsers() {
			sock.Write([]byte("1"))
		} else {
			sock.Write([]byte("0"))
		}
	}
}

func availUsers() bool {

	if numberOfConnectedUsers > 1 {
		return true
	}

	return false
}