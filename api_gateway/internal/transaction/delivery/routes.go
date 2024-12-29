package delivery

func (d *transactionDelivery) Routes() {
	d.r.HandleFunc("/transaction", d.CreateTransaction).Methods("POST")
	d.r.HandleFunc("/transaction/{contract_number}", d.GetTransaction).Methods("GET")
	d.r.HandleFunc("/transaction/{contract_number}", d.UpdateTransaction).Methods("PUT")
	d.r.HandleFunc("/transaction/{contract_number}", d.DeleteTransaction).Methods("DELETE")
}
