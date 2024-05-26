package database

import (
	"bufio"
	"errors"
	"io"
	"os/user"
	"strings"
)

func GetDotfilePath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	path := usr.HomeDir + "/go/db/passkey.json"
	return path
}

func getInput(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return text, nil

}
