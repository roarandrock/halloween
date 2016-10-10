package actions

import (
	"fmt"
	"gamecorehalloween/models"
)

//Test1 for testing random functionality in the game
func Test1() {

	//m, _ := Monstertest()
	cm := models.Chosenmonsterget()
	fmt.Println("Monster is in", cm.Position)
	if cm.Spawn == false {
		fmt.Println("Phone is in", models.ItemGet("phone").Loc)
	}
	//force room for phone and Monster
	/*
		cm.Position = 2
		models.Monsterupdate(cm, cm.Name)
		phone := models.ItemGet("phone")
		phone.Loc = 1
		models.Itemupdate(phone, phone.Name)
	*/
	/*	item := "phone"
		upitem := models.Item{}
		upitem.Loc = 1
		upitem.Name = "phone"
		upitem.Status = false
		models.Itemupdate(upitem, item)
	*/
	//i := models.Inventory()
	//fmt.Println(i["phone"])

	//r := models.House()
	//fmt.Println(r[1])

}
