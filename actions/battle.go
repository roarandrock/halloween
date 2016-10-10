package actions

import (
	"fmt"
	"halloween/inputs"
	"halloween/models"
)

var ( //for tweaking battles, light medium and heavy damage
	d1 = 10
	d2 = 30
	d3 = 70
)

//Checkposition checks current position in the environment against the map
func battle(cp models.Player, cm models.Monster) (models.Player, models.Monster) {
	stuff, q := Itinroom(cp.Position) //array of items (by name) and quanitity
	fmt.Println("1.Use item\n2.Talk\n3.Run")
	s1 := "?"
	switch r1 := inputs.Basicinput(s1); {
	case r1 == 1:
		if q == 0 {
			fmt.Println("there's nothing here you can use. and the monster knows that. he punishes you for your mistake.")
			cp.Health = cp.Health - d1
			return cp, cm
		}
		fmt.Println("which item?")
		for i := 0; i < q; i++ {
			fmt.Println(i+1, stuff[i])
		}
		r2 := inputs.Basicinput(s1)
		iname := stuff[r2-1]
		item1 := models.ItemGet(iname)
		cp, cm = Useitembattle(item1, cp, cm)
	case r1 == 2:
		fmt.Println("1.Hello\n2.Who are you?\n3.Why are you doing this?\n4.I'm calling the police")
		inputs.Basicinput(s1) //no need for a response now
		if cm.Number == 1 {   //for pyscho, he doesn't care to talk
			fmt.Println("He just smiles and cuts you with a knife")
			cp.Health = cp.Health - d2
		}
	case r1 == 3:
		fmt.Println("you run away")
		cp = Move(cp)
		if cm.Number == 1 {
			fmt.Println("the ", cm.Name, " lunges and slashes. you are cut.")
			cp.Health = cp.Health - d1
			if cp.Health <= 0 {
				return cp, cm
			}
			fmt.Println("but he does not follow you.")
			cm = Mrun(cp, cm)
		}
	}
	return cp, cm
}

//Useitembattle is a giant switch for using different items in battle. Defaults to no effect
func Useitembattle(item1 models.Item, cp models.Player, cm models.Monster) (models.Player, models.Monster) {

	switch {
	case item1.Name == "phone":
		fmt.Println("you throw the phone at the ", cm.Name)
		if cm.Number == 1 {
			fmt.Println("the busted phone falls to the floor.\nhe is bruised but undeterred. he cuts you.")
			cm.Health = cm.Health - d1
			cp.Health = cp.Health - d2
			item1.Used = true
		} else {
			fmt.Println(item1.Name, "has not impact on this beast.")
		}
	case item1.Name == "umbrella":
		if cm.Number == 1 {
			if item1.Loc == 20 {
				fmt.Println("umbrella attack!\nit keeps him at bay.\nhe cannot cut you and it frustrates him. he runs away.")
				cm = Mrun(cp, cm)
			} else {
				fmt.Println("you go for the umbrella. he slices you with a knife. but you got an umbrella now")
				item1.Loc = 20
				cp.Health = cp.Health - d1
			}
		}
	case item1.Name == "shoes":
		if item1.Loc == 20 {
			fmt.Println("You kick with your kicks")
			if cm.Number == 1 {
				fmt.Println("He is thrown back and hurt. frustrated he runs out of the room. sadly and arbitrarily, your shoes are ruined.")
				cm.Health = cm.Health - d2
				item1.Used = true
				cm = Mrun(cp, cm)
			}
		} else {
			fmt.Println("You throw the shoes at the ", cm.Name)
			if cm.Number == 1 {
				fmt.Println("He is bruised but undeterred. He cuts you.")
				cm.Health = cm.Health - d1
				cp.Health = cp.Health - d1
			}
		}
	case item1.Name == "oven":
		if cm.Number == 1 {
			if item1.Status == false {
				fmt.Println("he sticks his knife in your hand.\nyou take this painfully earned opportunity to stick his arm on the open flame.")
				cm.Health = cm.Health - d2
				cp.Health = cp.Health - d2
				item1.Used = true
			}
			if item1.Status == true {
				fmt.Println("you turn on the stove and a flame appears. while doing this, he cuts you with a knife.")
				item1.Status = false
				cp.Health = cp.Health - d2
			}
		} else {
			fmt.Println(item1.Name, "has not impact on this beast.")
		}
	case item1.Name == "butter knife":
		if cm.Number == 1 {
			if item1.Status == false {
				fmt.Println("you have a butter knife. you are not sure what kind of knife he has. you are not good with knives." +
					"\nthis is evident as he slashes you and your butter knife does not scratch him.")
				cp.Health = cp.Health - d1
			}
			if item1.Status == true {
				fmt.Println("you reach for the butter knife. you grab the butter knife. he cuts you with his knife.")
				item1.Status = false
				cp.Health = cp.Health - d1
			}
		} else {
			fmt.Println(item1.Name, "has not impact on this beast.")
		}
	case item1.Name == "skillet": //true is take action, false is set down action
		oven := models.ItemGet("oven")
		if cm.Number == 1 {
			switch {
			case item1.Status == true:
				if cp.Position == 2 && oven.Status == false {
					fmt.Println("you take the hot skillet from the oven and slam it in the monster's face.\nhe shrieks and runs out of the room. the smell of burnt toast lingers.")
					cm.Health = cm.Health - d3
					cm = Mrun(cp, cm)
					item1.Used = true
				}
				if cp.Position == 2 && oven.Status == true {
					fmt.Println("you take the skillet from the cold oven. his knife clangs against the cold metal, saving you from harm.\nfrustrated he flees the room.")
					item1.Status = true
					cm = Mrun(cp, cm)
				}
				if cp.Position != 2 {
					fmt.Println("you take the skillet from the floor. his knife clangs against the cold metal, saving you from harm.\nfrustrated he flees the room.")
					item1.Status = true
					cm = Mrun(cp, cm)
				}
			case item1.Status == false:
				if cp.Position == 2 {
					fmt.Println("1.Place the skillet on the oven\n2.Hit him with it")
					s1 := "?"
					r1 := inputs.Basicinput(s1)
					if r1 == 1 {
						fmt.Println("you place the skillet on the oven. the man with the knife cuts you.")
						cp.Health = cp.Health - d2
						item1.Status = true
					} else {
						fmt.Println("you hit the man with the cold metal pan. bruised but he slashes you with his cold sharp knife.")
						cm.Health = cm.Health - d1
						cp.Health = cp.Health - d1
					}
				}
				if cp.Position != 2 {
					fmt.Println("you hit the man with the cold metal pan. he is bruised but slashes you with his cold sharp knife.")
					cm.Health = cm.Health - d1
					cp.Health = cp.Health - d1
				}
			}
		} else {
			fmt.Println(item1.Name, "has not impact on this beast.")
		}
	case item1.Name == "magazine":
		if cm.Number == 1 {
			fmt.Println("with luck you catch the quick knife with your glossy shield.\nyou are spared but the magazine is destroyed." +
				"\nthe monster mutters something about icons and disappears into a different room.")
			item1.Used = true
			cm = Mrun(cp, cm)
		} else {
			fmt.Println(item1.Name, "has not impact on this beast.")
		}
	case item1.Name == "tv":
		if cm.Number == 1 {
			fmt.Println("the sudden sounds and images of people seem to confuse and frustrate him. he attacks the television." +
				"\nhis strength and ferocity makes short work of the appliance. now his attention is on you.")
			item1.Used = true
		} else {
			fmt.Println(item1.Name, "has not impact on this beast.")
		}
	case item1.Name == "golf club":
		if cm.Number == 1 {
			if item1.Status == false {
				fmt.Println("you swing the club and are rewarded with a solid thwack. it hits the monster on his face." +
					"\nhe is hurt and furious. he grabs the club away from you and breaks it across his knee.")
				cm.Health = cm.Health - d2
				item1.Used = true
			}
			if item1.Status == true {
				fmt.Println("you race down the hall for the golf club. the villain stares at you, waiting for your next move.\nis he scared?")
				item1.Status = false
				item1.Loc = 20
			}
		} else {
			fmt.Println(item1.Name, "has not impact on this beast.")
		}
	case item1.Name == "ladder":
		if cm.Number == 1 {
			fmt.Println("you race towards the ladder but the monster gets there first. he deters your climb by cutting your skin.")
			cp.Health = cp.Health - d1
		} else {
			fmt.Println("now is not the time.")
		} //only issue is trapping player in the attic? need to implement something for there.
	case item1.Name == "pillow":
		if cm.Number == 1 {
			if item1.Status == false {
				fmt.Println("you cower behind your pillow shield. it does a terrible job as a shield. the maniac slashes through it." +
					"like a knife through a pillow. you are wounded.")
				cp.Health = cp.Health - d2
				item1.Used = true
			}
			if item1.Status == true {
				fmt.Println("you grab the pillow and cower behind it. it does a terrible job as a shield. the maniac slashes through it." +
					"like a knife through a pillow. you are wounded.")
				cp.Health = cp.Health - d2
				item1.Used = true
			}
		} else {
			fmt.Println(item1.Name, "has not impact on this beast.")
		}
	case item1.Name == "bed":
		if cm.Number == 1 {
			fmt.Println("sorry you cannot sleep now. and you're not going to wake up, safe in bed at the end of this. to remind you of this, he cuts you.")
			cp.Health = cp.Health - d2
			item1.Used = true
		} else {
			fmt.Println("now is not the time.")
		}
	case item1.Name == "shower":
		if cm.Number == 1 {
			fmt.Println("you try to hide in the shower. he is not fooled. his knife breaks through the curtain and cuts you.")
			cp.Health = cp.Health - d1
			item1.Used = true
		} else {
			fmt.Println("now is not the time.")
		}
	case item1.Name == "toilet":
		if cm.Number == 1 {
			fmt.Println("gross. no.") //something better for toilet?
		} else {
			fmt.Println("now is not the time.")
		}
	case item1.Name == "hair spray":
		oven := models.ItemGet("oven")
		if cm.Number == 1 {
			switch {
			case item1.Status == false:
				if cp.Position == 2 && oven.Status == false {
					fmt.Println("the jet of hairspray hits the flame from the oven. the combustion throws the monster across the room.")
					cm.Health = cm.Health - d3
					item1.Used = true
				}
				if cp.Position == 2 && oven.Status == true {
					fmt.Println("you spray the flammable, semi-toxic aerosol in the monster's face. he cries in pain and blindly slashes. lucky bastard cuts you.")
					cp.Health = cp.Health - d1
					cm.Health = cm.Health - d1
				}
				if cp.Position != 2 {
					fmt.Println("you spray the chemical aerosol in the monster's face. he cries in pain and blindly slashes. lucky bastard cuts you.")
					cp.Health = cp.Health - d1
					cm.Health = cm.Health - d1
				}
			case item1.Status == true:
				fmt.Println("you grab your choice hairspray. he cuts you with his knife.")
				item1.Status = false
				item1.Loc = 20
				cp.Health = cp.Health - d1
			}
		} else {
			fmt.Println(item1.Name, "has not impact on this beast.")
		}
	case item1.Name == "red plunger":
		if cm.Number == 1 {
			if item1.Status == false {
				fmt.Println("plop. you stick it right on his face. you start laughing. he tears it off, taking some skin with it, and breaks it over his knee.")
				cm.Health = cm.Health - d1
				cp.Charisma = cp.Charisma + 30
				item1.Used = true
			}
			if item1.Status == true {
				fmt.Println("you grab the plunger. luckily, the wooden handle blocks his attack.")
				item1.Status = false
				item1.Loc = 20
			}
		} else {
			fmt.Println(item1.Name, "has not impact on this beast.")
		}
	case item1.Name == "saber":
		if cm.Number == 1 {
			if item1.Status == false {
				fmt.Println("rusty saber trumps knife." +
					"he is heavily wounded and flees the room.")
				cm.Health = cm.Health - d3
			}
			if item1.Status == true {
				fmt.Println("you try for the saber. the villain sees your aim and beats you there. taking it from the ground, he then runs and disappears into another room. your ancestors are disappointed.")
				item1.Used = true
				cm = Mrun(cp, cm)
			}
		} else {
			fmt.Println(item1.Name, "has not impact on this beast.")
		}
	case item1.Name == "bear trap":
		if cm.Number == 1 {
			if item1.Status == false {
				fmt.Println("his knife clangs off the metal trap. there's no time to set the trap but it's a superb shield.")
			} else {
				fmt.Println("the monster is in your way.")
			}
		} else {
			fmt.Println(item1.Name, "has not impact on this beast.")
		}
	case item1.Name == "front door":
		fmt.Println("the monster is between you and the door.")
		// add more items
	default:
		fmt.Println("it has no effect. bad plan")
	}
	models.Itemupdate(item1, item1.Name)
	return cp, cm
}

//Mrun moves the monster to any location that is not where the player is
func Mrun(cp models.Player, cm models.Monster) models.Monster {
	if cm.Health == 0 {
		fmt.Println("the monster does not get far. overwhelmed he collapses")
		return cm
	}
	models.Mset()
	cm.Position = models.Mpos()
	for cm.Position == cp.Position {
		models.Mset()
		cm.Position = models.Mpos()
	}
	return cm
}
