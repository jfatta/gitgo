package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

// flags
var (
	user string
)

func main() {
	// parse flags
	flag.Parse()

	// if user does not supply flags, print usage
	// we can clean this up later by putting this into its own function
	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)

	// "for... range" loop in GO allows us to iterate over each element of the array.
	// "range" keyword can return the index of the element (e.g. 0, 1, 2, 3 ...etc)
	// and it can return the actual value of the element.
	// Since GO does not allow unused variables, we use the "_" character to tell GO we don't care about the index, but
	// we want to get the actual user we're looping over to pass to the function.
	for _, u := range users {
		result := getUsers(u)
		fmt.Println(`Username:	`, result.Login)
		fmt.Println(`Name:		`, result.Name)
		fmt.Println(`Email:		`, result.Email)
		fmt.Println(`Bio:		`, result.Bio)
		fmt.Println("")
	}
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
