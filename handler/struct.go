package handler

type Person struct {
	Name string
	Age  int
	City string
}

type Server struct{
	Address string
	GoogleTable *DataSecret
}

type DataSecret struct {
	Token   string `json:"token"`
	TableID string `json:"tableID"`
}