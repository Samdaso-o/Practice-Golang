package main

import (
	"fmt"

	"./banking"
)

//대문자는 퍼블릭, 소문자는 프라이빗으로 상속이 안됩니다^^

func main() {
	account := banking.Account{Customer: "han", Balance: 10000}
	fmt.Println(account.Customer)
}
