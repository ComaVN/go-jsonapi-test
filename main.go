package main

import (
	"fmt"
	"strings"

	"github.com/google/jsonapi"
)

type Foo struct {
	Id   string `jsonapi:"primary,foo"`
	SomeAttr string `jsonapi:"attr,someattr"`
	Rel  *Bar   `jsonapi:"relation,rel"`
	RelRio  *BarRio   `jsonapi:"relation,relrio"`
}

type Bar struct {
	Id string `jsonapi:"primary,bar"`
	SomeAttr string `jsonapi:"attr,someattr"`
}

// https://jsonapi.org/format/#document-resource-identifier-objects
type BarRio struct {
	Id string `jsonapi:"primary,bar"`
}

func main() {
	var sb strings.Builder
	foo := Foo{
		Id:   "1",
		SomeAttr: "a",
		Rel: &Bar{
			Id: "2",
			SomeAttr: "b",
		},
		RelRio: &BarRio{
			Id: "3",
		},
	}
	err := jsonapi.MarshalPayloadWithoutIncluded(&sb, &foo)
	if err != nil {
		panic(err)
	}
	fmt.Println(sb.String())
	sb.Reset()
	err = jsonapi.MarshalPayload(&sb, &foo)
	if err != nil {
		panic(err)
	}
	fmt.Println(sb.String())
}
