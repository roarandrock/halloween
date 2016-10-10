/* Horror game for halloween
Longterm:
using _test.go instead of commenting out
fixing room movement so it's not limited to two rooms, like item update below, more flexible
redo item actions to be more dynamic instead of just 2, maybe an array? or a separate map of ["item"] = actions?
encumberance? carrying only one or two items? just check before carrying new items
stealth
 need to implement cleaner RNG calls
separate non-battle items from room items. just a flag in Item[]?
monster roaming - goes to adjacent rooms, can also work with a turn counter tied to the player model, how does this work with the attic?
Inventory and playerstats work, but could be better if kept in a separate file that is written and read from
relying less on range map searches, maybe templates?
charisma being useful?
damage description based on attacks?
crippling and statuses?
Author: roarandrock
*/

package main

import (
	"fmt"
	"halloween/narrative"
	"log"
)

func main() {
	player := narrative.Intro()
	check := narrative.Failcheck(player)
	for check == true {
		player = narrative.Midgame(player)
		check = narrative.Failcheck(player)
	}
	narrative.Ending(player)
	fmt.Println("Game Over at: ")
	log.Println(log.Ldate)
	fmt.Println("Happy Halloween")
}
