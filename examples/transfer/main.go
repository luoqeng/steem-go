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
	key    = "5JNHfZYKGaomSFvd4NUdQ9qMcEAC43kujbfjueTHpVapX1Kzq2n"
	toName string
	amount float64
	memo   string
)

func init() {
	flag.StringVar(&toName, "toname", "", "receive account name")
	flag.Float64Var(&amount, "amount", 1.0, "send amount")
	flag.StringVar(&memo, "memo", "", "remarks")
}

func main() {

	flag.Parse()
	if toName == "" {
		fmt.Println("toname required")
		return
	}

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

	cls.SetKeys(&client.Keys{AKey: []string{key}})

	resp, err := cls.Transfer("initminer", toName, memo, &types.Asset{Amount: amount, Symbol: "STEEM"})
	if err != nil {
		return err
	}

	json_fmt, _ := json.Marshal(resp)
	log.Printf("resp %s\n", json_fmt)

	return nil
}
