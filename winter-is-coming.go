package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"math/rand"
	"time"
	"winter-is-comming/GameEngineAPI"
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

			conn := startServer(PORT)

			rand.Seed(time.Now().UnixNano())

			GameEngineAPI.init_board()
			conn.Write("Start Game: Y/N")
			var startGame string
			fmt.Scanln(&startGame)

			if (strings.ToLower(strings.TrimSpace(startGame)) == "y"){
				GameEngineAPI.update_board(conn)

				for !board.gameOver{
					var Xcoord string
					fmt.Scanln(&Xcoord)

					if Xcoord, err := strconv.Atoi(Xcoord); err == nil{
						GameEngineAPI.moveJohnSnow(Xcoord)
						for (Xcoord > board.maxX || Xcoord < board.minX){
							conn.Write()
							conn.Write("X coordenate that you have entered is not valid please enter an integer between 0 and 9.")
							fmt.Scanln(&Xcoord)
							GameEngineAPI.moveJohnSnow(Xcoord)
						}
					}


					if (strings.TrimSpace(Xcoord) == "shoot"){
						if( !board.arrowOnBoard ){
							GameEngineAPI.board.arrow.y = 1
							GameEngineAPI.board.arrow.x = board.JohnSnow.x
							GameEngineAPI.board.arrowOnBoard = true
							conn.Write(GameEngineAPI.board.arrow)
						}else{
							fmt.Println()
							fmt.Println("Max number of arrows reached!")
							fmt.Println()
						}
					}
				}
			}

func startServer(PORT string) {
		fmt.Println("Init Server")

		l, err := net.Listen("tcp", ":" + PORT)
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		defer l.Close()

		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}

		go server_listen()

		return conn


}


func server_listen(){
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
