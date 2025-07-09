package structutil

import (
	"log"
	"testing"
)

func TestSetDefaults(t *testing.T) {
	type A struct {
		A string `default:"a"`
		B int    `default:"1"`
	}
	type d struct {
		E int    `default:"1"`
		F string `default:"f"`
		*A
	}

	dd := new(d)
	log.Printf("%+v \n", dd) // &{E:0 F: A:<nil>}
	if err := SetDefaults(dd); err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v \n", dd)   // &{E:1 F:f A:0xc0000082d0}
	log.Printf("%+v \n", dd.A) // &{A:a B:1}
}
