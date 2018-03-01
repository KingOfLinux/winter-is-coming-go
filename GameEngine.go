package main


import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var board Board

func main(){

  rand.Seed(time.Now().UnixNano())

	init_board()
  fmt.Println("Start Game: Y/N")
  var startGame string
  fmt.Scanln(&startGame)

  if (strings.ToLower(strings.TrimSpace(startGame)) == "y"){
    update_board()

  	for !board.gameOver{
  	  var Xcoord string
  	  fmt.Scanln(&Xcoord)

      if Xcoord, err := strconv.Atoi(Xcoord); err == nil{
        moveJohnSnow(Xcoord)
        for (Xcoord > board.maxX || Xcoord < board.minX){
          fmt.Println()
          fmt.Println("X coordenate that you have entered is not valid please enter an integer between 0 and 9.")
          fmt.Scanln(&Xcoord)
          moveJohnSnow(Xcoord)
        }
      }


  		if (strings.TrimSpace(Xcoord) == "shoot"){
        if( !board.arrowOnBoard ){
          board.arrow.y = 1
          board.arrow.x = board.JohnSnow.x
          board.arrowOnBoard = true
    			fmt.Println(board.arrow)
        }else{
          fmt.Println()
          fmt.Println("Max number of arrows reached!")
          fmt.Println()
        }
      }
  	}
  }

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

func init_board(){

	board = Board{false, false, 9, 29, 0, 0, JohnSnow{x: 5,y: 0 }, NightKing{ x: rand.Intn(10) ,y: 30 }, arrow{0, 0}}


  fmt.Println("Winter is Coming")
  fmt.Println("")
  fmt.Println("Rules:")
  fmt.Println("1: 10x30 board")
  fmt.Println("2: the max number of arrows on the board at one time is 1")
  fmt.Println("3: type shoot to shoot arrow")
  fmt.Println("4: type number to move John Snow")
  fmt.Println("5: Night King will move once every 5 seconds")
  fmt.Println("6: Arrows move 1 square every second in a straight line")
  fmt.Println(" Please enter integer X coordenate between 0 and 9 that John  will shoot arrow from the wall.")
}


func update_board(){
  boardPointer := &board
	go update_arrow(boardPointer)
	go move_night_king(boardPointer)
}

func update_arrow(board *Board){
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
      print_board()
    }

    if (board.arrow.x == board.NightKing.x && board.arrow.y == board.NightKing.y){
      game_over(true)
    }
  	}
  }

func print_board(){
  fmt.Println()
  fmt.Println(board.NightKing)
  fmt.Println(board.arrow)
  fmt.Println(board.JohnSnow)
  fmt.Println()
}


func move_night_king(board *Board){
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
      print_board()
    }
		if (board.NightKing.y == 0){
			game_over(false)
		}
	}
}

func game_over(win bool){
  var message string
  if win{
      message = "      The Night King is dead you have won!"
  }else{
    message = "              GAME OVER!               "
  }

  i := 0
	for ( i < 10) {
			fmt.Println()
			if (i == 5){
				fmt.Println(message)
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
