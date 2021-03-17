package lazadago

import (
	"fmt"
	"testing"
)

func TestAuthURL(t *testing.T) {
	client := NewClient(&ClientOptions{APIKey: "120429"})
	fmt.Println(client.GetAuthURL("120429", "https://erp.ds1689.cn/manystore/shop/shop/lazadabindget", "cn"))
}
