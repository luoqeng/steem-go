package main

import (
	"encoding/json"
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

	cls.SetKeys(&client.Keys{AKey: []string{key}})

	var arrtrans = []client.ArrTransfer{
		client.ArrTransfer{
			To:     "scott",
			Memo:   "10000",
			Amount: types.Asset{Amount: 1, Symbol: "TESTS"},
		},
		client.ArrTransfer{
			To:     "luoqeng",
			Memo:   "10000",
			Amount: types.Asset{Amount: 1, Symbol: "TESTS"},
		},
	}

	resp, err := cls.MultiTransfer("initminer", arrtrans)
	if err != nil {
		return err
	}

	json_fmt, _ := json.Marshal(resp)
	log.Printf("resp %s\n", json_fmt)

	return nil
}
