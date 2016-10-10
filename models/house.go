package models

//Room defines the rooms
type Room struct {
	RoomID int
	Name   string
	// Stuff1 string no longer needed
	//Stuff2 string
	Adj1 int //RoomIds adjacent to
	Adj2 int
	//Adj3   int
	//Adj4   int
}

/*RoomIDs
Foyer 1
Kitchen 2
Stairs 3
Living Room 4
Upper Landing 5
Bedroom 6
Bathroom 7
Attic 8 - reach through ladder item? no adjancey
wall 0
*/

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

var roomstats = map[int]Room{
	foyer.RoomID:      foyer,
	kitchen.RoomID:    kitchen,
	stairway.RoomID:   stairway,
	livingroom.RoomID: livingroom,
	uphall.RoomID:     uphall,
	bedroom.RoomID:    bedroom,
	bathroom.RoomID:   bathroom,
	attic.RoomID:      attic,
}

//House returns room details
func House() map[int]Room {
	return roomstats
}

//Adj checks adjancey
func Adj(cp Player) (string, string, int, int) {
	roomstats := House()
	id1 := roomstats[cp.Position].Adj1
	id2 := roomstats[cp.Position].Adj2
	var s1, s2 string
	var p1, p2 int
	if id1 != 0 {
		s1 = roomstats[id1].Name
		p1 = roomstats[id1].RoomID
	}
	if id2 != 0 {
		s2 = roomstats[id2].Name
		p2 = roomstats[id2].RoomID
	}
	if id1 == 0 {
		s1 = "wall"
		p1 = roomstats[cp.Position].RoomID
	}
	if id2 == 0 {
		s2 = "wall"
		p2 = roomstats[cp.Position].RoomID
	}
	return s1, s2, p1, p2
}
