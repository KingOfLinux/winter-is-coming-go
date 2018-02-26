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

	init_board()
	update_board()

	for !board.gameOver{
		fmt.Println(" Please enter integer X coordenate between 0 and 9 that John  will shoot arrow from the wall.")
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

		if (strings.TrimSpace(Xcoord) == "shoot" && board.numArrows <= board.maxArrows){
			arrow := arrow{board.johnSnow.x, 29 }
			board.arrows[board.numArrows] = arrow
			fmt.Println(arrow)

		}else{
			fmt.Println()
			fmt.Println("Invalid answer!")
			fmt.Println()
		}
		fmt.Println( board.johnSnow )
	  fmt.Println( board.nightKing )
	}
}

type Board struct{
	gameOver bool
	numArrows int
	maxArrows int
	maxX int
	maxY int
	minX int
	minY int
	johnSnow
	nightKing
	arrows []arrow
}

func init_board(){

	//Max arrows set to 3
	board = Board{false, 0, 3, 9, 29, 0, 0, johnSnow{x: 5,y: 30 }, nightKing{ x: rand.Intn(9),y: 0 }, []arrow{arrow{ 0, 0 }, arrow{ 0, 0 }, arrow{ 0, 0 } }}

  fmt.Println("Winter is Coming")
  fmt.Println("")
  fmt.Println("Rules:")
  fmt.Println("1: 10x30 board")
  fmt.Println("2: the max number of arrows on the board at one time is 3")
  fmt.Println("3: type shoot to shoot arrow")
  fmt.Println("4: type number to move John Snow")
  fmt.Println("5: Night King will move once every 5 seconds")
  fmt.Println("6: Arrows move 1 square every second in a straight line")
}

func update_board(){
	go update_arrows()
	go board.moveNightKing()


}

func update_arrows(){
	i := 0
	for (i < len(board.arrows)){
		board.arrows[i].y--
		i++
	}
}

func (board Board) moveNightKing(){
	for !board.gameOver{
		time.Sleep(time.Second * 1)
		board.nightKing.y = board.nightKing.y + 1
		direction := rand.Intn(3)
		switch direction{
			case 0:{
				if (board.nightKing.x <= board.minX){
					board.nightKing.x = board.nightKing.x + 1
				}else{
					board.nightKing.x = board.nightKing.x - 1
				}
			}
			case 1:
				board.nightKing.x = board.nightKing.x
			case 2:{
				if (board.nightKing.x >= board.maxX){
					board.nightKing.x = board.nightKing.x - 1
				}else{
					board.nightKing.x = board.nightKing.x + 1
				}
			}
		}
		fmt.Println(board.nightKing)

		if (board.nightKing.y == 30){
			game_over()
			board.gameOver = true
		}
	}
}

func game_over(){
  i := 0
	for ( i < 10) {
			fmt.Println()
			if (i == 5){
				fmt.Println("              GAME OVER!               ")
			}
			i++
		}
	board.gameOver = true
}

func moveJohnSnow(x int){
  board.johnSnow.x = x
}

type nightKing struct{
	x int
	y int
}

func (n nightKing) String() string{
	return fmt.Sprintf("Night King : (%d, %d) " , n.x, n.y )
}

type arrow struct{
	x int
	y int
}

func (a arrow) String() string{
	return fmt.Sprintf("Arrow : (%d, %d) " , a.x, a.y )
}

type johnSnow struct{
	x int
	y int
}

func (j johnSnow) String() string{
	return fmt.Sprintf("John Snow: (%d, %d) " , j.x, j.y )
}
