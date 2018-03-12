package GameEngineAPI

import (
	"fmt"
	"math/rand"
	"time"
	"net"
)

var board Board

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


  conn.Write([]byte("Winter is Coming" + "\n" ))
  conn.Write([]byte("" + "\n" ))
  conn.Write([]byte("Rules:"))
  conn.Write([]byte("1: 10x30 board" + "\n" ))
  conn.Write([]byte("2: the max number of arrows on the board at one time is 1" + "\n" ))
  conn.Write([]byte("3: type shoot to shoot arrow" + "\n" ))
  conn.Write([]byte("4: type number to move John Snow" + "\n" ))
  conn.Write([]byte("5: Night King will move once every 5 seconds" + "\n" ))
  conn.Write([]byte("6: Arrows move 1 square every second in a straight line" + "\n" ))
  conn.Write([]byte(" Please enter integer X coordenate between 0 and 9 that John  will shoot arrow from the wall." + "\n" ))
}


func update_board(conn net.Conn){
  boardPointer := &board
	connPointer := &conn
	go update_arrow(boardPointer, connPointer)
	go move_night_king(boardPointer, connPointer)
}

func update_arrow(board *Board, conn *net.Conn){
  for (!board.gameOver){
    time.Sleep(time.Second * 1)
    if (!(board.arrow.x == 0) && !(board.arrow.y == 0)){
      board.arrow.y =  board.arrow.y + 1
    }


    if (board.arrow.y >= 30){
      board.arrow.x = 0
      board.arrow.y = 0
      board.arrowOnBoard = false
    }
    if board.arrowOnBoard{
      print_board(*conn)
    }

    if (board.arrow.x == board.NightKing.x && board.arrow.y == board.NightKing.y){
      game_over(true, *conn)
    }
  	}
  }

func print_board(conn net.Conn){
  conn.Write([]byte( "\n" ))
  conn.Write([]byte(board.NightKing.String() + "\n" ))
  conn.Write([]byte(board.arrow.String() + "\n" ))
  conn.Write([]byte(board.JohnSnow.String() + "\n" ))
  conn.Write([]byte( "\n" ))
}


func move_night_king(board *Board, conn *net.Conn){
	for !board.gameOver{
		time.Sleep(time.Second * 5)
		board.NightKing.y = board.NightKing.y - 1
		direction := rand.Intn(3)
		switch direction{
			case 0:{
				if (board.NightKing.x <= board.minX){
					board.NightKing.x = board.NightKing.x + 1
				}else{
					board.NightKing.x = board.NightKing.x - 1
				}
			}
			case 1:
				board.NightKing.x = board.NightKing.x
			case 2:{
				if (board.NightKing.x >= board.maxX){
					board.NightKing.x = board.NightKing.x - 1
				}else{
					board.NightKing.x = board.NightKing.x + 1
				}
			}
		}

    if !board.arrowOnBoard{
      print_board(*conn)
    }
		if (board.NightKing.y == 0){
			game_over(false, *conn)
		}
	}
}

func game_over(win bool, conn net.Conn){
  var message string
  if win{
      message = "      The Night King is dead you have won!"
  }else{
    message = "              GAME OVER!               "
  }

  i := 0
	for ( i < 10) {
			conn.Write([]byte("\n" ))
			if (i == 5){
				conn.Write([]byte(message + "\n"))
			}
			i++
		}
	board.gameOver = true
}
func moveJohnSnow(x int){
  board.JohnSnow.x = x
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
