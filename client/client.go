package main

import(
	"log"
	"../calculate"
)

func main(){

	args := calculate.Pairs{5,6}

	var reply calculate.Result

	calculate.Add(args, &reply)

	log.Printf("Arith: %d*%d = %d", args.A, args.B, reply)
}

