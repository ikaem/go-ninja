// bill.go
package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{
			// "pie": 3.99, "tea": 11.89,
		},
		tip: 0,
	}

	return b
}

// format the bill
func (b bill) format() string {

	fs := "Bill breakdown: \n"
	var total float64 = 0

	// LIST OUT THE ITEMS IN THE ITEMS
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v)
		total += v
	}

	// ouztput tip
	fs += fmt.Sprintf("%-25v ... $%0.2f\n", "tip:", b.tip)
	total += b.tip

	// ouztput total
	fs += fmt.Sprintf("%-25v ... $%0.2f\n", "total:", total)

	return fs

}

func (b *bill) updateTip(tip float64) {
	// would this actually update the tip?
	// there is a struct we are updatring
	// and its string part
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill

func (b *bill) save() {
	// save thse to a bills folder
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")

}
