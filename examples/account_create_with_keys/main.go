package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/luoqeng/steem-go"
	"github.com/luoqeng/steem-go/types"
)

var (
	url                                   string
	creator, akey                         string
	newName, owner, active, posting, memo string
	fee                                   float64
)

func init() {
	flag.StringVar(&creator, "creator", "", "")
	flag.StringVar(&akey, "akey", "", "active private key")
	flag.StringVar(&newName, "newname", "", "new account name")
	flag.StringVar(&owner, "owner", "", "owner public key")
	flag.StringVar(&active, "active", "", "active public key")
	flag.StringVar(&posting, "posting", "", "posting public key")
	flag.StringVar(&memo, "memo", "", "memo public key")
	flag.StringVar(&url, "url", "ws://127.0.0.1:8090", "ws url")
	flag.Float64Var(&fee, "fee", 0.0, "")
}

func main() {
	flag.Parse()
	if creator == "" || akey == "" || newName == "" || owner == "" || active == "" || posting == "" || memo == "" {
		fmt.Println("Parameter cannot be empty")
		return
	}

	cls, err := client.NewClient([]string{url})
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer cls.Close()
	if err := run(cls); err != nil {
		log.Fatalln("Error:", err)
	}
}

func run(cls *client.Client) (err error) {
	cls.SetKeys(&client.Keys{AKey: []string{akey}})
	resp, err := cls.AccountCreateWithKeys(creator, newName, owner, active, posting, memo, &types.Asset{Amount: fee, Symbol: "STEEM"})
	if err != nil {
		return err
	}

	json_fmt, _ := json.Marshal(resp)
	log.Printf("resp %s\n", json_fmt)

	return nil
}
