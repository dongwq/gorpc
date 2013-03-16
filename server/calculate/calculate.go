package calculate

import(
	"log"
	"../../calculate"
)

type Pairs calculate.Pairs

func (t *Pairs) Add(args Pairs, reply *int) error {
	*reply = args.A + args.B

	log.Printf("recieved %d, %d", args.A, args.B)
	return nil
}

