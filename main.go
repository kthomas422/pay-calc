// TODO: have settings for pay values
package main

import (
	"fmt"
	"os"
	"strconv"
)

func calculatePay(netPay float64) {
	var (
		houseFund, debt, savings, funTickets, bills float64
	)
	houseFund = netPay * .5
	debt = netPay * .2
	savings = netPay * .1
	funTickets = netPay * .05
	bills = netPay * .15

	fmt.Printf("House Fund : $%.2f\n", houseFund)
	fmt.Printf("Debt       : $%.2f\n", debt)
	fmt.Printf("Savings    : $%.2f\n", savings)
	fmt.Printf("Fun Tickets: $%.2f\n", funTickets)
	fmt.Printf("Bills      : $%.2f\n", bills)
	fmt.Printf("Bill amount: $%.2f\n\n", 190.0)
	fmt.Printf("After Bills: $%.2f\n", bills-190.0)
	fmt.Printf("Total      : $%.2f\n", houseFund+debt+savings+funTickets+bills) // will be off due to float rounding
}

func usage(name string) {
	fmt.Printf("Usage: %s <total pay>\n", name)
}

func main() {
	if len(os.Args) != 2 {
		usage(os.Args[0])
		os.Exit(0)
	}
	f, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		usage(os.Args[0])
		os.Exit(1)
	}
	calculatePay(f)
}
