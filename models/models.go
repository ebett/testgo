package models

import (
	"fmt"
	"time"
//  "gopkg.in/mgo.v2/bson"
)

type Person struct { 
	FirstName, LastName string
	Dob time.Time 
	Email, 	Location     string 
}

func (p Person) toString(){
	fmt.Printf("FirstName: %s, LastName: %s, DateOfBirth: %s", p.FirstName, p.LastName, p.Dob.String())
}
