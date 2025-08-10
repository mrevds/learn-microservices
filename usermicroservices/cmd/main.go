package main

import (
	"microservices-learn/usermicroservices/database"
	"microservices-learn/usermicroservices/router"
)

func main() {
	router.InitRouter()
	database.InitDB()
}
