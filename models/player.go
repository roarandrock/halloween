package models

//Player defines player stats
type Player struct {
	Age      int
	Health   int
	Charisma int
	Points   int
	Continue bool
	Position int
}

//Item defines item stats
type Item struct {
	Name    string //name of the item
	Loc     int    //RoomID, on player, on monster
	Status  bool   //True for action1, false for action2 time
	Action1 string //attempt at item specific commands
	Action2 string
	Used    bool //true for used, false for not used. Can remove from game when used
}

/* room defs
var (
foyer      = Room{1, "foyer", 3, 2}
kitchen    = Room{2, "kitchen", 1, 4}
stairway   = Room{3, "stairway", 1, 5}
livingroom = Room{4, "living room", 2, 0}
uphall     = Room{5, "upstairs hall", 3, 6}
bedroom    = Room{6, "master bedroom", 5, 7}
bathroom   = Room{7, "bathroom", 6, 0}
	attic      = Room{8, "attic", 0, 0}
)
items at loc 20 are on player
*/

//items to be found and carried. Also fixed room items. can be separated in the future?
var (
	phone     = Item{"phone", 2, true, "pick up", "hang up", false}
	umbrella  = Item{"umbrella", 1, true, "carry", "open", false}
	shoes     = Item{"shoes", 1, true, "put on", "take off", false}
	fdoor     = Item{"front door", 1, true, "open", "close", false}
	oven      = Item{"oven", 2, true, "turn on", "switch off", false}
	knife     = Item{"butter knife", 2, true, "pick up", "drop", false}
	skillet   = Item{"skillet", 2, true, "take", "set down", false}
	magazine  = Item{"magazine", 4, true, "take", "read", false}
	tv        = Item{"tv", 4, true, "turn on", "turn off", false}
	golfclub  = Item{"golf club", 5, true, "pick up", "drop", false}
	ladder    = Item{"ladder", 5, true, "set up", "climb", false}
	pillow    = Item{"pillow", 6, true, "take", "set down", false}
	bed       = Item{"bed", 6, true, "nap in", "make", false}
	shower    = Item{"shower", 7, true, "use", "hide in", false}
	toilet    = Item{"toilet", 7, true, "1 in", "2 in", false}
	hairspray = Item{"hair spray", 7, true, "take", "drop", false}
	plunger   = Item{"red plunger", 7, true, "take", "drop", false}
	saber     = Item{"saber", 8, true, "take", "drop", false}
	trap      = Item{"bear trap", 8, true, "take", "place", false}

//add items, make sure to map,
)

var imap = map[string]Item{ //automate with range? i.e m[i].Name: m[i]
	"phone":        phone,
	"umbrella":     umbrella,
	"shoes":        shoes,
	fdoor.Name:     fdoor,
	oven.Name:      oven,
	knife.Name:     knife,
	skillet.Name:   skillet,
	magazine.Name:  magazine,
	tv.Name:        tv,
	golfclub.Name:  golfclub,
	ladder.Name:    ladder,
	pillow.Name:    pillow,
	bed.Name:       bed,
	shower.Name:    shower,
	toilet.Name:    toilet,
	hairspray.Name: hairspray,
	plunger.Name:   plunger,
	saber.Name:     saber,
	trap.Name:      trap,
}

//Inventory returns item map
func Inventory() map[string]Item {
	return imap
}

//Itemlocation takes current inventory and returns locations in a slice
func Itemlocation() []int {
	m := Inventory()
	a := [10]int{}
	i := 0
	for _, v := range m {
		a[i] = v.Loc
		i++
	}
	iloc := a[0:i]
	return iloc
}

//ItemGet grabs current item by name
func ItemGet(s string) Item {
	m := Inventory()
	i := m[s]
	return i
}

//Itemupdate updates items in the  inventory map
func Itemupdate(upitem Item, item string) {
	m := Inventory()
	m[item] = upitem
}

/* Possibly messes with updating items?
var imapint = map[int]Item{
	1: phone,
	2: umbrella,
	3: shoes,
}

//Inventint returns int for each item
func Inventint() map[int]Item {
	return imapint
}
*/
