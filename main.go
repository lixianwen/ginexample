package main

import "demo/router"

func main() {
        // setup router
        r := router.SetupRouter()

	r.Run(":2815")
}
