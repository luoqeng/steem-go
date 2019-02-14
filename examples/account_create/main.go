package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"

	"github.com/luoqeng/steem-go"
	"github.com/luoqeng/steem-go/types"
)

const key = "5JNHfZYKGaomSFvd4NUdQ9qMcEAC43kujbfjueTHpVapX1Kzq2n"

func main() {
	cls, err := client.NewClient([]string{"ws://127.0.0.1:8090"})
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer cls.Close()
	if err := run(cls); err != nil {
		log.Fatalln("Error:", err)
	}
}

func run(cls *client.Client) (err error) {

	flag.Parse()
	// Process args.
	args := flag.Args()

	if len(args) != 1 {
		return errors.New("1 arguments required")
	}
	newAccountName := args[0]

	cls.SetKeys(&client.Keys{AKey: []string{key}})

	password := client.GenPassword()
	log.Printf("---> account create name = %s password = %s\n", newAccountName, password)

	resp, err := cls.AccountCreate("initminer", newAccountName, password, &types.Asset{Amount: 0, Symbol: "TESTS"})
	if err != nil {
		return err
	}

	json_fmt, _ := json.Marshal(resp)
	log.Printf("resp %s\n", json_fmt)

	return nil
}
