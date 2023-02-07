package main


// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------
import ("time"; "math/rand"; "strconv"; "os"; "fmt")

const maxTurns = 5

func main() {
    rand.Seed(time.Now().UnixNano())

    if len(os.Args) != 2 {
        fmt.Println("You need to take a guess what the number is!!")
        return
    }

    args := os.Args[1:]
    guess, err := strconv.Atoi(args[0])
    if err != nil {
        fmt.Println("Not a number.")
        return
    } else if guess < 0 {
        fmt.Println("Please enter a number greater than 0")
        return
    }

    for turn := 1; turn <= maxTurns; turn++ {
        n := rand.Intn(guess + 1)

        if n == guess {
            switch rand.Intn(3) {
            case 0:
                fmt.Println("You're awesome")
                return
            case 1:
                fmt.Println("You won!")
                return
            case 2: 
                fmt.Println("Perfect!")
                return
            }


        }
    } 

    switch rand.Intn(3) {
    case 0:
    fmt.Println("You loser!")
    case 1:
    fmt.Println("Suck it")
    case 2:
    fmt.Println("Try harder, fool")
    }
}
