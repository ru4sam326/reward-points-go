package purchase_test

import (
	"testing"

	"github.com/reward-points-go/purchase"
)

func TestGetCustomerWisePurchaseDataPass(t *testing.T) {
	purchaseMap := purchase.GetCustomerWisePurchaseData("purchase_details.json")
	if len(purchaseMap) == 0 {
		t.Errorf("purchaseMap map should be greater than zero, But it is zero")
	}
}

func TestGetCustomerWisePurchaseDataFail(t *testing.T) {
	purchaseMap := purchase.GetCustomerWisePurchaseData("")
	if len(purchaseMap) != 0 {
		t.Errorf("purchaseMap map should be zero")
	}
}

func TestGetMonthWiseRewardsPass(t *testing.T) {
	purchaseMap := purchase.GetCustomerWisePurchaseData("purchase_details.json")
	monthWiseRewards := purchase.GetMonthWiseRewards(purchaseMap["John"])
	if len(monthWiseRewards) == 0 {
		t.Errorf("monthWiseRewards map should be greater than zero, But it is zero")
	}
}

func TestGetMonthWiseRewardsFail(t *testing.T) {
	purchaseMap := purchase.GetCustomerWisePurchaseData("purchase_details.json")
	monthWiseRewards := purchase.GetMonthWiseRewards(purchaseMap["Johny"])
	if len(monthWiseRewards) != 0 {
		t.Errorf("monthWiseRewards map should be  zero")
	}
}
