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

func CalcRewardPoints(amount float64) int {
	if amount <= 50.0 {
		return 0
	}
	if amount > 50.0 && amount <= 100 {
		return int(amount - 50.0)
	}

	return int((amount-100.0))*2 + 50
}
