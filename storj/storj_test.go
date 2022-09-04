package storj_test

import (
	"github.com/fritzkeyzer/go-fkv/storj"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func TestNewKV(t *testing.T) {
	kv, err := storj.NewKV("fkv", os.Getenv("STORJ_ACCESS"))
	if err != nil {
		t.Fatal(err)
	}

	type Object struct {
		Field string
		Data  []byte
	}

	o1 := Object{
		Field: "hello world",
		Data:  make([]byte, 100*1024*1024),
	}
	rand.Read(o1.Data)
	if err := kv.Set("file0", o1); err != nil {
		t.Fatal(err)
	}

	var o2 Object
	if _, err := kv.Get("file0", &o2); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(o1, o2) {
		t.Fatalf("got != want:\n\tgot: %v\n\twant: %v", o2, o1)
	}

	//if err := kv.Delete("file0"); err != nil {
	//	t.Fatal(err)
	//}

	var o3 Object
	found, err := kv.Get("file1", &o3)
	if found {
		t.Fatalf("file1 should not exist!")
	}
	if err != nil {
		t.Fatalf("expect no error, got: %v", err)
	}
}
