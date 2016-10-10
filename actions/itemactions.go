package actions

import (
	"fmt"
	"halloween/models"
)

//Useitem attempt at using items
func Useitem(item1 models.Item, action string, cp models.Player) (models.Item, models.Player) {
	monster := models.Chosenmonsterget()
	mid := monster.Number //number for chosen monster
	//For phone, but general structure should work for other items
	switch {
	case item1.Name == "phone":
		switch action {
		case item1.Action1:
			if monster.Spawn == false { //checks if monster has arrived yet, then monster specific greetings
				if mid == 1 {
					fmt.Println("you hear heavy breathing.\n\"I'm in your house\", says the man." +
						" you hang up the phone with disgust and confusion.")
				} else {
					fmt.Println("I'm in your house but have no specific details")
				}
				cp.Charisma = cp.Charisma - 10
				Monsterspawn()
			}
			if monster.Spawn == true {
				fmt.Println("there is no response. something happened to the phone line.") //will say this if player tries phone without leaving room, annoying
				cp.Charisma = cp.Charisma - 10
			}
		case item1.Action2:
			fmt.Println("with mild disgust.")
			cp.Charisma = cp.Charisma + 5
			item1.Status = true
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "umbrella":
		switch action {
		case item1.Action1:
			fmt.Println("it's pointy.")
			item1.Status = false
			item1.Loc = 20
		case item1.Action2:
			fmt.Println("bad luck. You close it again.")
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "shoes":
		switch action {
		case item1.Action1:
			fmt.Println("you're wearing new kicks.")
			item1.Status = false
			item1.Loc = 20
		case item1.Action2:
			fmt.Println("you're a bare foot ninja.")
			item1.Status = true
			item1.Loc = cp.Position
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "front door":
		switch action {
		case item1.Action1:
			if monster.Spawn == true {
				fmt.Println("the door refuses to open.")
			}
			if monster.Spawn == false {
				fmt.Println("it's an empty and poorly lit street.\nyou close the door. leaving home is overrated.")
			}
		case item1.Action2:
			fmt.Println("leaving home is overrated.")
			item1.Status = true
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "oven":
		switch action {
		case item1.Action1:
			fmt.Println("a small flame appears.")
			item1.Status = false
		case item1.Action2:
			fmt.Println("you watch the fire die.")
			item1.Status = true
		default:
			fmt.Println("it has no effect")
		}

	case item1.Name == "butter knife":
		switch action {
		case item1.Action1:
			fmt.Println("it is dull and non-threatening.")
			item1.Status = false
			item1.Loc = 20
		case item1.Action2:
			fmt.Println("you drop the dull knife.")
			item1.Status = true
			item1.Loc = cp.Position
		default:
			fmt.Println("it has no effect")
		}

	case item1.Name == "skillet":
		switch action {
		case item1.Action1:
			fmt.Println("it is your trusty skillet.")
			item1.Loc = 20
			item1.Status = false
		case item1.Action2:
			if cp.Position == 2 {
				fmt.Println("atop the oven.")
			} else {
				fmt.Println("in this room that lacks food or an oven.")
			}
			item1.Status = true
			item1.Loc = cp.Position
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "magazine":
		switch action {
		case item1.Action1:
			item1.Loc = 20
			item1.Status = false
		case item1.Action2:
			switch {
			case monster.Spawn == false:
				fmt.Println("a glossy page turner.")
			case monster.Spawn == true:
				fmt.Println("an engrossing read. you lose yourself in an article. from behind a monster approaches.")
				if monster.Number == 1 {
					fmt.Println("his laughter catches you by surprise. so does his knife splitting your throat open.")
					cp.Health = 0
				} else {
					fmt.Println("and it disappears, back into the shadows. waiting for the right moment")
				}
			}
		}
	case item1.Name == "tv":
		switch action {
		case item1.Action1:
			fmt.Println("the people and places on the screen are pretty.")
			cp.Charisma = cp.Charisma - 20
			item1.Status = false
		case item1.Action2:
			fmt.Println("the smiling people disappear. in the dark screen you see the shadow reflection of the room. did something just move behind you?")
			item1.Status = true
		}
	case item1.Name == "golf club":
		switch action {
		case item1.Action1:
			fmt.Println("it looks brand new. did you ever actually play golf?")
			item1.Status = false
			item1.Loc = 20
		case item1.Action2:
			fmt.Println("hits the ground with a dull thud")
			item1.Status = true
			item1.Loc = cp.Position
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "ladder":
		switch action {
		case item1.Action1:
			fmt.Println("the only way to reach the attic.")
			item1.Status = false
		case item1.Action2:
			switch {
			case cp.Position == 5:
				cp.Position = 8
				item1.Loc = 8
			case cp.Position == 8:
				fmt.Println("you climb down from the attic.")
				cp.Position = 5
				item1.Loc = 5
			}
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "pillow":
		switch action {
		case item1.Action1:
			fmt.Println("it is light and fluffy.")
			item1.Status = false
			item1.Loc = 20
		case item1.Action2:
			fmt.Println("harmlessly on the floor.")
			item1.Status = true
			item1.Loc = cp.Position
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "bed": // maybe implement hiding later
		switch action {
		case item1.Action1:
			switch {
			case monster.Spawn == false:
				fmt.Println("deep asleep. no nightmares stalk your slumber.")
			case monster.Spawn == true:
				fmt.Println("deep asleep. nothing can wake you.")
				if monster.Number == 1 {
					fmt.Println("not even the opening of the bedroom door. the footsteps to your bedside." +
						"not even the presence of a killer at your bedside. no, the only thing that wakes you is the knife piercing your chest.")
					cp.Health = 0
				} else {
					fmt.Println("and when you wake up, no one is there.")
				}
			}
			item1.Status = false
		case item1.Action2:
			fmt.Println("that looks good enough to sleep in.")
			item1.Status = true
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "shower": // maybe implement hiding later
		switch action {
		case item1.Action1:
			switch {
			case monster.Spawn == false:
				fmt.Println("nothing like a nice hot shower to pass the time.")
			case monster.Spawn == true:
				fmt.Println("nothing like a nice hot shower to calm your nerves.")
				if monster.Number == 1 {
					fmt.Println("the curtain jerks open. the knife goes up and down. your blood swirls down the drain." +
						"\"Did you think this would end well?\" asks your killer.")
					cp.Health = 0
				} else {
					fmt.Println("and you leave the shower feeling refreshed.")
				}
			} //could implement hiding later
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "toilet":
		switch action {
		case item1.Action1:
			fmt.Println("gross. no.")
			item1.Status = false
		case item1.Action2:
			fmt.Println("even grosser. seriously no.")
			item1.Status = true
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "hairspray":
		switch action {
		case item1.Action1:
			fmt.Println("uber hair lock. there are a multitude of little warning signs on the can.")
			item1.Loc = 20
			item1.Status = false
		case item1.Action2:
			fmt.Println("to the floor.")
			item1.Status = true
			item1.Loc = cp.Position
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "red plunger":
		switch action {
		case item1.Action1:
			fmt.Println("good ol' red. gets the job done.")
			item1.Status = false
			item1.Loc = 20
		case item1.Action2:
			fmt.Println("plop.")
			item1.Status = true
			item1.Loc = cp.Position
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "saber":
		switch action {
		case item1.Action1:
			fmt.Println("your ancestor's calvary saber. with rust and heft. attics are awesome.")
			item1.Status = false
			item1.Loc = 20
		case item1.Action2:
			fmt.Println("shunk. it sticks in the floor.")
			item1.Status = true
			item1.Loc = cp.Position
		default:
			fmt.Println("it has no effect")
		}
	case item1.Name == "bear trap":
		switch action {
		case item1.Action1:
			fmt.Println("a large, spring loaded bear trap. like a cartoon. a cartoon that can shatter leg bones.")
			item1.Status = false
			item1.Loc = 20
		case item1.Action2:
			fmt.Println("you carefully place it on the ground. you've seen this on tv, should be simple enough. like a big mousetrap." +
				"\n\n\"crunch\" metal teeth bite into your arm. the most pain. While tears, curses and blood flow, you eventually free yourself.") //funny. maybe implement later with stealth and monster roaming
			item1.Used = true
			cp.Health = cp.Health - d3
		default:
			fmt.Println("it has no effect")
		}
		//add other items
	default:
		fmt.Println("surprisingly, it does not do anything")
	}
	return item1, cp
}
