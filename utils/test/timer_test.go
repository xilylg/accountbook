package test

import (
	"fmt"
	"testing"

	"github.com/xilylg/accountbook/utils"
)

func TestStr2Unix(t *testing.T) {
	rst := utils.Day2Unix("2020-08-01")
	fmt.Print(rst)
}
