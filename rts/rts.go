package rts

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parallelcointeam/dgw/rpc"
)

var templates = make(map[string]*template.Template)

var last string

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["index"] = template.Must(template.ParseFiles("tpl/base.gohtml", "tpl/wallet/index.gohtml", "tpl/part/side.gohtml"))
	templates["send"] = template.Must(template.ParseFiles("tpl/base.gohtml", "tpl/wallet/send.gohtml", "tpl/part/side.gohtml"))
	templates["receive"] = template.Must(template.ParseFiles("tpl/base.gohtml", "tpl/wallet/receive.gohtml", "tpl/part/side.gohtml"))
	templates["setup"] = template.Must(template.ParseFiles("tpl/base.gohtml", "tpl/wallet/setup.gohtml", "tpl/part/side.gohtml"))
	templates["history"] = template.Must(template.ParseFiles("tpl/base.gohtml", "tpl/wallet/history.gohtml", "tpl/part/side.gohtml"))

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
		// "Time":       time,
		// uiVersion : currentApp.config['UI_VERSION']),
	}
	renderTemplate(w, "index", "base", data)
}

func SendHandler(w http.ResponseWriter, r *http.Request) {
	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
	}
	renderTemplate(w, "send", "base", data)
}
func SendSelectedAddressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
		"Address":    address,
	}
	renderTemplate(w, "send", "base", data)
}
func ReceiveHandler(w http.ResponseWriter, r *http.Request) {
	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
	}
	renderTemplate(w, "receive", "base", data)
}
func ReceiveSelectedAddressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
		"Address":    address,
	}
	renderTemplate(w, "receive", "base", data)
}

func NewAddressHandler(w http.ResponseWriter, r *http.Request) {
	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
	}
	renderTemplate(w, "index", "base", data)
}
func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
	}
	renderTemplate(w, "history", "base", data)
}
func SetupHandler(w http.ResponseWriter, r *http.Request) {
	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
	}
	renderTemplate(w, "setup", "base", data)
}
func EncryptWalletHandler(w http.ResponseWriter, r *http.Request) {
	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
	}
	renderTemplate(w, "index", "base", data)
}

func AddNodeHandler(w http.ResponseWriter, r *http.Request) {
	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
	}
	renderTemplate(w, "index", "base", data)
}
func SendReqHandler(w http.ResponseWriter, r *http.Request) {
	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
	}
	renderTemplate(w, "index", "base", data)
}
func SelectedCmdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cmd := vars["cmd"]

	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
		"Cmd":        cmd,
	}
	renderTemplate(w, "index", "base", data)
}
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	walletInfo := rpc.WalletInfo()
	data := map[string]interface{}{
		"WalletInfo": walletInfo,
	}
	renderTemplate(w, "index", "base", data)
}

func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
