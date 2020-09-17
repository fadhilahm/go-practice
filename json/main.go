// --- encoding json

package main

import (
	"fmt"
	"strings"
	"encoding/json"
	// "bytes"
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

// Profile declares `Profile` structure
// type Profile struct {
// 	Username		string
// 	followers		string
// }

// // Student declares `Student` structure
// type Student struct {
// 	FirstName, lastName		string
// 	HeightInMeters			float64
// 	IsMale					bool
// 	Languages				[2]string
// 	Subjects				[]string
// 	Grades					map[string]string
// 	Profile					Profile	
// }

// func main() {

// 	// some JSON data
// 	data := []byte(`
// 	{
// 		"FirstName": "John",
// 		"HeightInMeters": 1.75,
// 		"IsMale": null,
// 		"Languages": [ "English", "Spanish", "German" ],
// 		"Subjects": [ "Math", "Science" ],
// 		"Grades": { "Math": "A" },
// 		"Profile": {
// 			"Username": "johndoe91",
// 			"Followers": 1975
// 		}
// 	}`)

// 	// create data container
// 	var john Student = Student{
// 		IsMale: true,
// 		Subjects: []string{ "Art" },
// 		Grades: map[string]string{ "Science": "A+"},
// 	}

// 	// unmarshal `data`
// 	fmt.Printf("Error: %v\n", json.Unmarshal(data, &john))

// 	// print `john` struct
// 	fmt.Printf("%#v\n", john)
// }

// -- pointer

// Profile declares `Profile` structure
// type Profile struct {
// 	Username	string
// 	Followers	int
// }

// // Student declares `Student` structure
// type Student struct {
// 	FirstName, lastName		string
// 	HeightInMeters			float64
// 	IsMale					bool
// 	Languages				[2]string
// 	Subjects				[]string
// 	Grades					map[string]string
// 	Profile					*Profile
// }

// func main() {
// 	// some JSON data
// 	data := []byte(`
// 	{
// 		"FirstName": "John",
// 		"HeightInMeters": 1.75,
// 		"IsMale": null,
// 		"Languages": [ "English" ],
// 		"Subjects": [ "Math" , "Science"],
// 		"Grades": null,
// 		"Profile": { "Followers": 1975 }
// 	}`)

// 	// create data container
// 	var john Student = Student{
// 		IsMale: true,
// 		Languages: [2]string{ "Korean", "Chinese" },
// 		Subjects: nil,
// 		Grades: map[string]string{ "Math": "A" },
// 		Profile: &Profile{ Username: "johndoe91" },
// 	}

// 	// unmarshall `data`
// 	fmt.Printf("Error: %v\n\n", json.Unmarshal( data, &john ))
	
// 	// print `john` struct
// 	fmt.Printf( "%#v\n\n", john)
// 	fmt.Printf( "%#v\n", john.Profile)
// }

// --- promoted fields

// Profile declares `Profile` structure
// type Profile struct {
// 	Username 	string
// 	Followers	int
// }

// // Account declares `Account` structure
// type Account struct {
// 	IsMale bool
// 	Email string
// }

// // Student declares `Student` structure
// type Student struct {
// 	FirstName, lastName string
// 	HeightInMeters		float64
// 	IsMale				bool
// 	Profile
// 	Account
// }

// func main() {
// 	// some JSON data
// 	data := []byte(`
// 	{
// 		"FirstName": "John",
// 		"HeightInMeters": 1.75,
// 		"IsMale": true,
// 		"Username": "johndoe91",
// 		"Followers": 1975,
// 		"Account": { "IsMale": true, "Email": "john@doe.com"}
// 	}`)

// 	// create a data container
// 	var john Student

// 	// unmarshall `data`
// 	fmt.Printf( "Error: %v\n", json.Unmarshal( data, &john ) )

// 	// print `john` struct
// 	fmt.Printf( "%#v\n", john )
// }

// --- using structure tags

// Profile declares	`Profile` structure
// type Profile struct {
// 	Username	string	`json:"uname"`
// 	Followers	int		`json:"f_count"`
// }

// // Student declares `Student` structure
// type Student struct {
// 	Username		string		`json:"fname"`
// 	LastName		string		`json:"-"` // discard
// 	HeightInMeters	float64		`json:"height"`
// 	IsMale			bool		`json:"male"`
// 	Languages		[]string	`json:",omitempy"`
// 	Profile			Profile		`json:"profile"`
// }

// func main() {
// 	// some JSON data
// 	data := []byte(`
// 	{
// 		"fname": "John",
// 		"LastName": "Doe",
// 		"height": 1.75,
// 		"IsMale": true,
// 		"Languages": null,
// 		"profile": {
// 			"uname": "johndoe91",
// 			"Followers": 1975
// 		}
// 	}`)

// 	// create a data container
// 	var john Student = Student{
// 		Languages: []string{ "English", "French" },
// 	}

// 	// unmarshall `data`
// 	fmt.Printf("Error: %v\n", json.Unmarshal( data, &john ))

// 	// print `john` struct
// 	fmt.Printf( "%#v\n", john)
// }

// -- working with maps

// Student declares `Student` map
// type Student map[string]interface{}

// func main() {
// 	// some JSON data
// 	data := []byte(`
// 	{
// 		"id": 123,
// 		"fname": "John",
// 		"height": 1.75,
// 		"male": true,
// 		"languages": null,
// 		"subjects": [ "Math", "Science" ],
// 		"profile": {
// 			"uname": "johndoe91",
// 			"f_count": 1975
// 		}
// 	}`)

// 	// create a data container
// 	var john Student

// 	// unmarshall `data`
// 	fmt.Printf("Error: %v\nn", json.Unmarshal( data, &john ))

// 	// print `john` map
// 	fmt.Printf("%#v\n\n", john)

// 	// iterate through keys and values
// 	i := 1;
// 	for k, v := range john {
// 		fmt.Printf("%d: key ['%T']'%v', value ('%T')'%#v'\n", i, k, k, v, v)
// 		i++;
// 	}
// }

// --- even more complex data

// func main() {
// 	// some JSON data
// 	data := []byte(`
// 	{
// 		"id": 123,
// 		"fname": "John",
// 		"height": 1.75,
// 		"male": true,
// 		"languages": null,
// 		"subjects": [ "Math", "Science" ],
// 		"profile": {
// 			"uname": "johndoe91",
// 			"f_count": 1975
// 		}
// 	}`)

// 	// create a data container
// 	var john interface{}
// 	fmt.Printf( "Before: `type` of `john` is %T and its `value` is %v\n", john, john)

// 	// unmarshall `data`
// 	fmt.Printf( "Error: %v\n", json.Unmarshal(data, &john ))
// 	fmt.Printf( "After: `type` of `john` is %T\n\n", john )

// 	// print `john` map
// 	fmt.Printf( "%#v\n", john )

// 	// extract the concrete value
// 	johnData := john.(map[string]interface{})

// 	fmt.Println(johnData)
// }

// --- using Unmarshaller and TextUnmarshaller

// Profile declares `Profile` structure
// type Profile struct {
// 	Username	string
// 	Followers	string
// }

// // UnmarshalJSON - implement Unmarshaler interface
// func ( p *Profile ) UnmarshalJSON( data []byte ) error {

// 	// unmarshall JSON
// 	var container map [string]interface{}
// 	_ = json.Unmarshal( data, &container )
// 	fmt.Printf("container: %T / %#v\n\n", container, container)

// 	// extract interface values
// 	iuserName, _ := container["Username"]
// 	ifollowers, _ := container["f_count"]
// 	fmt.Printf( "iusername: %T/%#v\n", iuserName, iuserName )
// 	fmt.Printf( "ifollowers: %T/%#v\n\n", ifollowers, ifollowers)

// 	// extract concrete values
// 	userName, _ := iuserName.(string) // get `string` value
// 	followers, _ := ifollowers.(float64) // get `float64` value
// 	fmt.Printf("userName: %T/%#v\n", userName, userName)
// 	fmt.Printf("followers: %T/%#v\n\n", followers, followers)

// 	// assign values
// 	p.Username = strings.ToUpper(userName)
// 	p.Followers = fmt.Sprintf("%.2fk", followers / 1000)

// 	return nil
// }

// // Student declares `Student` structure
// type Student struct {
// 	FirstName	string
// 	Profile		Profile
// }

// func main() {
// 	// some JSON data
// 	data := []byte(`
// 	{
// 		"FirstName": "John",
// 		"Profile": {
// 			"Username": "johndoe91",
// 			"f_count": 1975
// 		}
// 	}`)

// 	// create a data container
// 	var john Student

// 	// unmarshall `data`
// 	fmt.Printf("Error: %#v\n", json.Unmarshal( data, &john ))

// 	// print `john` struct
// 	fmt.Printf("%#v\n", john)
// }

// --- encoder and decoder

// Person declares `Person` structure
// type Person struct {
// 	Name 	string
// 	Age		int
// }

// func main() {
// 	// create a buffer to hold JSON data
// 	buf := new(bytes.Buffer)
// 	// create JSON encode for `buf`
// 	bufEncoder := json.NewEncoder(buf)

// 	// encode JSON from `Person` struct
// 	bufEncoder.Encode(Person{"Ross Geller", 28})
// 	bufEncoder.Encode(Person{"Monica Geller", 27})
// 	bufEncoder.Encode(Person{"Jack Geller", 56})

// 	// print contents of the `buf`
// 	fmt.Println(buf) // calls `buf.String()` method
// }

// --- decoder

// Person declares `Person` structure
type Person struct {
	Name	string
	Age		int
}

func main() {
	// create a strings reader
	jsonStream := strings.NewReader(`
	{"Name": "Ross Geller", "Age": 28}
	{"Name": "Monica Geller", "Age": 27}
	{"Name": "Jack Geller", "Age": 56}
	`)

	// create JSON decoder using `jsonStream`
	decoder := json.NewDecoder(jsonStream)

	// create `Person` struct to hold decoded data
	var ross, monica Person

	// decode JSON from `decoder` one line at a time
	decoder.Decode(&ross)
	decoder.Decode(&monica)

	// see value of the `ross` and `monica`
	fmt.Printf("ross: %#v\n", ross)
	fmt.Printf("monica: %#v\n", monica)
}