package util

import (
	"io/ioutil"
	"log"
)

func GetFileContent(path string) ([]byte, error) {
	if path == "" {
		log.Println("No file path provided")
		return make([]byte, 0), nil
	}
	return ioutil.ReadFile(path)
}

// func CalcRewardPoints(data PurchaseData) float64 {
// 	return 0.0
// }
