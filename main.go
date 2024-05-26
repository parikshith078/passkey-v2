package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"pari/passkey-v2/database"
)

func main() {

	var (
		set  = flag.Bool("set", false, "Add new key value pair")
		get  = flag.String("get", "", "Get value using key")
		del  = flag.String("del", "", "Delete a pair using key")
		list = flag.Bool("list", false, "List all key value pairs")
	)

	flag.Parse()

	data := database.DB{}
	filePath := database.GetDotfilePath()

	if err := data.Load(filePath); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *set:
		args := flag.Args()
		if len(args) != 2 {
			fmt.Fprintln(os.Stderr, errors.New("2 arguments only").Error())
			os.Exit(1)
		}
		if err := data.Set(args[0], args[1]); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		if err := data.Store(filePath); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *get != "":
		err, res := data.Get(*get)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Println(res.Value)

	case *del != "":
		if err := data.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		if err := data.Store(filePath); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *list:
		data.List()

	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
