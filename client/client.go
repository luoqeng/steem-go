package client

import (
	// Stdlib
	"encoding/json"
	"io/ioutil"
	"log"

	// Vendor
	"github.com/pkg/errors"

	// RPC
	"github.com/asuleymanov/golos-go"
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/transports/websocket"
	"github.com/asuleymanov/golos-go/types"
)

const fdt = `"20060102t150405"`

type User struct {
	Name  string `json:"username"`
	Chain string `json:"chain"`
	Url   string `json:"url"`
	PKey  string `json:"posting_key"`
	AKey  string `json:"active_key"`
	OKey  string `json:"owner_key"`
	MKey  string `json:"memo_key"`
}

type Client struct {
	Rpc   *rpc.Client
	User  *User
	Chain *transactions.Chain
}

type BResp struct {
	ID       string
	BlockNum uint32
	TrxNum   uint32
	Expired  bool
}

func readconfig() *User {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		log.Fatal(errors.Wrapf(e, "Error read config.json: "))
	}

	var jsontype *User
	if erru := json.Unmarshal(file, &jsontype); erru != nil {
		log.Println(errors.Wrapf(erru, "Error Unmarshal config.json: "))
		return nil
	}
	return jsontype
}

func initclient(url string) *rpc.Client {
	// Инициализация Websocket
	t, err := websocket.NewTransport(url)
	if err != nil {
		panic(errors.Wrapf(err, "Error Websocket: "))
	}

	// Инициализация RPC клиента
	client, err := rpc.NewClient(t)
	if err != nil {
		panic(errors.Wrapf(err, "Error RPC: "))
	}
	//defer client.Close()
	return client
}

func initChainId(str string) *transactions.Chain {
	var ChainId transactions.Chain
	// Определяем ChainId
	switch str {
	case "steem":
		ChainId = *transactions.SteemChain
	case "golos":
		ChainId = *transactions.GolosChain
	}
	return &ChainId
}

func NewApi() *Client {
	tmpUser := readconfig()
	return &Client{
		User:  tmpUser,
		Rpc:   initclient(tmpUser.Url),
		Chain: initChainId(tmpUser.Chain),
	}
}

func (api *Client) Send_Trx(strx types.Operation) (*BResp, error) {
	// Получение необходимых параметров
	props, err := api.Rpc.Database.GetDynamicGlobalProperties()
	if err != nil {
		return nil, errors.Wrapf(err, "Error get DynamicGlobalProperties: ")
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Добавление операций в транзакцию
	tx.PushOperation(strx)

	// Получаем необходимый для подписи ключ
	privKeys := api.Signing_Keys(strx)

	// Подписываем транзакцию
	if err := tx.Sign(privKeys, api.Chain); err != nil {
		return nil, errors.Wrapf(err, "Error Sign: ")
	}

	// Отправка транзакции
	resp, err := api.Rpc.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)

	if err != nil {
		return nil, errors.Wrapf(err, "Error BroadcastTransactionSynchronous: ")
	} else {
		var bresp BResp

		bresp.ID = resp.ID
		bresp.BlockNum = resp.BlockNum
		bresp.TrxNum = resp.TrxNum
		bresp.Expired = resp.Expired

		return &bresp, nil
	}
}

func (api *Client) Send_Arr_Trx(strx []types.Operation) (*BResp, error) {
	// Получение необходимых параметров
	props, err := api.Rpc.Database.GetDynamicGlobalProperties()
	if err != nil {
		return nil, errors.Wrapf(err, "Error get DynamicGlobalProperties: ")
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Добавление операций в транзакцию
	for _, val := range strx {
		tx.PushOperation(val)
	}

	// Получаем необходимый для подписи ключ
	privKeys := api.Signing_Keys(strx[0])

	// Подписываем транзакцию
	if err := tx.Sign(privKeys, api.Chain); err != nil {
		return nil, errors.Wrapf(err, "Error Sign: ")
	}

	// Отправка транзакции
	resp, err := api.Rpc.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)

	if err != nil {
		return nil, errors.Wrapf(err, "Error BroadcastTransactionSynchronous: ")
	} else {
		var bresp BResp

		bresp.ID = resp.ID
		bresp.BlockNum = resp.BlockNum
		bresp.TrxNum = resp.TrxNum
		bresp.Expired = resp.Expired

		return &bresp, nil
	}
}

func (api *Client) Verify_Trx(strx types.Operation) (bool, error) {
	// Получение необходимых параметров
	props, err := api.Rpc.Database.GetDynamicGlobalProperties()
	if err != nil {
		return false, errors.Wrapf(err, "Error get DynamicGlobalProperties: ")
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return false, err
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Добавление операций в транзакцию
	tx.PushOperation(strx)

	// Получаем необходимый для подписи ключ
	privKeys := api.Signing_Keys(strx)

	// Подписываем транзакцию
	if err := tx.Sign(privKeys, api.Chain); err != nil {
		return false, errors.Wrapf(err, "Error Sign: ")
	}

	// Отправка транзакции
	resp, err := api.Rpc.Database.GetVerifyAuthoruty(tx.Transaction)

	if err != nil {
		return false, errors.Wrapf(err, "Error BroadcastTransactionSynchronous: ")
	} else {
		return resp, nil
	}
}