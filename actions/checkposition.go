package actions

import (
	"fmt"
	"gamecorehalloween/inputs"
	"gamecorehalloween/models"
)

//Checkposition checks current positiion in the environment against the map
func Checkposition(cp models.Player) string {
	roomstats := models.House()
	s := roomstats[cp.Position].Name
	return s
}

//Monstercheck checks the new player position and starts battle loop
func Monstercheck(cp models.Player) models.Player {
	cm := models.Spawnmonsterget()
	if cp.Position == cm.Position {
		bcheck := 1 //1 means fight, 0 means stop
		if cm.Met == true {
			fmt.Println(cm.Found)
		}
		if cm.Met == false {
			fmt.Println("You see a", cm.Name)
			fmt.Println(cm.Intro)
			cm.Met = true
			models.Monsterupdate(cm, cm.Name)
		}
		for bcheck == 1 { //need to check if player is alive, monster is alive, and they are in the same room
			cp, cm = battle(cp, cm)
			if cm.Position == cp.Position {
				models.Monsterhealth(cm)
			}
			bcheck = 0
			if cm.Health > 0 && cp.Health > 0 {
				if cp.Position == cm.Position {
					bcheck = 1
				}
			}
		}
		if cm.Health <= 0 {
			cp.Continue = false
			fmt.Println(cm.Outrom)
		}
		if cp.Health <= 0 && cm.Health > 0 { //player wins ties
			cp.Continue = false
			//fmt.Println(cm.Outrop) moved to ending.go
		}
		models.Monsterupdate(cm, cm.Name)
		return cp
	}
	return cp
}

//Monsterspawn gets chosen monster and spawn it
func Monsterspawn() {
	cm := models.Chosenmonsterget()
	cm.Spawn = true
	models.Monsterupdate(cm, cm.Name)
}

//Monstertest quick test for RNG phone and monster pos
/*
func Monstertest() (string, string) {
	roomstats := models.House()
	m := models.Mpos()
	s1 := roomstats[m].Name
	p := models.Pcheck()
	s2 := roomstats[p].Name
	return s1, s2
}
*/

//Stq sets up the player
func Stq(np models.Player) models.Player {
	s := "How old are you?"
	i := inputs.Basicinput(s)
	np.Age = i
	if i < 18 {
		fmt.Println("Sorry, you are too young for this experience.\nGo play Monopoly with your babysitter.")
		np.Continue = false
		return np
	}
	//s = "What's your name?"
	//s = "What's your identifying gender?"
	np.Health = 100
	np.Charisma = 100
	np.Continue = true
	np.Position = 1
	return np
}

//Sti setsup Inventory
func Sti() {
	p := models.Pcheck()             //New random phone location
	newph := models.ItemGet("phone") //grabs phone item
	newph.Loc = p
	models.Itemupdate(newph, "phone")
}

//Stmonster setsup monster
func Stmonster() {
	mpos := models.Mpos()       //random position in house
	mchoose := models.Mchoice() // returns int for randomly chosen monster
	mmap := models.Monstermap()
	var cname string
	for _, v := range mmap {
		if v.Number == mchoose {
			cname = v.Name
		}
	}
	newm := models.Monsterget(cname)
	newm.Position = mpos
	newm.Chosen = true
	models.Monsterupdate(newm, newm.Name)
}
