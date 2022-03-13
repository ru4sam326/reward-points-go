package util_test

import (
	"testing"

	"github.com/reward-points-go/util"
)

func TestGetFileContentPass(t *testing.T) {
	content, err := util.GetFileContent("reviews.json")
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
