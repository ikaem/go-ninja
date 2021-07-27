// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')


	return strings.TrimSpace(input), err

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name - ", reader)
	b := newBill(name)

	fmt.Println("The bill created - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option - (a: add item, s: save bill, t - add tip)", reader)

	// fmt.Println("You chose:", opt)

	switch opt {
	case "a", "f":
		fmt.Println("You chose 'a'")
		name, _ := getInput("What is the item name?", reader)
		price, _ := getInput("What is the item price?", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)
		fmt.Printf("You have added an item %v, with price of %v", name, price)
		promptOptions(b)
	case "t":
		fmt.Println("You chose 't'")
		tipValue, _ := getInput("What is the tip amount? $", reader)

		t, err := strconv.ParseFloat(tipValue, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updateTip(t)

		fmt.Printf("You have added a tip of $%T", t)
		promptOptions(b)
	case "s":

		fmt.Println("You chose 's'")
		b.save()
	default:
		fmt.Println("Not a valid option")
	}

}

func main() {

	myBill := createBill()

	promptOptions(myBill)

	fmt.Println(myBill)

}
