package main


import (
	"fmt"
	"math/rand"
)

var n  nightKing
var j  johnSnow

var maxArrows int
var maxX int
var maxY int
var minX int
var minY int

func main(){

  maxArrows = 3
  maxX = 9
  maxY = 29
  minX = 0
  minY = 0



  n = nightKing{ rand.Intn(9), 0 }
  j = johnSnow{ 5, 30 }

  fmt.Println("Winter is Coming")
  fmt.Println("")
  fmt.Println("Rules:")
  fmt.Println("1: 10x30 board")
  fmt.Println("2: the max number of arrows on the board at one time is 3")
  fmt.Println("3: type shoot to shoot arrow")
  fmt.Println("4: type number to move John Snow")
  fmt.Println("5: Night King will move once every 5 seconds")
  fmt.Println("6: Arrows move 1 square every second in a straight line")



  fmt.Println(" Please enter integer X coordenate between 0 and 9 that John  will shoot arrow from the wall.")
  var Xcoord string
  fmt.Scanln(&Xcoord)
  if Xcoord, err := strconv.Atoi(Xcoord); err == nil{
    for (Xcoord > maxX || Xcoord < minX){
      fmt.Println("X coordenate that you have entered is not valid please enter an integer between 0 and 9.")
      fmt.Scanln(&Xcoord)
    }
  }

  moveJohnSnow(Xcoord)


  fmt.Println( j )

  fmt.Println( n )


}

func moveJohnSnow(x int){
  j.x = x

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
