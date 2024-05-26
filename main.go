package main

import (
	"pari/passkey-v2/database"
)

func main() {
	data := database.DB{}
	path := database.GetDotfilePath()
	// err := data.Set("new", "new2")
	err := data.Load(path)
	if err != nil {
		panic(err)
	}
	// err = data.Set("tes2", "tes1")
	// if err != nil {
	// 	panic(err)
	// }
	// err, res := data.Get("new")
	// if err != nil {
	// 	panic(err)
	// }
	data.List()
	err = data.Store(path)
	if err != nil {
		panic(err)
	}
}
