package main

import (
	"github.com/fritzkeyzer/go-fkv/disk"
	"log"
	"reflect"
)

func main() {
	kv := disk.NewKV("test_dir")
	//kv := mem.NewKV()

	type Object struct {
		Field string
	}

	o1 := Object{Field: "hello world"}
	if err := kv.Set("file0", o1); err != nil {
		log.Fatal(err)
	}

	var o2 Object
	if _, err := kv.Get("file0", &o2); err != nil {
		log.Fatal(err)
	}

	if !reflect.DeepEqual(o1, o2) {
		log.Fatalf("got != want:\n\tgot: %v\n\twant: %v", o2, o1)
	}

	//if err := kv.Delete("file0"); err != nil {
	//	log.Fatal(err)
	//}

	// delete test_dir
	//err := os.RemoveAll("test_dir")
	//if err != nil {
	//	log.Fatal(err)
	//}
}
