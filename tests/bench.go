package tests

import (
	"github.com/fritzkeyzer/go-fkv"
	"log"
)

func Bench(kv fkv.FKV, n int) {
	type Object struct {
		Field string
	}

	o1 := Object{Field: "hello world"}
	for i := 0; i < n; i++ {
		if err := kv.Set("file0", o1); err != nil {
			log.Fatal(err)
		}

		var o2 Object
		if f, err := kv.Get("file0", &o2); err != nil || !f {
			log.Fatal(err)
		}

		if err := kv.Delete("file0"); err != nil {
			log.Fatal(err)
		}
	}
}
