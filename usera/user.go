package usera

import "example.com/hello/doer"

type Usera struct {
	Doer doer.Doer
}

func (u *Usera) Use() error {
	return u.Doer.DoSomething(123, "Hello GoMock")
}
