package models

import (
  "time"
//  "gopkg.in/mgo.v2/bson"
)

type Person struct { 
	FirstName, LastName string
	Dob time.Time 
	Email, 	Location     string 
}
