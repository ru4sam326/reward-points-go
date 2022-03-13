package main

import (
	"fmt"
	"log"
	"os"

	"github.com/reward-points-go/purchase"
)

func main() {
	args := os.Args[1:]
	filePath := "purchase_details.json"

	if len(args) == 0 {
		log.Println("Purchase details file wss not provided, Going with the sample data")
	} else {
		filePath = args[0]
	}
	purchaseMap := purchase.GetCustomerWisePurchaseData(filePath)
	for name, history := range purchaseMap {
		fmt.Println("Reward point Summary Of Customer:" + name)
		monthWiseRewards := purchase.GetMonthWiseRewards(history)
		total := 0
		for key, rewards := range monthWiseRewards {
			fmt.Printf("%s : %d", key, rewards)
			fmt.Println()
			total += rewards
		}

		fmt.Printf("Total Reward points: %d", total)
		fmt.Println()
	}
}
