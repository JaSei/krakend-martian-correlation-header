// Package elastic registers a header modifier for generating correlation header
package elastic

import (
	"github.com/JaSei/krakend-martian-correlation-header/martian/modifier"
	"github.com/google/martian/parse"
)

func init() {
	parse.Register("header.Correlation", modifier.FromJSON)
}
