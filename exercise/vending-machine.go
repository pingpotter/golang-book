package main

import "fmt"
import "github.com/BellBezel/vendingMachine"

var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
var items = map[string]int{"SD": 18, "CC": 12}

func main() {
	vm := vendingMachine.NewVendingMachine(coins, items)
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 18
	can := vm.SelectSD()
	fmt.Println(can) //SD

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 20
	can = vm.SelectCC()
	fmt.Println(can) //CC, F, TW ,O

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 25
	coin := vm.CoinReturn()
	fmt.Println(coin) //T, T, F
}
