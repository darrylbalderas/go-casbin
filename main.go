package main

import (
	"log"

	"github.com/casbin/casbin"
)

func main() {
	e := casbin.NewEnforcer("model.conf", "policy.csv")

	sub := "group1"  // the user that wants to access a resource.
	dom := "domain1" // the user that wants to access a resource.
	obj := "data1"   // the resource that is going to be accessed.
	act := "read"    // the operation that the user performs on the resource.

	if res := e.Enforce(sub, dom, obj, act); res {
		log.Println("Able to read data")
	} else {
		log.Println("Not able to read data")
	}
}
