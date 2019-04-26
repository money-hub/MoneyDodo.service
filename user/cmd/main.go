package main

import (
	_ "github.com/money-hub/MoneyDodo.service/swagger"
	service "github.com/money-hub/MoneyDodo.service/user/cmd/service"
)

func main() {
	service.Run()
}
