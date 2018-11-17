package rpc

import (
	"log"

	"github.com/parallelcointeam/pod/rpcclient"
)

func WalletInfo() interface{} {
	var client *rpcclient.Client
	connCfg := &rpcclient.ConnConfig{
		Host:         "192.168.1.118:11048",
		User:         "duo",
		Pass:         "pass",
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	getBalance, err := client.GetBalance("")
	if err != nil {
	}
	getUnconfirmedBalance, err := client.GetUnconfirmedBalance("")
	if err != nil {
	}

	getAccountAddress, err := client.GetAccountAddress("")
	if err != nil {
	}

	listUnspent, err := client.ListUnspent()
	if err != nil {
	}

	getReceivedByAccount, err := client.GetReceivedByAccount("")
	if err != nil {
	}
	// getReceivedByAddress, err := client.GetReceivedByAddress(nil)
	// if err != nil {
	// }

	getInfo, err := client.GetInfo()
	if err != nil {
	}

	// getHeaders, err := client.GetHeaders()
	// if err != nil {
	// }

	getBlockChainInfo, err := client.GetBlockChainInfo()
	if err != nil {
	}

	getRawMempool, err := client.GetRawMempool()
	if err != nil {
	}

	bestBlockHash, err := client.GetBestBlockHash()
	if err != nil {
	}

	bestBlock, err := client.GetBlock(bestBlockHash)
	if err != nil {
	}

	getNetTotals, err := client.GetNetTotals()
	if err != nil {
	}

	listTx, err := client.ListTransactions("")
	if err != nil {
	}

	result := map[string]interface{}{
		"Balance":     getBalance,
		"Unconfirmed": getUnconfirmedBalance,
		"Address":     getAccountAddress,
		"Info":        getInfo,

		"BlockChainInfo":    getBlockChainInfo,
		"BestBlock":         bestBlock,
		"BestBlockHash":     bestBlockHash,
		"Unspent":           listUnspent,
		"PooledTxs":         getRawMempool,
		"ReceivedByAccount": getReceivedByAccount,
		// "ReceivedByAddress": getReceivedByAddress,
		"TxList": listTx,

		"NetTotals": getNetTotals,
		"GUI":       "KaputVerzion",
	}
	return result
}

// type immatureCoins struct {
// 	Confirmations int
// 	Amount        float64
// }

// func ImmatureCoins() immatureCoins {
// 	var confirmations int
// 	var amount float64
// 	getList, _ := client.ListUnspent()
// 	for i := range getList {
// 		// x["confirmations"] < 500, getList

// 		confirmations += getList.Confirmations
// 		amount += getList.Amount
// 	}

// 	return immatureCoin{
// 		Confirmations: confirmations,
// 		Amount:        amount,
// 	}
// }

// func ListTx() []btcjson.ListTransactionsResult {
// 	listTx := client.ListTransactions('*', 60)

// 	// txid = filter(lambda x : x["txid"] != txid, list_tx)
// 	return listTx
// }

// func UnspentTx() {
// 	get_list = client.ListUnspent(0)
// 	pass
// }
// func EncryptWallet() {
// 	return client.CreateEncryptedWallet("password")
// }

// func LockWallet() {
// 	return client.WalletLock()
// }
// func AddNode() {
// 	return client.AddNode("ipaddress", "add")
// }

// func Qrcode(address, amount, name, msg string) {
// 	return fmt.Sprintf("coin:%s?amount=%s&label=%s&message=%s", address, amount, name, msg)
// }
