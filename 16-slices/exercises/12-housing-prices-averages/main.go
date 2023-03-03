// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

import (
    "fmt"
    "strconv"
    "strings"
)
func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

    var (
        location []string
        size,beds,baths,price []int
    )

    rows := strings.Split(data, "\n")

    for _, row := range rows {
        cols := strings.Split(row, separator)

        location = append(location, cols[0])

        n, _ := strconv.Atoi(cols[1])
        size = append(size, n)

        n, _ = strconv.Atoi(cols[2])
        beds = append(beds, n)

        n, _ = strconv.Atoi(cols[3])
        baths = append(baths, n)

        n, _ = strconv.Atoi(cols[4])
        price = append(price, n)
    }

    for _, h := range strings.Split(header, separator){
        fmt.Printf("%-15s", h)
    }
    
    fmt.Printf("\n%s\n", strings.Repeat("=", 75))

    var (
        sizeSum, bedSum, bathSum, priceSum int
    )

    for index := range rows {
        fmt.Printf("%-15s", location[index])
        fmt.Printf("%-15d", size[index])
        sizeSum += size[index]
        fmt.Printf("%-15d", beds[index])
        bedSum += beds[index]
        fmt.Printf("%-15d", baths[index])
        bathSum += baths[index]
        fmt.Printf("%-15d", price[index])
        priceSum += price[index]
        fmt.Println()
    }

    fmt.Printf("\n%s\n", strings.Repeat("=", 75))

    fmt.Printf("%-15s", " ")

    fmt.Printf("%-15.2f", float64(sizeSum) / float64(len(size)))
    fmt.Printf("%-15.2f", float64(bedSum) / float64(len(beds)))
    fmt.Printf("%-15.2f", float64(bathSum) / float64(len(baths)))
    fmt.Printf("%-15.2f", float64(priceSum) / float64(len(price)))
}
