package cowbody

import "fmt"

type Cowbody struct {
}

func NewCowbody() *Cowbody {
	return &Cowbody{}
}

func (cowbody Cowbody) Decode(v interface{}) error {
	fmt.Printf("conclusion cowbody")
	return nil
}
