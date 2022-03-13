package util_test

import (
	"testing"

	"github.com/reward-points-go/util"
)

func TestGetFileContentPass(t *testing.T) {
	content, err := util.GetFileContent("purchase_details.json")
	if len(content) == 0 || err != nil {
		t.Errorf("File Content should be read, But unable to read")
	}
}

func TestGetFileContentFail(t *testing.T) {
	content, _ := util.GetFileContent("")
	if len(content) != 0 {
		t.Errorf("File Content should be zero")
	}
}

func TestCalcRewardPoints(t *testing.T) {
	rewards := util.CalcRewardPoints(5.0)
	if rewards != 0 {
		t.Errorf("Reward points should be zero")
	}

	rewards = util.CalcRewardPoints(50.0)
	if rewards != 0 {
		t.Errorf("Reward points should be zero")
	}

	rewards = util.CalcRewardPoints(51.0)
	if rewards != 1 {
		t.Errorf("Reward points should be 1")
	}

	rewards = util.CalcRewardPoints(59.0)
	if rewards != 9 {
		t.Errorf("Reward points should be 9")
	}

	rewards = util.CalcRewardPoints(120.0)
	if rewards != 90 {
		t.Errorf("Reward points should be 90")
	}
}
