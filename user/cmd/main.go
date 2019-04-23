package main

import (
	service "github.com/money-hub/MoneyDodo.service/user/cmd/service"
	_ "github.com/money-hub/MoneyDodo.service/user/cmd/swagger"
)

func main() {
	service.Run()
}
