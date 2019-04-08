package main

import (
	"bank/account"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var accounts []*account.Account

func getAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := NewMessage()
	message.AddParam("accounts", accounts)
	_ = json.NewEncoder(w).Encode(message)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := NewMessage()
	params := mux.Vars(r)
	for _, item := range accounts {
		id, _ := uuid.Parse(params["id"])
		if item.Id == id {
			message.AddParam("accounts", []*account.Account{item})
			_ = json.NewEncoder(w).Encode(message)
			return
		}
	}
	message.Status = NOT_FOUND
	_ = json.NewEncoder(w).Encode(message)
}

/*
	{
	 "params": {
	  "balance" : 12
	 }
	}
 */
func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := NewMessage()
	_ = json.NewDecoder(r.Body).Decode(message)
	balance := message.GetParam("balance")
	a := account.New(int(balance.(float64)))
	accounts = append(accounts, a)
	message.ClearParams()
	message.Status = OK
	message.AddParam("account", []*account.Account{a})
	_ = json.NewEncoder(w).Encode(message)
}

/*
	{
	 "params": {
	  "from" : "4bc2a229-b30c-4a5a-aaca-d24e50ea96ca",
	  "to" : "2689ca42-d36b-47fd-b54b-06463c2c2a75",
	  "sum" : 10
	 }
	}
 */
func transferRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := NewMessage()
	_ = json.NewDecoder(r.Body).Decode(message)
	from, _ := uuid.Parse(message.GetParam("from").(string))
	to, _ := uuid.Parse(message.GetParam("to").(string))
	sum := int(message.GetParam("sum").(float64))
	message.ClearParams()
	transfer(from, to, sum, message)
	_ = json.NewEncoder(w).Encode(message)
}

func transfer(from uuid.UUID, to uuid.UUID, sum int, message *Message) {
	fromAccount := findAccount(from)
	toAccount := findAccount(to)
	if fromAccount == nil || toAccount == nil {
		message.Status = NOT_FOUND
		return
	}

	if fromAccount.Balance-sum < 0 {
		message.Status = INSUFFICIENT_FUNDS
		return
	}

	fromAccount.Balance -= sum
	toAccount.Balance += sum

	message.Status = OK
}

func findAccount(id uuid.UUID) *account.Account {
	for _, item := range accounts {
		if item.Id == id {
			return item
		}
	}
	return nil
}

func main() {
	router := mux.NewRouter()

	accounts = append(accounts, account.New(100))
	accounts = append(accounts, account.New(100))

	router.HandleFunc("/bank/accounts", getAccounts).Methods("GET")
	router.HandleFunc("/bank/accounts/{id}", getAccount).Methods("GET")
	router.HandleFunc("/bank/accounts/create", createAccount).Methods("POST")
	router.HandleFunc("/bank/accounts/transfer", transferRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(":8888", router))
}
