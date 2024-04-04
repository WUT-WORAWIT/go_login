package main

import (
	Router "logins/Router"
)

func main() {
	r := Router.SetupRouter()
	//running
	r.Run()
}
