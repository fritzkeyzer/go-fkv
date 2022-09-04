package tests

import (
	"fmt"
	"github.com/fritzkeyzer/go-fkv"
	"github.com/fritzkeyzer/go-fkv/disk"
	"github.com/fritzkeyzer/go-fkv/mem"
	"github.com/fritzkeyzer/go-fkv/storj"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestDiskKV(t *testing.T) {
	kv := disk.NewKV("test_dir")
	test(kv, t)
}

func TestMemKV(t *testing.T) {
	kv := mem.NewKV()
	test(kv, t)
}

func TestStorjKV(t *testing.T) {
	token := os.Getenv("STORJ_ACCESS")
	if len(strings.TrimSpace(token)) == 0 {
		fmt.Println("Unable to test storj KV. 'STORJ_ACCESS' env variable not set")
		return
	}

	kv, err := storj.NewKV("test-bucket", token)
	if err != nil {
		t.Fatal(err)
	}
	test(kv, t)
}

func test(kv fkv.FKV, t *testing.T) {
	type Object struct {
		Field string
	}

	o1 := Object{Field: "hello world"}
	if err := kv.Set("file0", o1); err != nil {
		t.Fatal(err)
	}

	var o2 Object
	if f, err := kv.Get("file0", &o2); err != nil {
		t.Fatal(err)
	} else if !f {
		t.Fatalf("could not find fil0")
	}

	if !reflect.DeepEqual(o1, o2) {
		t.Fatalf("got != want:\n\tgot: %v\n\twant: %v", o2, o1)
	}

	if err := kv.Delete("file0"); err != nil {
		t.Fatal(err)
	}

	var o3 Object
	if found, err := kv.Get("file0", &o3); err != nil {
		t.Fatalf("expect no error, got: %v", err)
	} else if found {
		t.Fatalf("file0 should not exist!")
	}
}

func BenchmarkDiskKV(b *testing.B) {
	kv := disk.NewKV("test_dir")
	bench(kv, b.N)
}

func BenchmarkMemKV(b *testing.B) {
	kv := mem.NewKV()
	bench(kv, b.N)
}

func BenchmarkStorjKV(b *testing.B) {
	kv, err := storj.NewKV("test-bucket", os.Getenv("STORJ_ACCESS"))
	if err != nil {
		log.Fatal(err)
	}
	bench(kv, b.N)
}

func bench(kv fkv.FKV, n int) {
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
