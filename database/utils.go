package database

import "os/user"

func GetDotfilePath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	path := usr.HomeDir + "/go/db/passkey.json"
	return path
}

