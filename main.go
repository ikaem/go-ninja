package main

import (
	"fmt"
	"sort"
)

func main() {

	// greeting := "Gello there freinds"

	// fmt.Println(strings.Contains(greeting, "Gello")) // returns boolean // false
	// fmt.Println((strings.ReplaceAll(greeting, "Gello", "hello")))
	// fmt.Println(strings.Index(greeting, "Ge"))
	// fmt.Println(strings.Split(greeting, " ")) // [Gello there freinds]

	// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	// sort.Ints(ages)

	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 99)

	// fmt.Println("index", index) // 8

	names := []string{"karlo", "ivan", "martina"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "ivan")) // 0

}
