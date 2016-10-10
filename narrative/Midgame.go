package narrative

import (
	"fmt"
	"gamecorehalloween/actions"
	"gamecorehalloween/inputs"
	"gamecorehalloween/models"
)

//Midgame is the guts of the game, trigger encounters
func Midgame(cp models.Player) models.Player {
	/* not a point game
	p := cp.Points
	//check for endgame condition
	if p > 15 {
		cp.Continue = false
		Ending(cp)
		return cp
	}
	*/
	//For testing things
	//actions.Test1()

	//check current position and return options
	loc := actions.Checkposition(cp)
	Pringcheck(cp)
	fmt.Println("you are in the ", loc)
	fmt.Println("what do you do?")
	fmt.Println("1.Look around the room\n2.Move to an adjacent room\n3.Use something in the room")
	s1 := "?"
	r1 := inputs.Basicinput(s1)
	if r1 == 1 {
		cp = actions.Lookr(cp)
	}
	if r1 == 2 {
		cp = actions.Move(cp)
	}
	if r1 == 3 {
		cp = actions.Usecheck(cp)
	}
	return cp
}

//Pringcheck checks if phone has been answered
func Pringcheck(cp models.Player) {
	phone := models.ItemGet("phone")
	m := models.Chosenmonsterget()
	if m.Spawn == false {
		if phone.Loc == cp.Position {
			fmt.Println("there is a ringing phone in this room")
			return
		}
		fmt.Println("you hear a phone ringing in another room")
	}
}
