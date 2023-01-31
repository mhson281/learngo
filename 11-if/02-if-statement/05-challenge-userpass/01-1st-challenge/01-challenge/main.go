// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import ("os" ; "fmt")

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

type users struct {
    username, password string
}

const (
    username = "jack"
    password = "1888"
    usage = `Usage: [username] [password]`
    wrongUser = "Access denied for %q.\n"
    wrongPass = "Invalid password for %q.\n"
    acceptAccess = "Access granted to %q.\n"
)

func main() {
    argsLength := len(os.Args)
    args := os.Args

    if argsLength != 3 {
        fmt.Println(usage)
        return
    }


    u, p := args[1], args[2]

    user1 := users{u, p}

    if user1.username != username {
        fmt.Printf(wrongUser, user1.username)
    } else if user1.password != password {
        fmt.Printf(wrongPass, user1.password)
    } else {
        fmt.Printf(acceptAccess, user1.username)
    }
}
