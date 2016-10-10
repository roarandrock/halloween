package narrative

import (
	"fmt"
	"gamecorehalloween/models"
)

//Failcheck checks for death and zero charisma
func Failcheck(currentp models.Player) bool {
	h := currentp.Health
	if h == 0 {
		return false
	}
	c := currentp.Charisma
	if c == 0 {
		fmt.Println("You have zero charisma!")
		return false
	}
	r := currentp.Continue
	if r == false {
		return false
	}
	return true
}
