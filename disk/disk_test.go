package disk_test

import (
	"github.com/fritzkeyzer/go-fkv/disk"
	"os"
	"reflect"
	"testing"
)

func TestNewKV(t *testing.T) {
	kv := disk.NewKV("test_dir")

	type Object struct {
		Field string
	}

	o1 := Object{Field: "hello world"}
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

	if err := kv.Delete("file0"); err != nil {
		t.Fatal(err)
	}

	// delete test_dir
	err := os.RemoveAll("test_dir")
	if err != nil {
		t.Fatal(err)
	}
}
