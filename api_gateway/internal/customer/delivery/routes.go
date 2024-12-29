package delivery

func (d *customerDelivery) Routes() {
	d.r.HandleFunc("/customer", d.CreateCustomer).Methods("POST")
	d.r.HandleFunc("/customer/{nik}", d.GetCustomer).Methods("GET")
	d.r.HandleFunc("/customer/{nik}", d.UpdateCustomer).Methods("PUT")
	d.r.HandleFunc("/customer/{nik}", d.DeleteCustomer).Methods("DELETE")
}
