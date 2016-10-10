package actions

import (
	"fmt"
	"gamecore/inputs"
	"halloween/models"
)

//Itinroom checks for items in room and on person and returns array of items
func Itinroom(rid int) ([]string, int) {
	q := 0 //number of items in room
	m := models.Inventory()
	vlist := [10]string{} //make this into a generic search function for maps?
	i := 0
	for _, v := range m {
		if v.Used == false {
			if v.Loc == rid || v.Loc == 20 {
				vlist[i] = v.Name
				i++
				q++
			}
		}
	}
	if q == 0 {
		empty := []string{}
		return empty, q
	}
	stuff := vlist[0:i]
	return stuff, q
}

//Lookr at room
func Lookr(cp models.Player) models.Player {
	loc := cp.Position
	stuff, q := Itinroom(loc)
	//fmt.Println("There are items: ", q) //test
	if q == 0 {
		fmt.Println("nothing interesting here")
		//health check
		fmt.Println("check yourself out?")
		fmt.Println("1.Yes\n2.No")
		s1 := "?"
		r1 := inputs.Basicinput(s1)
		if r1 == 1 {
			Healthcheck(cp)
		}
		return cp
	}
	fmt.Println("in this room you see:")
	for _, v := range stuff {
		fmt.Println(v)
	}
	//health check
	fmt.Println("check yourself out?")
	fmt.Println("1.Yes\n2.No")
	s1 := "?"
	r1 := inputs.Basicinput(s1)
	if r1 == 1 {
		Healthcheck(cp)
	}
	return cp
}

//Healthcheck returns strings describing player's health
func Healthcheck(cp models.Player) {
	//fmt.Println(cp.Health) for testing
	h := cp.Health
	switch {
	case h < 15:
		fmt.Println("blood and pain everywhere. having trouble keeping awake. you're hanging on by a narrow thread.")
	case h < 30:
		fmt.Println("you have never been this damaged before. you are still moving but cannot take much more punishment.")
	case h < 70:
		fmt.Println("there is blood and pain. you are not doing great. time to focus and get even.")
	case h < 100:
		fmt.Println("you are hurt but you will recover. assuming you survive the night.")
	case h == 100:
		fmt.Println("looking good. you narcissist.")
	}
}

//Usecheck items
func Usecheck(cp models.Player) models.Player {
	loc := cp.Position
	stuff, q := Itinroom(loc)
	if q == 0 {
		fmt.Println("nothing interesting here")
		return cp
	}
	fmt.Println("here you can use:")
	for i, v := range stuff {
		fmt.Println(i+1, v)
	}
	fmt.Println("choose an item")
	r1 := inputs.Basicinput("?")
	if r1 == 0 || r1 > q {
		fmt.Println("what? that's not here.")
		return cp
	}
	pick := stuff[r1-1]
	upitem := models.ItemGet(pick)
	var a string
	if upitem.Status == false {
		a = upitem.Action2
	}
	if upitem.Status == true {
		a = upitem.Action1
	}
	fmt.Println(a, "the", upitem.Name, "?")
	fmt.Println("1.Yes\n2.No")
	r2 := inputs.Basicinput("?")
	if r2 == 1 {
		fmt.Println("you", a, "the", upitem.Name)
		upitem, cp = Useitem(upitem, a, cp)
		models.Itemupdate(upitem, pick)
	}
	return cp
}

//Move options
func Move(cp models.Player) models.Player {

	adj1, adj2, p1, p2 := models.Adj(cp)
	fmt.Println("where would you like to go?")
	if adj1 != "wall" {
		fmt.Println("1.", adj1)
		if adj2 != "wall" {
			fmt.Println("2.", adj2)
		}
		r1 := inputs.Basicinput("?")
		if r1 == 1 {
			cp.Position = p1
		}
		if r1 == 2 {
			cp.Position = p2
		}
	}
	cm := models.Spawnmonsterget()
	if cm.Spawn == true {
		cp = Monstercheck(cp)
	}
	return cp
}
