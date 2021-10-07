package main

import (
	"github.com/JaSei/krakend-martian-correlation-header/martian/modifier"
	"github.com/devopsfaith/krakend-martian/register"
)

func init() {
	register.Set("header.Correlation", []register.Scope{register.ScopeRequest}, func(b []byte) (interface{}, error) {
		return modifier.FromJSON(b)
	})
}

func main() {

}
