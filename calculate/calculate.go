package calculate

import(
	"../service"
)

type Pairs struct{
	A, B int
}

type Result int

const(
	AddService = "Pairs.Add"
	Address = "127.0.0.1:8081"
)

func Add(pairs Pairs, result *Result) error{
	err := service.Call(Address, AddService, pairs, result)
	return err
}

