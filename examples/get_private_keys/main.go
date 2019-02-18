package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/luoqeng/steem-go"
)

var (
	newName, password string
)

func init() {
	flag.StringVar(&newName, "newname", "", "new account name")
	flag.StringVar(&password, "password", "", "")
}

func main() {
	flag.Parse()
	if newName == "" {
		fmt.Println("newName cannot be empty")
		return
	}
	if password == "" {
		password := client.GenPassword()
		log.Printf("---> account create name: %s ,password: %s\n", newName, password)
	}

	if err := run(); err != nil {
		log.Fatalln("Error:", err)
	}
}

func run() (err error) {
	type Keys struct {
		Private string
		Public  string
		Role    string
	}

	var listKeys = make(map[string]Keys)
	roles := [4]string{"owner", "active", "posting", "memo"}

	for _, val := range roles {
		priv := client.GetPrivateKey(newName, val, password)
		pub := client.GetPublicKey("STM", priv)
		listKeys[val] = Keys{Private: priv, Public: pub, Role: val}
	}

	for _, val := range listKeys {
		log.Printf("---> role: %s ,priv: %s ,pub: %s\n", val.Role, val.Private, val.Public)
	}

	return nil
}
