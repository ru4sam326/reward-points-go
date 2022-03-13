package main

import (
	"log"
	"os"

	"github.com/reward-points-go/purchase"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Println("Purchase details file is required")
		return
	}
	purchaseMap := purchase.GetCustomerWisePurchaseData(args[0])
	log.Println(purchaseMap)
}
