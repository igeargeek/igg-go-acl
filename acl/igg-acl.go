package acl

import "fmt"

type confI interface {
	MiddleWareAcl()
}

type conf struct {
	Roles      []string
	Permission []string
}

func NewConf(r []string, p []string) confI {
	return &conf{
		Permission: p,
		Roles:      r,
	}
}

func (c *conf) MiddleWareAcl() {
	fmt.Println(*c)
}
