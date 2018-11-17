/*************************************
    __    __    __    __    __    __
   /  \  /  \  /  \  /  \  /  \  /  \
  /  __\/  __\/  __\/  __\/  __\/  __\
 /  /__/  /__/  /__/  /__/  /__/  /__/
  \   / \   / \   / \   / \   / \   /
   \_/   \_/   \_/   \ /   \_/   \_/

App : dgw
Name: DUO GUI Wallet

*************************************/
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parallelcointeam/dgw/rts"
)

func main() {

	css := http.FileServer(http.Dir("./static/"))
	css = http.StripPrefix("/css/", css)

	r := mux.NewRouter()
	r.PathPrefix("/css").Handler(css)

	r.Path("/").HandlerFunc(rts.IndexHandler).Name("Index")
	r.Path("/send").HandlerFunc(rts.SendHandler).Name("Send")
	r.Path("/send/{address}").HandlerFunc(rts.SendSelectedAddressHandler).Name("SendSelectedAddress")

	r.Path("/receive").HandlerFunc(rts.ReceiveHandler).Name("Receive")
	r.Path("/receive/{address}").HandlerFunc(rts.ReceiveSelectedAddressHandler).Name("ReceiveSelectedAddress")

	r.Path("/newaddress").HandlerFunc(rts.NewAddressHandler).Name("NewAddress")

	r.Path("/history").HandlerFunc(rts.HistoryHandler).Name("History")

	r.Path("/setup").HandlerFunc(rts.SetupHandler).Name("setup")

	r.Path("/encryptwallet").HandlerFunc(rts.EncryptWalletHandler).Name("EncryptWallet")

	r.Path("/addnode").HandlerFunc(rts.AddNodeHandler).Name("AddNode")

	r.Path("/sendreq").HandlerFunc(rts.SendReqHandler).Name("SendReq")
	r.Path("/sendreq/{cmd}").HandlerFunc(rts.SelectedCmdHandler).Name("SelectedCmd")

	r.Path("/download").HandlerFunc(rts.DownloadHandler).Name("download")

	go log.Fatal(http.ListenAndServe(":5599", r))

}
