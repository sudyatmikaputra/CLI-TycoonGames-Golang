/*	SID	: 1301180494
	Name: Ketut Sudyatmika Putra
	This program is a game called tycoon. It's simulate a trading system and the goals is to find a big profit.
*/

package main

import (
	"fmt"
	"math/rand"
)

const max = 50

type catArray struct {
	price     [max]int
	item      [max]int
	pricesell [max]int
	itemsell  [max]int
}

type catItem [5]catArray

var (
	item        = [5]string{"silk", "silver", "gold", "diamond", "storage"}
	itemprice   = [5]int{250, 625, 1250, 2500, 125}
	itemqty     = [5]int{0, 0, 0, 0, 0}
	itemqtyware = [4]int{0, 0, 0, 0}
	account     = 1300
	bought      = [5]int{0, 0, 0, 0, 0}
	sold        = [5]int{0, 0, 0, 0, 0}
	catalog     catItem
	itemqtytemp     = [5]int{0, 0, 0, 0, 0}
	capacity    int = 10
	totcap      int = 0
	totrentcap  int = 0
	rentcap     int = 0
	dayleft     int
	y           int = 0
)

/* 	I.S. quantity, good and name variable needed to run the function
F.S. buy function return changes on quantity, capacity, account, etc variable that needed for further function.
*/
func buy(quantity int, good string, name string) {
	var dec string
	if account > 0 {
		switch good {
		case "silk":
			fmt.Print("Buying ", quantity, " silk bought at ", itemprice[0], "? ")
			backmoney := account
			backqty := itemqty[0]
			fmt.Scanln(&dec)
			if dec == "yes" {
				account = account - itemprice[0]*quantity
				itemqty[0] = itemqty[0] + quantity
				totcap = totcap + quantity
				if totcap <= capacity {
					if account < 0 {
						account = backmoney
						itemqty[0] = backqty
						fmt.Print("I am sorry, ", name, ". You don't have enough money")
						fmt.Println()
					} else {
						fmt.Println("You bought", quantity, "silk for", itemprice[0]*quantity)
						catalog[0].item[bought[0]] = quantity
						catalog[0].price[bought[0]] = itemprice[0]
						bought[0]++
					}
				} else {
					fmt.Println("Your storage is full")
					account = backmoney
					itemqty[0] = backqty
					totcap = totcap - quantity
				}
			} else if dec == "no" {
				fmt.Println("Trading canceled")
			}

		case "silver":
			fmt.Print("Buying ", quantity, " silver bought at ", itemprice[1], "? ")
			backmoney := account
			backqty := itemqty[1]
			fmt.Scanln(&dec)
			if dec == "yes" {
				account = account - itemprice[1]*quantity
				itemqty[1] = itemqty[1] + quantity
				totcap = totcap + quantity
				if totcap <= capacity {
					if account < 0 {
						account = backmoney
						itemqty[1] = backqty
						fmt.Print("I am sorry, ", name, ". You don't have enough money")
						fmt.Println()
					} else {
						fmt.Println("You bought", quantity, "silk for", itemprice[1]*quantity)
						catalog[1].item[bought[1]] = quantity
						catalog[1].price[bought[1]] = itemprice[1]
						bought[1]++
					}
				} else {
					fmt.Println("Your storage is full")
					account = backmoney
					itemqty[1] = backqty
					totcap = totcap - quantity
				}
			} else if dec == "no" {
				fmt.Println("Trading canceled")
			}

		case "gold":
			fmt.Print("Buying ", quantity, " gold bought at ", itemprice[2], "? ")
			backmoney := account
			backqty := itemqty[2]
			fmt.Scanln(&dec)
			if dec == "yes" {
				account = account - itemprice[2]*quantity
				itemqty[2] = itemqty[2] + quantity
				totcap = totcap + quantity
				if totcap <= capacity {
					if account < 0 {
						account = backmoney
						itemqty[2] = backqty
						fmt.Print("I am sorry, ", name, ". You don't have enough money")
						fmt.Println()
					} else {
						fmt.Println("You bought", quantity, "silk for", itemprice[2]*quantity)
						catalog[2].item[bought[2]] = quantity
						catalog[2].price[bought[2]] = itemprice[2]
						bought[2]++
					}
				} else {
					fmt.Println("Your storage is full")
					account = backmoney
					itemqty[2] = backqty
					totcap = totcap - quantity
				}
			} else if dec == "no" {
				fmt.Println("Trading canceled")
			}

		case "diamond":
			fmt.Print("Buying ", quantity, " diamond bought at ", itemprice[3], "? ")
			backmoney := account
			backqty := itemqty[3]
			fmt.Scanln(&dec)
			if dec == "yes" {
				account = account - itemprice[3]*quantity
				itemqty[3] = itemqty[3] + quantity
				totcap = totcap + quantity
				if totcap <= capacity {
					if account < 0 {
						account = backmoney
						itemqty[3] = backqty
						fmt.Print("I am sorry, ", name, ". You don't have enough money")
						fmt.Println()
					} else {
						fmt.Println("You bought", quantity, "silk for", itemprice[3]*quantity)
						catalog[3].item[bought[3]] = quantity
						catalog[3].price[bought[3]] = itemprice[3]
						bought[3]++
					}
				} else {
					fmt.Println("Your storage is full")
					account = backmoney
					itemqty[3] = backqty
					totcap = totcap - quantity
				}
			} else if dec == "no" {
				fmt.Println("Trading canceled")
			}
		case "storage":
			fmt.Print("Expanding your storage for ", quantity, " space at ", itemprice[4], "? ")
			backmoney := account
			backqty := capacity
			fmt.Scanln(&dec)
			if dec == "yes" {
				account = account - itemprice[4]*quantity
				capacity = capacity + quantity
				if account < 0 {
					account = backmoney
					capacity = backqty
					fmt.Print("I am sorry, ", name, ". You don't have enough money")
					fmt.Println()
				} else {
					fmt.Println("You bought", quantity, "storage for", itemprice[4]*quantity)
					catalog[4].item[bought[4]] = quantity
					catalog[4].price[bought[4]] = itemprice[4]
					bought[4]++
				}
			} else if dec == "no" {
				fmt.Println("Trading canceled")
			}

		}
	} else {
		fmt.Println("You don't have money.")
	}
}

/* 	I.S. quantity, good and name variable needed to run the function
F.S. sell function return changes on quantity, capacity, account, etc variable that needed for further function.
*/
func sell(quantity int, good string, name string) {
	var dec string
	switch good {
	case "silk":
		fmt.Print("Selling ", quantity, " silk at ", itemprice[0], "? ")
		backqty := itemqty[0]
		backmoney := account
		fmt.Scanln(&dec)
		if dec == "yes" {
			if itemqty[0] > 0 {
				account = account + itemprice[0]*quantity
				itemqty[0] = itemqty[0] - quantity
				if itemqty[0] < 0 {
					itemqty[0] = backqty
					account = backmoney
					fmt.Print("I am sorry, ", name, ". You don't have enough item to sell.")
					fmt.Println()
				} else {
					fmt.Println("You sold", quantity, "silk for", itemprice[0]*quantity)
					catalog[0].itemsell[sold[0]] = quantity
					catalog[0].pricesell[sold[0]] = itemprice[0]
					sold[0]++
					totcap = totcap - quantity
				}
			} else {
				fmt.Print("Sorry, ", name, ". You don't have enough item to sell.")
				fmt.Println()
			}
		} else if dec == "no" {
			fmt.Println("Trading canceled")
		}

	case "silver":
		fmt.Print("Selling ", quantity, " silver at ", itemprice[1], "? ")
		backqty := itemqty[1]
		backmoney := account
		fmt.Scanln(&dec)
		if dec == "yes" {
			if itemqty[1] > 0 {
				account = account + itemprice[1]*quantity
				itemqty[1] = itemqty[1] - quantity
				if itemqty[1] < 0 {
					itemqty[1] = backqty
					account = backmoney
					fmt.Print("I am sorry, ", name, ". You don't have enough item to sell.")
					fmt.Println()
				} else {
					fmt.Println("You sold", quantity, "silver for", itemprice[1]*quantity)
					catalog[1].itemsell[sold[1]] = quantity
					catalog[1].pricesell[sold[1]] = itemprice[1]
					sold[1]++
					totcap = totcap - quantity
				}
			} else {
				fmt.Print("Sorry, ", name, ". You don't have enough item to sell.")
				fmt.Println()
			}
		} else if dec == "no" {
			fmt.Println("Trading canceled")
		}

	case "gold":
		fmt.Print("Selling ", quantity, " gold at ", itemprice[2], "? ")
		backqty := itemqty[2]
		backmoney := account
		fmt.Scanln(&dec)
		if dec == "yes" {
			if itemqty[2] > 0 {
				account = account + itemprice[2]*quantity
				itemqty[2] = itemqty[2] - quantity
				if itemqty[2] < 0 {
					itemqty[2] = backqty
					account = backmoney
					fmt.Print("I am sorry, ", name, ". You don't have enough item to sell.")
					fmt.Println()
				} else {
					fmt.Println("You sold", quantity, "gold for", itemprice[2]*quantity)
					catalog[2].itemsell[sold[2]] = quantity
					catalog[2].pricesell[sold[2]] = itemprice[2]
					sold[2]++
					totcap = totcap - quantity
				}
			} else {
				fmt.Print("Sorry, ", name, ". You don't have enough item to sell.")
				fmt.Println()
			}
		} else if dec == "no" {
			fmt.Println("Trading canceled")
		}

	case "diamond":
		fmt.Print("Selling ", quantity, " diamond at ", itemprice[3], "? ")
		backqty := itemqty[3]
		backmoney := account
		fmt.Scanln(&dec)
		if dec == "yes" {
			if itemqty[3] > 0 {
				account = account + itemprice[3]*quantity
				itemqty[3] = itemqty[3] - quantity
				if itemqty[3] < 0 {
					itemqty[3] = backqty
					account = backmoney
					fmt.Print("I am sorry, ", name, ". You don't have enough item to sell.")
					fmt.Println()
				} else {
					fmt.Println("You sold", quantity, "diamond for", itemprice[3]*quantity)
					catalog[3].itemsell[sold[3]] = quantity
					catalog[3].pricesell[sold[3]] = itemprice[3]
					sold[3]++
					totcap = totcap - quantity
				}
			} else {
				fmt.Print("Sorry, ", name, ". You don't have enough item to sell.")
				fmt.Println()
			}
		} else if dec == "no" {
			fmt.Println("Trading canceled")
		}

	case "storage":
		fmt.Print("Selling ", quantity, " storage at ", itemprice[4], "? ")
		backqty := capacity
		backmoney := account
		fmt.Scanln(&dec)
		if dec == "yes" {
			if capacity > 0 {
				account = account + itemprice[4]*quantity
				capacity = capacity - quantity
				if capacity < 0 {
					capacity = backqty
					account = backmoney
					fmt.Print("I am sorry, ", name, ". You don't have enough storage to sell.")
					fmt.Println()
				} else {
					fmt.Println("You sold", quantity, "storage for", itemprice[4]*quantity)
					catalog[4].itemsell[sold[4]] = quantity
					catalog[4].pricesell[sold[4]] = itemprice[4]
					sold[4]++
				}
			} else {
				fmt.Print("Sorry, ", name, ". You don't have enough storage to sell.")
				fmt.Println()
			}
		} else if dec == "no" {
			fmt.Println("Trading canceled")
		}
	}
}

/* 	I.S. random and random2 variable declared
F.S. return up, down, neutral depends on the cases.
*/
func random() string {
	random := rand.Intn(100)
	random2 := rand.Intn(100)

	if random2 <= 50 {
		if random <= 70 {
			return "up"
		} else {
			return "down"
		}
	} else {
		return "neutral"
	}
}

/* 	I.S. take case from function "random" : up, down, neutral
F.S. up, down, and neutral case filled.
*/
func priceChange(changePrice string) {
	switch changePrice {
	case "up":
		for i := 0; i < 4; i++ {
			itemprice[i] = itemprice[i] + (itemprice[i] * 25 / 100)
		}
		fmt.Println("*Price info : Price up today")
		fmt.Println()
	case "down":
		for i := 0; i < 4; i++ {
			itemprice[i] = itemprice[i] - (itemprice[i] * 15 / 100)
		}
		fmt.Println("*Price info : Price down today")
		fmt.Println()
	case "neutral":
		fmt.Println("*Price info : Price doesn't change today.")
		fmt.Println()
	}
}

/* 	I.S. Catalog empty
F.S. Catalog filled
*/
func showCatalog() {
	fmt.Println("Current commodity prices:")
	fmt.Println("silk", itemprice[0])
	fmt.Println("silver", itemprice[1])
	fmt.Println("gold", itemprice[2])
	fmt.Println("diamond", itemprice[3])
	fmt.Println("storage", itemprice[4])
}

/* 	I.S. Account states empty
F.S. Account states filled
*/
func showAccount() {
	i := 0
	a := 0
	b := 0
	c := 0
	d := 0
	e := 0
	i2 := 0
	a2 := 0
	b2 := 0
	c2 := 0
	d2 := 0

	fmt.Println("Your balance is", account)
	fmt.Println("You have in your storage: ")
	fmt.Println()

	if itemqty[0] > 0 {
		for i < bought[0] {
			fmt.Println(catalog[0].item[i], "silk, bought at", catalog[0].price[i])
			i++
		}
		for i2 < sold[0] {
			fmt.Println(catalog[0].itemsell[i2], "silk, sold at", catalog[0].pricesell[i2])
			i2++
		}
	} else {
		fmt.Println("0 silk")
	}
	fmt.Println("Total silk:", itemqty[0])
	fmt.Println()

	if itemqty[1] > 0 {
		for a < bought[1] {
			fmt.Println(catalog[1].item[a], "silver, bought at", catalog[1].price[a])
			a++
		}
		for a2 < sold[1] {
			fmt.Println(catalog[1].itemsell[a2], "silver, sold at", catalog[1].pricesell[a2])
			a2++
		}
	} else {
		fmt.Println("0 silver")
	}
	fmt.Println("Total silver:", itemqty[1])
	fmt.Println()

	if itemqty[2] > 0 {
		for b < bought[2] {
			fmt.Println(catalog[2].item[b], "gold, bought at", catalog[2].price[b])
			b++
		}
		for b2 < sold[2] {
			fmt.Println(catalog[2].itemsell[b2], "gold, sold at", catalog[2].pricesell[b2])
			b2++
		}
	} else {
		fmt.Println("0 gold")
	}
	fmt.Println("Total gold:", itemqty[2])
	fmt.Println()

	if itemqty[3] > 0 {
		for c < bought[3] {
			fmt.Println(catalog[3].item[c], "diamond, bought at", catalog[3].price[c])
			c++
		}
		for c2 < sold[3] {
			fmt.Println(catalog[3].itemsell[c2], "diamond, sold at", catalog[3].pricesell[c2])
			c2++
		}
	} else {
		fmt.Println("0 diamond")
	}
	fmt.Println("Total diamond:", itemqty[3])
	fmt.Println()

	if capacity > 0 {
		for d < bought[4] {
			fmt.Println(catalog[4].item[d], "storage, bought at", catalog[4].price[d])
			d++
		}
		for d2 < sold[4] {
			fmt.Println(catalog[4].itemsell[d2], "storage, sold at", catalog[4].pricesell[d2])
			d2++
		}
	}
	itemqty[4] = capacity
	fmt.Println("Used capacity:", itemqty[0]+itemqty[1]+itemqty[2]+itemqty[3])
	fmt.Println("Total capacity:", capacity)
	fmt.Println()
	fmt.Println("Used warehouse capacity:", totrentcap)
	fmt.Println("Total warehouse capacity:", rentcap)
	fmt.Println("Warehouse rent dayleft:", dayleft)
	fmt.Println()
	fmt.Println("Warehouse info:")
	for i = 0; i < 4; i++ {
		fmt.Println(itemqtyware[i], item[i])
	}

	if itemqty[0] > itemqtytemp[0] || itemqty[1] > itemqtytemp[1] || itemqty[2] > itemqtytemp[2] || itemqty[3] > itemqtytemp[3] || itemqty[4] > itemqtytemp[4] {
		fmt.Println("Thank you for your business. New price list has been issued.")
	}

	e = 0
	for e < 5 {
		itemqtytemp[e] = itemqty[e]
		e++
	}

}

/* 	I.S. take quantity, name variable
F.S. run function and knowing it's rent or not
*/
func rent(quantity int, name string) {
	account = account - quantity
	if account < 0 {
		account = account + quantity
		fmt.Print("Sorry, ", name, ". You don't have enough money.")
		fmt.Println()
	} else {
		fmt.Print("Thank you, ", name, ". You rent a warehouse for ", quantity, " day.")
		fmt.Println()
	}

}

/* 	I.S. take quntity and good variable
F.S. "store" function return changes on quantity, capacity, account, etc variable that needed for further function.
*/
func store(quantity int, good string) {
	totrentcap = totrentcap + quantity
	if totrentcap <= rentcap {
		switch good {
		case "silk":
			itemqtyware[0] = itemqtyware[0] + quantity
			itemqty[0] = itemqty[0] - quantity
			totcap = totcap - quantity
			if itemqty[0] >= 0 {
				fmt.Println("You silk has been stored.")
			} else {
				itemqtyware[0] = itemqtyware[0] - quantity
				itemqty[0] = itemqty[0] + quantity
				totcap = totcap + quantity
				fmt.Println("You don't have enough silk.")
			}
		case "silver":
			itemqtyware[1] = itemqtyware[1] + quantity
			itemqty[1] = itemqty[1] - quantity
			totcap = totcap - quantity
			if itemqty[0] >= 0 {
				fmt.Println("You silver has been stored.")
			} else {
				itemqtyware[1] = itemqtyware[1] - quantity
				itemqty[1] = itemqty[1] + quantity
				totcap = totcap + quantity
				fmt.Println("You don't have enough silver.")
			}

		case "gold":
			itemqtyware[2] = itemqtyware[2] + quantity
			itemqty[2] = itemqty[2] - quantity
			totcap = totcap - quantity
			if itemqty[0] >= 0 {
				fmt.Println("You gold has been stored.")
			} else {
				itemqtyware[2] = itemqtyware[2] - quantity
				itemqty[2] = itemqty[2] + quantity
				totcap = totcap + quantity
				fmt.Println("You don't have enough gold.")
			}

		case "diamond":
			itemqtyware[3] = itemqtyware[3] + quantity
			itemqty[3] = itemqty[3] - quantity
			totcap = totcap - quantity
			if itemqty[0] >= 0 {
				fmt.Println("You diamond has been stored.")
			} else {
				itemqtyware[3] = itemqtyware[3] - quantity
				itemqty[3] = itemqty[3] + quantity
				totcap = totcap + quantity
				fmt.Println("You don't have enough diamond.")
			}
		}
	} else {
		totrentcap = totrentcap - quantity
		fmt.Println("Your warehouse is full.")
	}
}

/* 	I.S. take quntity and good variable
F.S. "unload" function return changes on quantity, capacity, account, etc variable that needed for further function.
*/
func unload(quantity int, good string) {
	totrentcap = totrentcap - quantity
	if totrentcap >= 0 {
		switch good {
		case "silk":
			itemqtyware[0] = itemqtyware[0] - quantity
			itemqty[0] = itemqty[0] + quantity
			totcap = totcap + quantity
			if itemqty[0] >= 0 {
				fmt.Println("You silk has been unloaded.")
			} else {
				itemqtyware[0] = itemqtyware[0] + quantity
				itemqty[0] = itemqty[0] - quantity
				totcap = totcap - quantity
				fmt.Println("You don't have that much silk.")
			}
		case "silver":
			itemqtyware[1] = itemqtyware[1] - quantity
			itemqty[1] = itemqty[1] + quantity
			totcap = totcap + quantity
			if itemqty[0] >= 0 {
				fmt.Println("You silver has been unloaded.")
			} else {
				itemqtyware[1] = itemqtyware[1] + quantity
				itemqty[1] = itemqty[1] - quantity
				totcap = totcap - quantity
				fmt.Println("You don't have that much silver.")
			}

		case "gold":
			itemqtyware[2] = itemqtyware[2] - quantity
			itemqty[2] = itemqty[2] + quantity
			totcap = totcap + quantity
			if itemqty[0] >= 0 {
				fmt.Println("You gold has been unloaded.")
			} else {
				itemqtyware[2] = itemqtyware[2] + quantity
				itemqty[2] = itemqty[2] - quantity
				totcap = totcap - quantity
				fmt.Println("You don't have that much gold.")
			}

		case "diamond":
			itemqtyware[3] = itemqtyware[3] - quantity
			itemqty[3] = itemqty[3] + quantity
			totcap = totcap + quantity
			if itemqty[0] >= 0 {
				fmt.Println("You diamond has been unloaded.")
			} else {
				itemqtyware[3] = itemqtyware[3] + quantity
				itemqty[3] = itemqty[3] - quantity
				totcap = totcap - quantity
				fmt.Println("You don't have that much diamond.")
			}
		}
	} else {
		totrentcap = totrentcap + quantity
		fmt.Println("You don't have that much item.")
	}

}

/* 	I.S. data unsorted
F.S. data sorted
*/
func iSort(accsort *[max]float32, accday *[max]float32) {
	for i := 1; i < y; i++ {
		j := i
		for j > 0 && accsort[j] < accsort[j-1] {
			accsort[j], accsort[j-1] = accsort[j-1], accsort[j]
			accday[j], accday[j-1] = accday[j-1], accday[j]
			j -= 1
		}
	}
}

/* 	I.S. data still not printed yet
F.S. data printed
*/
func printTable(accsort [max]float32, accday [max]float32) {
	i := 0
	for i < y {
		fmt.Println(float64(accsort[i]), "on day", accday[i])
		i++
	}
}

/* 	I.S. data still not searched
F.S. data searched
*/
func bSearch(accsort [max]float32, X float32, accday [max]float32) {
	var m int
	right := y
	left := 0
	for left < right && accsort[m] != X {
		m = (right + left) / 2
		if accsort[m] == X {
			fmt.Println("Search", X, "found on day", accday[m])
		} else if accsort[m] > X {
			right = m - 1
		} else {
			left = m + 1
		}
	}
	if accsort[m] != X {
		fmt.Println("Search not found. Last value is", accsort[m])
	}
}

func main() {
	var name, function, good, categories string
	var quantity, initaccount int
	var cost, totcost, maxa, mina, accreal, search float32
	var accsort [max]float32
	var accday [max]float32

	x := rand.Intn(5)
	day := 1
	initaccount = account
	accreal = float32(account)
	maxa = accreal
	mina = accreal
	totcost = 0

	fmt.Print("Welcome sir/mam, your name is? ")
	fmt.Scanln(&name)
	fmt.Print(name, ", as a startup, you've been given account 1300. Trade wisely.")
	fmt.Println()

	for function != "quit" {
		x = rand.Intn(5)
		for x == 0 {
			x = rand.Intn(5)
		}

		fmt.Print("[Day ", day, "]What's your command? ")
		fmt.Scan(&function)

		for function != "show" && function != "buy" && function != "sell" && function != "rent" && function != "store" && function != "unload" && function != "quit" {
			fmt.Print("Unknown command, please input a new command: ")
			fmt.Scan(&function)
		}

		if function != "show" {
			priceChange(random())
			itemprice[4] = itemprice[4] + 1
			if itemprice[4] == itemprice[0] || itemprice[4] >= itemprice[0] {
				itemprice[4] = itemprice[0]
			}
			dayleft = dayleft - x
			if dayleft < 0 {
				rentcap = 0
				dayleft = 0
			}
			if dayleft <= 0 && totrentcap > 0 {
				account = account - x
				fmt.Println("**You exceed your rent time and your balance decreased by 1 unit each day.**")
				fmt.Println("**Please unload your items in the warehouse**")
				fmt.Println()
			}
		}
		switch function {
		case "buy":
			fmt.Scanln(&quantity, &good)
			buy(quantity, good, name)
			fmt.Println()
		case "sell":
			fmt.Scanln(&quantity, &good)
			sell(quantity, good, name)
			fmt.Println()
		case "show":
			fmt.Scan(&categories)
			if categories == "account" {
				showAccount()
				fmt.Println()
			} else if categories == "catalog" {
				showCatalog()
				fmt.Println()
			}
		case "rent":
			fmt.Scan(&quantity)
			dayleft = quantity
			rentcap = quantity
			rent(quantity, name)
			fmt.Println()
		case "store":
			fmt.Scanln(&quantity, &good)
			store(quantity, good)
			fmt.Println()
		case "unload":
			fmt.Scanln(&quantity, &good)
			unload(quantity, good)
			fmt.Println()
		}
		if function != "show" {
			for i := 0; i < x; i++ {
				totcost = totcost + cost
				cost = 0.01 * float32(capacity)
			}
			accday[y] = float32(day)
			day = day + x
			accreal = float32(account)
			accsort[y] = accreal - totcost
			y++
		}
		if accreal-totcost > maxa {
			maxa = accreal - totcost
		} else if accreal-totcost < mina {
			mina = accreal - totcost
		}
	}

	fmt.Print(name, ", we liquidate all your asset")
	fmt.Println()
	fmt.Println("Your maintenance cost throughout the day is,", totcost)
	fmt.Println("Your final asset value (final balance - maintenance cost) is", accreal-totcost)
	fmt.Println("Your maximum account value is", maxa)
	fmt.Println("Your minimum account value is", mina)
	fmt.Println("Initial account:", initaccount)
	fmt.Println("Final account:", accreal-totcost)
	fmt.Println("The difference between your initial to final account is", (accreal-totcost)-float32(initaccount))
	fmt.Println()
	if account > initaccount {
		fmt.Println("Apparently you WIN")
		fmt.Println()
	} else {
		fmt.Print("Sorry ", name, ", you lose. Goodluck another time!")
		fmt.Println()
		fmt.Println()
	}
	iSort(&accsort, &accday)
	fmt.Print("Sorting...")
	fmt.Println()
	printTable(accsort, accday)
	fmt.Println()
	fmt.Print("Input number to search: ")
	fmt.Scan(&search)
	bSearch(accsort, search, accday)
}

//bSearch sometimes not found because it's still contain invisible decimal numbers
