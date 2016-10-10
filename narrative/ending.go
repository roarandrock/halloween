package narrative

import (
	"fmt"
	"gamecorehalloween/actions"
	"gamecorehalloween/models"
	//"log"
)

//Ending shows final outcome, returns nothing
func Ending(endp models.Player) {
	cm := models.Chosenmonsterget()
	if endp.Health <= 0 { //player is dead
		fmt.Println(cm.Outrop)
	}
	fmt.Println("Health:", endp.Health, "Charisma:", endp.Charisma, "Age:", endp.Age)
	itemlist, _ := actions.Itinroom(20)
	fmt.Println("items on corpse:", itemlist)
	endm := models.Chosenmonsterget()
	fmt.Println("Monster's health:", endm.Health)
	/* Point check
	c := endp.Points
	fmt.Println("Final Points: ", c)
	switch {
	case c < 10:
		println("Loser")
	case c < 20:
		println("Middle of the road")
	case c > 30:
		println("you win!")
	default:
		log.Fatal("Crash", log.Ldate)
	}
	*/
}
