package main

import "github.com/igeargeek/igg-go-acl/acl"

func main() {
	roles := []string{
		"admin",
		"editor",
	}

	permission := []string{
		"read",
		"write",
	}

	middleware := acl.NewConf(roles, permission)
	middleware.MiddleWareAcl()
}
