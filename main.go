package main

import "smart-recommendation/router"

func main() {
	route := router.NewRouter()
	route.Run("localhost:8080")
}
