// --- encoding json

package main

import (
	"fmt"
	"encoding/json"
)

// Student declares `Student` structure
// type Student struct {
// 	FirstName, lastName		string
// 	Email					string
// 	Age						int
// 	HeightInMeters			float64
// 	IsMale					bool
// }

// func main() {
// 	// define `john` struct
// 	john := Student{
// 		FirstName: "John",
// 		lastName: "Doe",
// 		Age: 21,
// 		HeightInMeters: 1.75,
// 		IsMale: true,
// 	}

// 	// encode `john` as JSON
// 	johnJSON, _ := json.Marshal(john)
// 	johnJSONPrettify, _ := json.MarshalIndent(john, "", "" )

// 	// print JSON string
// 	fmt.Println( string(johnJSON ))
// 	fmt.Println( string(johnJSONPrettify ))
// }

// --- encode map

// Student declares `Student` map
// type Student map[string]interface{}

// func main() {
// 	// define `john` struct
// 	john := Student{
// 		"FirstName": "John",
// 		"lastName": "Doe",
// 		"Age": 21,
// 		"HeightInMeters": 1.75,
// 		"IsMale": true,
// 	}

// 	// encode `john` as JSON
// 	johnJSON, _ := json.Marshal(john)

// 	// print JSON string
// 	fmt.Println( string(johnJSON) )

// }

// --- abstract data types

// // Profile declares `Profile` structure
// type Profile struct {
// 	Username		string
// 	followers		int
// 	Grades			map[string]string
// }

// // Student declares `Student` structure
// type Student struct {
// 	FirstName, lastName		string
// 	Age						int
// 	Profile					
// 	Languages				[]string
// }

// func main() {

// 	var john Student

// 	// define `john` struct
// 	john = Student{
// 		FirstName: 	"John",
// 		lastName: 	"Doe",
// 		Age:		21,
// 		Profile:	Profile{
// 			Username:	"johndoe91",
// 			followers:	1975,
// 			// Grades: map[string]string {"Math": "A", "Science": "A+"},
// 		}	,
// 		Languages:	[]string{"English", "French"},
// 	}

// 	// encode `john` as JSON
// 	johnJSON, err := json.MarshalIndent(john, "", "  ")

// 	// print JSON string
// 	fmt.Println( string(johnJSON), err)

// }

// --- pointer and interface

// // ProfileI interface defines `Follow` method
// type ProfileI interface {
// 	Follow()
// }

// // Profile declares `Profile` structure
// type Profile struct {
// 	Username	string
// 	Followers	int
// }

// // Follow method implementation
// func (p *Profile) Follow() {
// 	p.Followers++
// }

// // Student declares `Student` structure
// type Student struct {
// 	FirstName, lastName		string
// 	Age						int
// 	Primary					ProfileI
// 	Secondary				ProfileI
// }

// func main() {
// 	// define `john` structure (pointer)
// 	john := &Student{
// 		FirstName: 	"John",
// 		lastName:	"Doe",
// 		Age:		21,
// 		Primary:	&Profile{
// 			Username:	"johndoe91",
// 			Followers:	1975,
// 		},
// 	}

// 	// follow `john`
// 	john.Primary.Follow()

// 	// encode `john` as JSON
// 	johnJSON, err := json.MarshalIndent(john, "", "  ")

// 	// print JSON string
// 	fmt.Println( string(johnJSON), err)
// }

// --- data type conversion

// Profile declares `Profile` structure
// type Profile struct {
// 	Username	string
// 	Followers 	int
// }

// // MarshalJSON - implement `Marshaller` interface
// func (p Profile) MarshalJSON() ([]byte, error) {
// 	// return JSON value (TODO: handle error gracefully)
// 	return []byte(fmt.Sprintf(`{"f_count": "%d"}`, p.Followers)), nil
// }

// // Age declares `Age` type
// type Age int

// // MarshalText - Implement `TextMarshaler` interface
// func (a Age) MarshalText() ([]byte, error) {
// 	// return string value (TODO: handle error gracefully)
// 	return []byte(fmt.Sprintf(`{"age": %d}`, int(a))), nil
// }

// // Student declares `Student` structure
// type Student struct {
// 	FirstName, lastName		string
// 	Age						Age
// 	Profile					Profile
// }

// func main() {
	
// 	// define `john` struct (pointer)
// 	john := &Student{
// 		FirstName:	"John",
// 		lastName:	"Doe",
// 		Profile:	Profile{
// 			Username:	"johndoe91",
// 			Followers:	1975,
// 		},
// 	}

// 	// encode `john` as JSON
// 	johnJSON, err := json.MarshalIndent(john, "", "  ")

// 	// print JSON string
// 	fmt.Println( string(johnJSON), err)
// }

// --- using structure tags

// Profile declares `Profile` structure
// type Profile struct {
// 	Username 	string 	`json:"uname"`
// 	Followers	int		`json:"followers,omitempty,string"`
// }

// // Student declares `Student` structure
// type Student struct {
// 	FirstName		string		`json:"fname"` // `fname` as field name
// 	LastName		string		`json:"lname,omitempty"` // discard if value is empty
// 	Email			string		`json:"-"` // always discard
// 	Age				int			`json:"-,"` // `-` as field name
// 	IsMale			bool		`json:",string"` // keep original field name, coerce to a string
// 	Profile			Profile		`json:""` // no effect
// }

// func main() {
// 	// define `john` struct (pointer)
// 	john := &Student{
// 		FirstName: 	"John",
// 		LastName:	"", // empty
// 		Age:		21,
// 		Email:		"john@doe.com",
// 		Profile:	Profile{
// 			Username:	"johndoe91",
// 			Followers:	1975,
// 		},
// 	}

// 	// encode `john` as JSON
// 	johnJSON, _ := json.MarshalIndent(john, "", "  ")

// 	// print JSON string
// 	fmt.Println( string (johnJSON))
// }

// --- decoding JSON

// func main() {
// 	// some JSON data
// 	data := []byte(`
// 	{
// 		"FirstName": "John",
// 		"Age": 21,
// 		"Username": "johndoe91",
// 		"Grades": null,
// 		"Languages": [
// 			"English",
// 			"French"
// 		]
// 	}`)

// 	// check if `data` is a valid JSON
// 	isValid := json.Valid(data)
// 	fmt.Println(isValid)
// }

// --- Unmarshall

// Student declares `Student` structure
// type Student struct {
// 	FirstName, lastName		string
// 	Email					string
// 	Age						int
// 	HeightInMeters			float64
// }

// func main() {
// 	// some JSON data
// 	data := []byte(`
// 	{
// 		"FirstName": "John",
// 		"lastName": "Doe",
// 		"Age": 21,
// 		"HeightInMeters": 175,
// 		"Username": "johndoe91"
// 	}
// 	`)

// 	// create a data container
// 	var john Student

// 	// unmarshall `data`
// 	fmt.Printf(" Error: %v\n", json.Unmarshal(data, &john) );

// 	// print `john` struct
// 	fmt.Printf("%#v\n", john)
// }

// --- handling complex data