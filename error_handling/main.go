package main

import (
	"fmt"
	// "errors"
)

// // MyError represents `MyError` struct
// type MyError struct {}

// // struct implements `Error` method
// func (MyError *MyError) Error() string {
// 	return "Something unexpected happened!"
// }

// func main() {
// 	// create error
// 	// MyError := &MyError{}
// 	 myErr := errors.New("Something unexpected happened!")

// 	// print error message
// 	// fmt.Println(myErr)
// 	fmt.Printf("Type of myErr is %T \n", myErr)
// 	fmt.Printf("Value of myErr is %#v \n", myErr)
// }

// // divide two number
// func Divide(a int, b int) (int, error) {
// 	// can not divide by `0`
// 	if b == 0 {
// 		return 0, errors.New("Can not divide by Zero!")
// 	} else {
// 		return (a / b), nil
// 	}
// }

// func main() {
// 	// divide 4 by 0
// 	if result, err := Divide(4, 0); err != nil {
// 		fmt.Println("Error occured: ", err)
// 	} else {
// 		fmt.Println("4/0 is", result)
// 	}
// }

// HttpError represents `HttpError` struct
// type HttpError struct {
// 	status	int
// 	method 	string
// }

// // HttpError implements Error method
// func (HttpError *HttpError) Error() string {
// 	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.", HttpError.method, HttpError.status)
// }

// // getUserEmail represents mock function: sends HTTP request
// func getUserEmail(userId int) (string, error) {
// 	// request failed, return an error
// 	return "", &HttpError{403, "GET"}
// }

// func main() {
// 	// get user email address
// 	if email, err := getUserEmail(1); err != nil {
// 		fmt.Println(err) // print error

// 		// if user is unauthorized, perform session cleaning
// 		if errVal := err.(*HttpError); errVal.status == 403 {
// 			fmt.Println("Performing session cleaning...")
// 		}
// 	} else {
// 		// do something with the `email`
// 		fmt.Println("User email is", email)
// 	}
// }

// // NetworkError represents `NetworkError` struct
// type NetworkError struct {}

// // `NetworkError` struct implement error interface
// func (e *NetworkError) Error() string {
// 	return "A network connection was aborted."
// }

// // FileSaveFailedError represents `FileSaveFailedError` struct
// type FileSaveFailedError struct {}

// // `FileSaveFailedError` implement error interface
// func (e *FileSaveFailedError) Error() string {
// 	return "The requested file could not be saved."
// }

// // a function that can return either of the above errors
// func saveFileToRemote() error {
// 	result := 2 // mock result of save operation

// 	if result == 1 {
// 		return &NetworkError{}
// 	} else if result == 2 {
// 		return &FileSaveFailedError{}
// 	} else {
// 		return nil
// 	}
// }

// func main() {
// 	// check type
// 	switch err := saveFileToRemote(); err.(type) {
// 	case nil:
// 		fmt.Println("File successfully saveed")
// 	case *NetworkError:
// 		fmt.Println("Network Error:", err)
// 	case *FileSaveFailedError:
// 		fmt.Println("File save error:", err)
// 	}
// }

// -- saving original error as a context (wrapping an error)

// // UnauthorizedError represents simple unauthorized error
// type UnauthorizedError struct {
// 	UserId 			int
// 	OriginalError	error
// }

// // add some context to the original error message
// func (httpErr *UnauthorizedError) Error() string {
// 	return fmt.Sprintf("User unauthorized Error: %v", httpErr.OriginalError)
// }

// // mock function call to validate user, returns error
// func validateUser( userId int) error {

// 	// mock general error from a function call: getSession(userId)
// 	err := fmt.Errorf("Session invalid for user id %d", userId)

// 	// return UnauthorizedError with original error
// 	return &UnauthorizedError{userId, err}
// }

// func main() {
	
// 	// validate user with id `1`
// 	err := validateUser(1)

// 	if err != nil {
// 		fmt.Print(err)
// 	} else {
// 		fmt.Println("User is allowed to perform this action!")
// 	}
// }

// -- now use struct promotion

// // UnauthorizedError represents simple unauthorized error
// type UnauthorizedError struct {
// 	UserId	int
// 	error
// }

// // mock function call to validate user, returns error
// func validateUser ( userId int ) error {

// 	// mock general error from a function call: getSession(userId)
// 	err := fmt.Errorf("Session invalid for user id %d", userId)

// 	// return UnauthorizedError with original error
// 	return &UnauthorizedError{userId, err}
// }

// func main() {

// 	// validate user with id `1`
// 	err := validateUser(1)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("User is allowed to perform this action!")
// 	}
// }

// -- Getting error context using methods

// // UnauthorizedError represents simple user unauthorized error
// type UnauthorizedError struct {
// 	UserId		int
// 	SessionId	int
// 	error
// }

// // check if user is logged in
// func (err *UnauthorizedError) isLoggedIn() bool {
// 	return err.SessionId != 0 // SessionId is 0 for non-logged in users
// }

// // mock function call to validate user, returns error
// func validateUser(userId int) error {

// 	// mock general error from a function call: getSession(userId)
// 	err := fmt.Errorf("Session invalid for user id %d", userId)

// 	// return UnauthorizedError with original error
// 	return &UnauthorizedError{userId, 18783664, err}
// }

// func main() {
// 	// validate user with id `1`
// 	err := validateUser(1)

// 	if err != nil {

// 		// extract error object from `err` interface
// 		if errVal, ok := err.(*UnauthorizedError); ok {
// 			fmt.Println(errVal)
// 			fmt.Println("Is user logged in:", errVal.isLoggedIn())
// 		} else {
// 			fmt.Println(err)
// 		}

// 	} else {
// 		fmt.Println("User is allowed to perform this action!")
// 	}
// }

// -- Custom error interfaces

// UserSessionState represents `UserSessionState` interface
type UserSessionState interface {
	error
	isLoggedIn() bool
	getSessionId() int
}

// UnauthorizedError represents simple unauthorized error
type UnauthorizedError struct {
	UserId		int
	SessionId	int
}

// check if user is logged in
func (err *UnauthorizedError) isLoggedIn() bool {
	return err.SessionId != 0 // SessionId is 0 for non-logged in users
}

// get user session id
func (err *UnauthorizedError) getSessionId() int {
	return err.SessionId
}

// return error message
func (httpErr *UnauthorizedError) Error() string {
	return fmt.Sprintf("User with id %v unauthorized to perform this action", httpErr.UserId)
}

// mock function call to validate user, returns error
func validateUser(userId int) error {
	// return an error
	return &UnauthorizedError{userId, 18783664}
}

func main() {
	// validate user with id `1`
	err := validateUser(1)

	if err != nil {

		// print error
		fmt.Println(err)

		// check if error is of `UserSessionState` interface type
		if errVal, ok := err.(UserSessionState); ok {
			if errVal.isLoggedIn() {
				fmt.Printf("Cleaning user session with session id %v\n", errVal.getSessionId() )

				// to get nested struct
				errValNested, _ := err.(*UnauthorizedError)
				fmt.Println(errValNested)
			}
		}
	} else {
		fmt.Println("User is allowed to perform this action!")
	}
}