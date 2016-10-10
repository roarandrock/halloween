package narrative

import (
	"fmt"
	"halloween/actions"
	"halloween/models"
)

// Intro launches the start of the game
func Intro() models.Player {
	np := models.Player{} //Sets up new player with basic stat
	models.Mset()         // sets up monster location
	np = actions.Stq(np)
	fmt.Println("it is late and you are home alone")
	actions.Sti() //sets up Inventory and monster stats
	actions.Stmonster()
	return np
}
