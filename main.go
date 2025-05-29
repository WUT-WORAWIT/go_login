package main

import (
	R "go_login/Router"
)

func main() {
	r := R.SetupRouter()
	//running
	r.Run()
}
