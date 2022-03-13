package main

import "makkonen.com/go-mongo/routes"

func main() {

	//there should be another way to have my routes in as separate file
	// and import it into main
	r := routes.InitializeRoutes()
	//router.Run("localhost:4000")
	r.Run("localhost:4000")
}
