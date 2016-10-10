package models

import (
	"math/rand"
	"time"
)

var m int        //monster location
var p int        //phone location
var r int        //numbers of room
var moptions int //number of monsters
var monster int  //monster chosen

//Mset sets the monster and phone
func Mset() {
	roommap := House()
	r = len(roommap)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	m = r1.Intn(r) + 1 //for Monstercheck
	for m == 1 || m == 8 {
		m = r1.Intn(r) + 1
	}
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	p = r2.Intn(r) + 1
	for p == m || p == 8 || p == 7 {
		p = r2.Intn(r) + 1
	}
	mmap := Monstermap() // choosing monster
	moptions = len(mmap)
	s3 := rand.NewSource(time.Now().UnixNano())
	r3 := rand.New(s3)
	monster = r3.Intn(moptions) + 1
}

//Mpos returns roomid for monster
func Mpos() int {
	return m
}

//Pcheck checks if phone is in the room
func Pcheck() int {
	return p
}

//Mchoice returns chosen monster
func Mchoice() int {
	return monster
}
