package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var (
	response = make([]byte, 1024)
)
func checkErr(err error) {
	if err != nil {
		log.Fatal((err))
	}
}

func checkForArgs() {
	if len(os.Args) != 3 {
		log.Fatal("Usage client <hostname> <port>")
	}

	fmt.Printf("Connected to server %s on port %d \n", os.Args[1], os.Args[2])
}

func main() {

	checkForArgs()

	sock, err := net.Dial("tcp", os.Args[1]+":"+os.Args[2])
	checkErr(err)

	handleOperation(sock)
	sock.Close()
}

func handleOperation(sock net.Conn) {

	_,err := sock.Write([]byte("checking"))

	checkErr(err)

	res,err := sock.Read(response)

	println(int(res))
	if int(res) == 1 {
		println("Start chatting")
	} else {
		println("Please wait until another user log in ...")
	}
}