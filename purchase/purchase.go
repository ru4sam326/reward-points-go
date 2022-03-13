package purchase

import (
	"encoding/json"
	"log"
	"time"

	"github.com/reward-points-go/util"
)

type PurchaseData struct {
	CustomerName      string    `json:"customer_name"`
	TransactionDate   time.Time `json:"transaction_date"`
	TransactionAmount float64   `json:"transaction_amount"`
}

func GetCustomerWisePurchaseData(path string) map[string][]PurchaseData {
	customerMap := make(map[string][]PurchaseData)
	if path == "" {
		log.Println("File path required for reading the file " + path)
		return customerMap
	}
	content, err := util.GetFileContent(path)
	if err != nil {
		log.Println("Error occurred while reading the file " + path)
		return customerMap
	}
	data := make([]PurchaseData, 0)
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Println("Error occurred while reading the file content " + path + ", Not in right format")
		return customerMap
	}
	for _, record := range data {
		log.Println(record, "record")
		purchaseData := customerMap[record.CustomerName]
		if len(purchaseData) == 0 {
			purchaseData = make([]PurchaseData, 0)
		}
		purchaseData = append(purchaseData, record)
		customerMap[record.CustomerName] = purchaseData
	}
	return customerMap
}
