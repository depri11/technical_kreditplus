package delivery

func (d *transactionDelivery) Routes() {
	d.r.HandleFunc("/transaction", d.CreateTransaction).Methods("POST")
	d.r.HandleFunc("/transaction/{nik}", d.GetTransaction).Methods("GET")
	d.r.HandleFunc("/transaction/{nik}", d.UpdateTransaction).Methods("PUT")
	d.r.HandleFunc("/transaction/{nik}", d.DeleteTransaction).Methods("DELETE")
}
