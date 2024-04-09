package main

import (
	R "logins/Router"
)

func main() {
	r := R.SetupRouter()
	//running
	r.Run()
}
