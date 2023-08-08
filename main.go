package main

import "smart-recommendation/router"

func main() {
	route := router.NewRouter()
	route.Run(":8080")
}
