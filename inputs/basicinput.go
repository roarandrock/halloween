package inputs

import (
	"fmt"
	"log"
)

//Basicinput outputs the statement and repeats it as necessary
func Basicinput(s string) int {
	var i int
	println(s)
	if _, err := fmt.Scan(&i); err != nil {
		log.Println("Input failed ", err, "\nTry again")
		println(s)
		if _, err := fmt.Scan(&i); err != nil {
			log.Fatal("Input failed ", err, "\nGiving up")
		}
	}
	return i
}
