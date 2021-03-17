package lazadago

import (
	"fmt"
	"testing"
)

func TestAuthURL(t *testing.T) {
	fmt.Println(GetAuthURL("120429", "https://erp.ds1689.cn/manystore/shop/shop/lazadabindget", "cn"))
}
