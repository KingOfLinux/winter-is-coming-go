package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"math/rand"
	"time"
)

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        fmt.Println("Can not get IP")
				os.Exit(100)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

var board Board


func main() {
			fmt.Println("Your IP is " + GetOutboundIP().String())
		  PORT := "1945"
			fmt.Println("Your IP and port are " + GetOutboundIP().String() + ":" + PORT + " Please give this to your friend so they can connect.")

			rand.Seed(time.Now().UnixNano())


			startServer(PORT)
			os.Exit(100)
		}

		type Board struct{
			gameOver bool
		  arrowOnBoard bool
			maxX int
			maxY int
			minX int
			minY int
			JohnSnow
			NightKing
		  arrow
		}

		func init_board(conn net.Conn){

			board = Board{false, false, 9, 29, 0, 0, JohnSnow{x: 5,y: 0 }, NightKing{ x: rand.Intn(10) ,y: 30 }, arrow{0, 0}}


		  conn.Write([]byte("Winter is Coming" + "\n"))
		  conn.Write([]byte(""+ "\n"))
		  conn.Write([]byte("Rules:"+ "\n"))
		  conn.Write([]byte("1: 10x30 board"+ "\n"))
		  conn.Write([]byte("2: the max number of arrows on the board at one time is 1"+ "\n"))
		  conn.Write([]byte("3: type shoot to shoot arrow"+ "\n"))
		  conn.Write([]byte("4: type number to move John Snow"+ "\n"))
		  conn.Write([]byte("5: Night King will move once every 5 seconds"+ "\n"))
		  conn.Write([]byte("6: Arrows move 1 square every second in a straight line"+ "\n"))
		  conn.Write([]byte(" Please enter integer X coordenate between 0 and 9 that John  will shoot arrow from the wall."+ "\n"))
		}


func startServer(PORT string) {
		fmt.Println("Init Server")

		l, err := net.Listen("tcp", ":" + PORT)
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		defer l.Close()

		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}

		init_board(c)

		for{
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				os.Exit(100)
			}
			fmt.Print("-> ", string(netData))


			if strings.TrimSpace(string(netData)) == "STOP" {
				fmt.Println("Exiting TCP server!")
				return
			}
		}
}

type NightKing struct{
	x int
	y int
}

func (n NightKing) String() string{
	return fmt.Sprintf("Night King : (%d, %d) " , n.x, n.y )
}

type arrow struct{
	x int
	y int
}

func (a arrow) String() string{
	return fmt.Sprintf("Arrow : (%d, %d) " , a.x, a.y )
}

type JohnSnow struct{
	x int
	y int
}

func (j JohnSnow) String() string{
	return fmt.Sprintf("John Snow: (%d, %d) " , j.x, j.y )
}
