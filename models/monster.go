package models

import "fmt"

//Monster defines stats
type Monster struct {
	Name     string
	Health   int
	Charisma int
	Position int
	Spawn    bool //exists in the house
	Chosen   bool //exists in the game
	Number   int  //unique number for selecting chosen
	Intro    string
	Outrom   string //for monster death
	Outrop   string //for player death
	Met      bool   //for after player meets
	Found    string
}

var (
	i1 = "a man with glistening skin and black hair that covers but cannot hide his eyes.\n" +
		"eyes that never blink as they follow you.\nhe approaches you slowly. knife in hand." +
		"\n\"Pleasure to carve you\""
	om1 = "he falls to the floor. he doesn't get up. he was only a man. and you defeated him."
	op1 = "he watches you collapsing to the floor and bleeding out. \ndid you ever stand a chance against such a killer?" +
		"\nhe cleans his knife. who is next?"
	f1 = "\"Found you\" says the knife-wielding fiend"
)

var (
	psychopath = Monster{"psychopath", 100, 75, 1, false, false, 1, i1, om1, op1, false, f1} //can bump up health later
	//evil kid
	//werewolf
	//alien
)

var mmap = map[string]Monster{
	psychopath.Name: psychopath,
}

//Monstermap returns map
func Monstermap() map[string]Monster {
	return mmap
}

//Monsterget grabs monster by name
func Monsterget(name string) Monster {
	m := Monstermap()
	cm := m[name]
	return cm
}

//Spawnmonsterget grabs current spawned monster
func Spawnmonsterget() Monster {
	m := Monstermap()
	a := [1]string{} //only for one monster
	i := 0
	cm := Monster{}
	for _, v := range m {
		if v.Spawn == true {
			a[i] = v.Name
			i++
		}
	}
	cm = m[a[0]]
	return cm
}

//Chosenmonsterget grabs monster chosen for game
func Chosenmonsterget() Monster {
	m := Monstermap()
	a := [1]string{} //only for one monster
	i := 0
	cm := Monster{}
	for _, v := range m {
		if v.Chosen == true {
			a[i] = v.Name
			i++
		}
	}
	cm = m[a[0]]
	return cm
}

//Monsterupdate updates items in the  inventory map
func Monsterupdate(upm Monster, name string) {
	m := Monstermap()
	m[name] = upm
}

//Monsterhealth prints strings detailing the monster's state
func Monsterhealth(cm Monster) {
	//fmt.Println(cp.Health) for testing
	h := cm.Health
	if cm.Number == 1 {
		switch {
		case h < 15:
			fmt.Println("limping and blood soaked, he is a sliver from death. nothing less is going to stop this monster.")
		case h < 30:
			fmt.Println("he is seriously damaged. he should be running for a hospital. but he does not stop.")
		case h < 70:
			fmt.Println("he has several wounds and is frustrated by your resilience. but he's not slowing down.")
		case h < 100:
			fmt.Println("you have hurt him but nothing a few bandages cannot fix.")
		case h == 100:
			fmt.Println("he appears determined, abnormally strong, and uninjured.")
		}
	} else {
		fmt.Println("this monster is inscrutable.")
	}
}
