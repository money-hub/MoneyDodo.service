package main

import (
	service "github.com/money-hub/MoneyDodo.service/authentication/cmd/service"
	_ "github.com/money-hub/MoneyDodo.service/swagger"
)

func main() {
	service.Run()
}
